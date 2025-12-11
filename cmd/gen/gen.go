package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

type (
	//  cldr-core/availableLocales.json
	Locales struct {
		AvailableLocales struct {
			Full []string `json:"full"`
		} `json:"availableLocales"`
	}

	// cldr-core/supplemental/weekData.json
	WeekData struct {
		Supplemental struct {
			WeekData struct {
				// Map of country to day name:
				//   "AD": "mon",
				//   "AE": "sat",
				FirstDay map[string]string `json:"firstDay"`
			} `json:"weekData"`
		} `json:"supplemental"`
	}

	// cldr-dates-full/main/*/ca-gregorian.json
	Dates struct {
		Main map[string]struct { // Key is locale name, "nl" or "en-NZ"
			Identity struct {
				Language  string `json:"language"`
				Territory string `json:"territory"`
				Script    string `json:"script"`
			} `json:"identity"`

			Dates struct {
				Calendars struct {
					Gregorian Gregorian `json:"gregorian"`
				} `json:"calendars"`
			}
		} `json:"main"`
	}
	Gregorian struct {
		Months struct {
			Format struct {
				Abbreviated map[int]string `json:"abbreviated"`
				Wide        map[int]string `json:"wide"`
			} `json:"format"`
		} `json:"months"`

		Days struct {
			Format struct {
				Abbreviated map[string]string `json:"abbreviated"`
				Wide        map[string]string `json:"wide"`
			} `json:"format"`
		} `json:"days"`
		DayPeriods struct {
			Format struct {
				Abbreviated struct {
					AM string `json:"am"`
					PM string `json:"pm"`
				} `json:"abbreviated"`
			} `json:"format"`
		} `json:"dayPeriods"`

		DateFormats struct {
			Full   string          `json:"full"`
			Long   string          `json:"long"`
			Medium string          `json:"medium"`
			Short  json.RawMessage `json:"short"`
		} `json:"dateFormats"`
		TimeFormats struct {
			Full   string `json:"full"`
			Long   string `json:"long"`
			Medium string `json:"medium"`
			Short  string `json:"short"`
		} `json:"timeFormats"`
		DateTimeFormats struct {
			Full   string `json:"full"`
			Long   string `json:"long"`
			Medium string `json:"medium"`
			Short  string `json:"short"`
		} `json:"dateTimeFormats"`
	}

	// supplemental/ordinals.json
	// Ordinals struct {
	// 	Supplemental struct {
	// 		PluralsTypeOrdinal map[string]struct {
	// 			PluralRuleCountZero  string `json:"pluralRule-count-zero"`
	// 			PluralRuleCountOne   string `json:"pluralRule-count-one"`
	// 			PluralRuleCountTwo   string `json:"pluralRule-count-two"`
	// 			PluralRuleCountFew   string `json:"pluralRule-count-few"`
	// 			PluralRuleCountMany  string `json:"pluralRule-count-many"`
	// 			PluralRuleCountOther string `json:"pluralRule-count-other"`
	// 		}
	// 	} `json:"plurals-type-ordinal"`
	// }

	// TODO cldr-misc-full/main/nl/contextTransforms.json
	// TODO cldr-misc-full/main/nl/delimiters.json
	// TODO cldr-rbnf/rbnf

	All struct {
		Language, Region, Script string
		FirstDay                 int
		Calendar                 Gregorian
	}
)

// Can be string or object as {"_value": "d/M/yy", "_numbers": "M=romanlow"}
func shortString(b json.RawMessage) string {
	if len(b) == 0 {
		return ""
	}

	if b[0] == '"' {
		var s string
		err := json.Unmarshal(b, &s)
		F(err)
		return s
	}

	var s struct {
		Value string `json:"_value"`
	}
	err := json.Unmarshal(b, &s)
	F(err)
	return s.Value
}

// Use the JSON data because the XML is a pain to parse.
//
// git clone --depth=1 git@github.com:unicode-org/cldr-json.git
//
// http://cldr.unicode.org/translation/date-time-1/date-time-patterns
func main() {
	var locales Locales
	readJSON(&locales, "cldr-core/availableLocales.json")

	var weekData WeekData
	readJSON(&weekData, "cldr-core/supplemental/weekData.json")

	var all []All
	for _, l := range locales.AvailableLocales.Full {
		var dates Dates
		readJSON(&dates, "cldr-dates-full/main/"+l+"/ca-gregorian.json")

		all = append(all, All{
			Language: dates.Main[l].Identity.Language,
			Region:   dates.Main[l].Identity.Territory,
			Script:   dates.Main[l].Identity.Script,
			Calendar: dates.Main[l].Dates.Calendars.Gregorian,
		})
	}

	// Replace the {0} and {1} references in the datetimes. Bit more space, but
	// it's just easier later on.
	for i := range all {
		c := all[i].Calendar
		c.DateTimeFormats.Full = strings.ReplaceAll(c.DateTimeFormats.Full, "{0}", c.TimeFormats.Full)
		c.DateTimeFormats.Long = strings.ReplaceAll(c.DateTimeFormats.Long, "{0}", c.TimeFormats.Long)
		c.DateTimeFormats.Medium = strings.ReplaceAll(c.DateTimeFormats.Medium, "{0}", c.TimeFormats.Medium)
		c.DateTimeFormats.Short = strings.ReplaceAll(c.DateTimeFormats.Short, "{0}", c.TimeFormats.Short)
		c.DateTimeFormats.Full = strings.ReplaceAll(c.DateTimeFormats.Full, "{1}", c.DateFormats.Full)
		c.DateTimeFormats.Long = strings.ReplaceAll(c.DateTimeFormats.Long, "{1}", c.DateFormats.Long)
		c.DateTimeFormats.Medium = strings.ReplaceAll(c.DateTimeFormats.Medium, "{1}", c.DateFormats.Medium)
		c.DateTimeFormats.Short = strings.ReplaceAll(c.DateTimeFormats.Short, "{1}", shortString(c.DateFormats.Short))

		n, ok := weekData.Supplemental.WeekData.FirstDay[all[i].Region]
		if !ok {
			n = weekData.Supplemental.WeekData.FirstDay[strings.ToUpper(all[i].Language)]
		}

		all[i].FirstDay = map[string]int{"": 1, "sun": 0, "mon": 1, "wed": 2, "tue": 3, "thu": 4, "fri": 5, "sat": 6}[n]

		all[i].Calendar = c
	}

	funs := make(map[string][]string)
	for _, a := range all {
		f := makeFun(a)
		funs[f] = append(funs[f], a.Tag())
	}

	// Sort so we're sure "fr" is before "fr-CA"
	for k, v := range funs {
		sort.Strings(v)
		funs[k] = v
	}

	var (
		show   = make(map[string]string)
		linked = make(map[string]string)
	)
	for k, v := range funs {
		// If the first one is language, i.e. "fr", then we don't need to write
		// anything else as "fr-XX" will match "fr" too. So we can just write
		// out the first one.
		//
		// Only if it's e.g. "fr-DJ fr-DZ" do we need to explicitly link.
		if len(v) > 1 && len(v[0]) > 2 {
			for _, l := range v[1:] {
				linked[l] = v[0]
			}
		}

		show[v[0]] = k
	}

	out := new(strings.Builder)
	out.WriteString(`package z18n

import "golang.org/x/text/language"

var mapping = map[language.Tag]*localize {
`)
	for _, l := range locales.AvailableLocales.Full {
		f := strings.ReplaceAll(l, "-", "_")
		_, ok := show[l]
		if ok {
			fmt.Fprintf(out, "language.MustParse(%q): l_%s,\n", l, f)
			continue
		}

		link, ok := linked[l]
		if !ok {
			continue
		}

		f = strings.ReplaceAll(link, "-", "_")
		fmt.Fprintf(out, "language.MustParse(%q): l_%s, // Link\n", l, f)
	}

	out.WriteString("}\n\n")

	for _, l := range locales.AvailableLocales.Full {
		code, ok := show[l]
		if !ok {
			continue
		}
		fmt.Fprintf(out, "var l_%s = &%s\n", strings.ReplaceAll(l, "-", "_"), code)
	}

	f, err := format.Source([]byte(out.String()))
	if err != nil {
		fmt.Println(err)
		fmt.Println(out.String())
		os.Exit(1)
	}
	fmt.Print(string(f))
}

func readJSON(dst interface{}, file string) {
	data, err := os.ReadFile("cldr-json/cldr-json/" + file)
	if err != nil {
		F(fmt.Errorf("reading %q: %s", file, err))
	}
	err = json.Unmarshal(data, dst)
	if err != nil {
		F(fmt.Errorf("reading %q: %s", file, err))
	}
}

func F(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (a All) Tag() string {
	t := []string{a.Language}
	if a.Script != "" {
		t = append(t, a.Script)
	}
	if a.Region != "" {
		t = append(t, a.Region)
	}
	return strings.Join(t, "-")
}

func (a All) Func() string {
	t := []string{a.Language}
	if a.Script != "" {
		t = append(t, a.Script)
	}
	if a.Region != "" {
		t = append(t, a.Region)
	}
	return strings.ToUpper(strings.Join(t, "_"))
}

func makeFun(a All) string {
	b := new(strings.Builder)
	c := a.Calendar

	fmt.Fprintf(b, "localize{\nam: %q,\npm: %q,\nfirstDay: %d,\ntwentyFourHours: false,\n",
		c.DayPeriods.Format.Abbreviated.AM, c.DayPeriods.Format.Abbreviated.PM, a.FirstDay)

	fmt.Fprintf(b, "months: months{\nshort: %s,\n long: %s,\n},\n",
		months(c.Months.Format.Abbreviated), months(c.Months.Format.Wide))
	fmt.Fprintf(b, "days: days{\nshort: %s,\nlong: %s,\n},\n",
		days(c.Days.Format.Abbreviated), days(c.Days.Format.Wide))

	fmt.Fprintf(b, "timeFormat: timeFormat{\ndate: %s,\ntime: %s,\ndatetime: %s,\n},\n",
		dateformats(c.DateFormats.Full, c.DateFormats.Long, c.DateFormats.Medium, shortString(c.DateFormats.Short)),
		dateformats(c.TimeFormats.Full, c.TimeFormats.Long, c.TimeFormats.Medium, c.TimeFormats.Short),
		dateformats(c.DateTimeFormats.Full, c.DateTimeFormats.Long, c.DateTimeFormats.Medium, c.DateTimeFormats.Short))
	b.WriteString("}")
	return b.String()
}

func months(m map[int]string) string {
	b := new(strings.Builder)
	b.WriteString("[12]string{")
	for i := 1; i <= 12; i++ {
		fmt.Fprintf(b, "%q, ", m[i])
	}
	b.WriteRune('}')
	return b.String()
}

func days(m map[string]string) string {
	b := new(strings.Builder)
	b.WriteString("[7]string{")
	// Start at Sunday to match Go Weekday()
	for _, i := range []string{"sun", "mon", "wed", "tue", "thu", "fri", "sat"} {
		fmt.Fprintf(b, "%q, ", m[i])
	}
	b.WriteRune('}')
	return b.String()
}

func dateformats(s ...string) string {
	r := make([]string, len(s))
	for i, ss := range s {
		r[i] = dateformat(ss)
	}
	return "[4]string{" + strings.Join(r, ", ") + "}"
}

// https://www.unicode.org/reports/tr35/tr35-dates.html#Date_Format_Patterns
var replFormat = strings.NewReplacer(
	// Date
	"yy", "06", "y", "2006",
	"MMMM", "January", "MMM", "Jan", "MM", "01", "M", "1",
	"dd", "02", "d", "2",
	"EEEE", "Monday", "EEE", "Mon", "EE", "Mon", "E", "Mon", "eeee", "Monday", "eee", "Mon",

	// Time
	// Note: "H" always printed as "00", never "0"; Go doesn't support that (I think?)
	"a", "PM",
	"hh", "03", "h", "3", "HH", "15", "H", "15",
	"mm", "04", "m", "4",
	"ss", "05", "s", "5",

	// Timezone
	// Full TZ name; not supported in Go
	//
	// This is stupid anyway; do you know what "Western Argentina Time" or
	// "Bhutan Time" is? Just use abbrevation + offset.
	"zzzz", "MST -0700",

	"zzz", "MST", "zz", "MST", "z", "MST",

	// Z..ZZZ	-0700
	// ZZZZ		GMT-07:00
	// ZZZZZ	-07:00
	// O		GMT-8
	// OOOO		GMT-08:00
)

func dateformat(s string) string {
	var (
		buf    bytes.Buffer
		result bytes.Buffer
		quote  bool
	)
	for _, c := range s {
		// Pattern fields are always letters, but any text between single quotes
		// should be shown as literal text:
		// M 'in' y
		if c == '\'' {
			quote = !quote
			continue
		}
		if quote {
			result.WriteRune(c)
			continue
		}
		if unicode.IsLetter(c) {
			buf.WriteRune(c)
			continue
		}

		//

		if buf.Len() > 0 {
			result.WriteString(replFormat.Replace(buf.String()))
			buf.Reset()
		}

		result.WriteRune(c)
	}
	if buf.Len() > 0 {
		result.WriteString(replFormat.Replace(buf.String()))
	}

	return fmt.Sprintf("%q", result.String())
}
