package z18n

import (
	"context"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
	"zgo.at/zstd/zstring"
)

var ctxkey = &struct{}{}

// Locale is a single localisation.
type Locale struct {
	*localize
	bundle *Bundle
	tags   []language.Tag
}

// TimeFormat controls which standard time format to print
type TimeFormat uint8

type (
	// A localize contains all locale data for one language or region.
	//
	// TODO: this can probably be folded in to the Locale type; a number of these
	// methods can be exported as well.
	localize struct {
		firstDay        uint8  // First day of week; 0=sun, 1=mon, ..., 6=sat
		twentyFourHours bool   // 24-hour clock?
		am, pm          string // Names for AM and PM
		months          months
		days            days
		timeFormat      timeFormat
	}
	months struct { // Month names, 0=january
		short [12]string
		long  [12]string
	}
	days struct { // Day names, 0=sunday
		short [7]string
		long  [7]string
	}
	timeFormat struct { // Full, long, medium, short
		date     [4]string
		time     [4]string
		datetime [4]string
	}
)

// Standard date, time, and datetime lengths.
//
// See localize.Date(), localize.Time(), and localize.Datetime() for some examples.
//
// For time and datetime the CLDR specifies that a full time zone name should be
// printed for most regions, such as "Western Argentina Time" or "Bhutan Time".
// Do you know what actual TZ those are? Probably not. So this replaces that
// with "MST -0700" instead (abbrevated timezone name + offset), which is far
// more useful in almost all cases (and also avoids having to include a large
// list of timezone names for every language).
const (
	TimeFormatFull TimeFormat = iota
	TimeFormatLong
	TimeFormatMedium
	TimeFormatShort
)

// With returns a copy of the context with the Locale as a value.
func With(ctx context.Context, l *Locale) context.Context {
	return context.WithValue(ctx, ctxkey, l)
}

// Get the Locale value from the context.
func Get(ctx context.Context) *Locale {
	l, _ := ctx.Value(ctxkey).(*Locale)
	return l
}

// T translates a string, like Locale.T. the Locale is fetched from the context.
func T(ctx context.Context, id string, data ...interface{}) string {
	l := Get(ctx)
	if l == nil {
		return "%(z18n ERROR: T: no context)"
	}
	return l.T(id, data...)
}

// Thtml is like T, but returns template.HTML instead of a string.
func Thtml(ctx context.Context, id string, data ...interface{}) template.HTML {
	return template.HTML(T(ctx, id, data...))
}

// T translates a message for this locale.
//
// It will return the message in the bundler's defaultLang if the message is not
// translated in this language (yet).
//
// The ID can contain any character except a |. Everything after the first | is
// used as the default message.
//
// Variables can be inserted as %(varname), and text can be wrapped in HTML tags
// with %[varname translated text]. Wrapping in HTML requires passing a Tagger
// interface (such as Tag).
//
// Pass N() as any argument to apply pluralisation.
func (l Locale) T(id string, data ...interface{}) string {
	def := ""
	if p := strings.Index(id, "|"); p > -1 {
		id, def = id[:p], id[p+1:]
	}

	// data can contain be any of four things:
	//   - A single variable, in case there is just one placeholder.
	//   - A map for mapped variables, in case of multiple placeholders
	//   - Either of the above with a Plural
	//   - Anything else for a single variable.
	// The z18n tool checks the parameters, so we don't need to do a lot of that
	// here and we can be a bit relaxed.
	var (
		pl     *Plural
		params = make(P)
		oneVar bool
	)
	for _, d := range data {
		if p, ok := d.(Plural); ok {
			pl = &p
		} else if p, ok := d.(P); ok {
			params = p
		} else if p, ok := d.(map[string]interface{}); ok {
			params = p
		} else if p, ok := d.(map[string]string); ok {
			for k, v := range p {
				params[k] = v
			}
		} else {
			oneVar, params = true, P{"": d}
		}
	}
	if pl != nil {
		params["n"] = pl
	}

	// TODO: why do this here? There must be a reason, but I forgot
	//
	// Match returns the best match for any of the given tags, along with
	// a unique index associated with the returned tag and a confidence
	// score.
	_, i, _ := l.bundle.matcher.Match(l.tags...)
	tag := l.bundle.tags[i]

	m, ok := l.bundle.messages[tag]
	if ok {
		msg, ok := m[id]
		if ok && msg.Default != "" {
			msg.ID = id
			msg.data = params
			msg.oneVar = oneVar
			msg.Plural = pl
			msg.bundle = l.bundle
			msg.tag = tag
			return msg.Display(&l)
		}
	}

	return Msg{
		bundle:  l.bundle,
		ID:      id,
		Default: def,
		data:    params,
		oneVar:  oneVar,
		Plural:  pl,
		tag:     tag,
	}.Display(&l)
}

// FormatTime works like time.Format, except that names such as "Monday" and
// "March" are translated to this locale.
func (l localize) FormatTime(t time.Time, format string) string {
	// TODO: probably slow, and ugly. But it gets me ahead for now.
	format = strings.ReplaceAll(format, "Monday", l.days.long[t.Weekday()])
	format = strings.ReplaceAll(format, "January", l.months.long[t.Month()-1])
	format = strings.ReplaceAll(format, "Mon", l.days.short[t.Weekday()])
	format = strings.ReplaceAll(format, "Jan", l.months.short[t.Month()-1])

	// TODO: AM/PM
	return t.Format(format)
}

// Date formats the time t in the standard date format for this language.
func (l localize) Date(t time.Time, format TimeFormat) string {
	return l.FormatTime(t, l.timeFormat.date[format])
}

// Time formats the time t in the standard time format for this language.
func (l localize) Time(t time.Time, format TimeFormat) string {
	return l.FormatTime(t, l.timeFormat.time[format])
}

// Datetime formats the time t in the standard datetime format for this
// language.
func (l localize) Datetime(t time.Time, format TimeFormat) string {
	return l.FormatTime(t, l.timeFormat.datetime[format])
}

// WeekdayName gets the name of the weekday in t.
//
// If timeFormatShort is passed it will return a short name, such as "Wed". Any
// other value will get the full name: "Wednesday".
func (l localize) WeekdayName(t time.Time, n TimeFormat) string {
	if n == TimeFormatShort {
		return l.days.short[t.Weekday()]
	}
	return l.days.long[t.Weekday()]
}

// MonthName gets the name of the month in t.
//
// If timeFormatShort is passed it will return a short name, such as "Apr". Any
// other value will get the full name: "April".
func (l localize) MonthName(t time.Time, n TimeFormat) string {
	if n == TimeFormatShort {
		return l.months.short[t.Month()-1]
	}
	return l.months.long[t.Month()-1]
}

// StartOfWeek gets the day this weeks starts.
//
// 0 is Sunday, 1 is Monday, 6 is Saturday.
func (l localize) StartOfWeek() int { return int(l.firstDay) }

// TwentyFourHours reports if this region uses a 24-hour clock.
func (l localize) TwentyFourHours() bool { return l.twentyFourHours }

// TODO: support formatting telephone numbers (LC_TELEPHONE)
func l10n(l *Locale, tag language.Tag, v interface{}, modifiers map[string]string) string {
	switch vv, t := simplify(v); t {
	case reflect.String:
		return vv.String()

	case reflect.Float64:
		var opts []number.Option
		if f, ok := modifiers[""]; ok {
			min, max := zstring.Split2(f, ":")

			var maxi, mini int
			if max == "" { // TODO: error?
				maxi, _ = strconv.Atoi(min)
			} else {
				mini, _ = strconv.Atoi(min)
				maxi, _ = strconv.Atoi(max)
			}

			opts = append(opts, number.MinFractionDigits(mini), number.MaxFractionDigits(maxi))
		}
		if _, ok := modifiers["raw"]; ok {
			opts = append(opts, number.NoSeparator())
		}
		if _, ok := modifiers["$"]; ok {
			// TODO: this is actually not so great:
			//
			// - Always prints as "€ 5.00", but just "€ 5" is okay.
			// - Not all currencies have fractions, so need to be clever-y about
			//   that anyway.
			// - Always prints as "€ 5,00", but e.g. German style is "5,00 €"

			// And also:
			// - No way to specify the currency being used; important because
			//   otherwise we have the wrong symbol. It just uses the locale's
			//   symbol now, but that can be wrong.
			//
			// Solution: Create something like:
			//
			//   z18n.Currency("EUR", 4.12)
			//   z18n.Currency("USD", 4.12)
			//
			// Then we don't need the $ in the variable any more either, and
			// then combine the currency info + locale info to print something
			// nice.

			// c, _ := currency.FromTag(language.German)
			// cc := c.Amount(vv.Float())
			// ccc := currency.Symbol(cc)
			// return message.NewPrinter(tag).Sprintf("%d", ccc)
		}

		//fmt.Println(tag)
		return message.NewPrinter(tag).Sprintf("%f", number.Decimal(vv.Float(), opts...))

	default:
		switch vvv := vv.Interface().(type) {
		case time.Time:
			return l10nDate(vvv, l, modifiers)
		case fmt.Stringer:
			return vvv.String()
		}

		return vv.String()
	}
}

func l10nDate(t time.Time, l *Locale, modifiers map[string]string) string {
	f := l.Datetime
	if _, ok := modifiers["date"]; ok {
		f = l.Date
	} else if _, ok := modifiers["time"]; ok {
		f = l.Time
	}

	if format, ok := modifiers[""]; ok {
		return l.FormatTime(t, format)
	}

	tf := TimeFormatMedium
	if _, ok := modifiers["full"]; ok {
		tf = TimeFormatFull
	} else if _, ok := modifiers["long"]; ok {
		tf = TimeFormatLong
	} else if _, ok := modifiers["medium"]; ok {
		tf = TimeFormatMedium
	} else if _, ok := modifiers["short"]; ok {
		tf = TimeFormatShort
	}
	return f(t, tf)
}
