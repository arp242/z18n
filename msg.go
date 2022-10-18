package z18n

import (
	"fmt"
	"html"
	"html/template"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"zgo.at/z18n/plural"
	"zgo.at/zstd/zstring"
)

// Msg is a localized message.
type Msg struct {
	bundle *Bundle
	tag    language.Tag
	data   P

	// TODO: would be better if we would also allow:
	//
	//   %(varname) %[text]
	//
	// Without having to create a map.
	oneVar bool

	ID      string  // Message ID.
	Plural  *Plural // Plural value; may be nil.
	Default string  // CLDR "other" plural (default is more intuitive IMO).
	Zero    string  // CLDR "zero" plural.
	One     string  // CLDR "one" plural.
	Two     string  // CLDR "two" plural.
	Few     string  // CLDR "few" plural.
	Many    string  // CLDR "many" plural.
}

type (
	// Plural signals to T that this parameter is used to pluralize the string,
	// rather than a data parameter.
	Plural int

	// P is a shortcut for T() map parameters.
	P map[string]interface{}

	Tagger interface {
		Open() string
		Close() string
		Text() string
	}
)

func (p Plural) String() string { return strconv.Itoa(int(p)) }

// N returns a plural of n.
func N(n int) Plural { return Plural(n) }

type tag struct {
	tag, content string
	innerHTML    string
}

func (t tag) Open() string {
	if t.content == "" {
		return "<" + t.tag + ">"
	}
	return "<" + t.tag + " " + t.content + ">"
}
func (t tag) Close() string  { return "</" + t.tag + ">" }
func (t tag) Text() string   { return t.innerHTML }
func (t tag) String() string { return t.Open() + html.EscapeString(t.Text()) + t.Close() }

// Tag creates a new tag.
func Tag(tagName, content string, innerHTML ...string) tag {
	return tag{tag: tagName, content: content, innerHTML: strings.Join(innerHTML, "")}
}

func (m Msg) tpl(l *Locale, str string) string {
	var (
		tags  = zstring.IndexPairs(str, "%[", "]")
		vars  = zstring.IndexPairs(str, "%(", ")")
		total = len(tags) + len(vars)
	)
	// fmt.Printf("XXX %q\n", str)
	// if len(m.data) > 0 || total > 0 {
	// 	fmt.Printf("\tvars: %v; tags: %v\n\t%#v\n\n", tags, vars, m.data)
	// }
	if total == 0 {
		if l.bundle.Mark {
			return "«" + str + "»"
		}
		return str
	}

	str = m.tplTags(str, tags)

	// We need to check this again, as the indexes probably changed.
	// TODO: this can be more efficient.
	vars = zstring.IndexPairs(str, "%(", ")")

	if l.bundle.Mark {
		return "«" + m.tplVars(l, str, vars) + "»"
	}
	return m.tplVars(l, str, vars)
}

func (m Msg) tplTags(str string, pairs [][]int) string {
	for _, p := range pairs {
		start, end := p[0], p[1]
		text := str[start+2 : end]
		varname := ""
		if len(text) > 0 && text[0] == '%' {
			varname, text = zstring.Split2(text, " ")
			varname = varname[1:]
		}

		key := varname
		if m.oneVar {
			key = ""
		}
		value, ok := m.data[key]
		if !ok { // TODO: update CLI to detect this
			str = str[:start] + fmt.Sprintf("!(z18n ERROR: no value for tag %q)", varname) + str[end+1:]
			continue
		}
		t, ok := value.(Tagger)
		if !ok { // TODO: update CLI to detect this.
			str = str[:start] + "!(z18n ERROR: value for " + varname + " is not a Tagger)" + str[end+1:]
			continue
		}

		if tt := t.Text(); tt != "" {
			text = tt
		} else if !m.bundle.NoHTML { // TODO: check for template.HTML
			text = template.HTMLEscapeString(text)
		}
		str = str[:start] + t.Open() + text + t.Close() + str[end+1:]
	}
	return str
}

var funcmap = map[string]func(string) string{
	"lower":       strings.ToLower,
	"upper":       strings.ToUpper,
	"upper_first": zstring.UpperFirst,
	"html":        template.HTMLEscapeString, // Don't run if already escaped.

	// These are not "true" functions but tags used later.
	"raw":    nil,
	"$":      nil,
	"%":      nil,
	"‰":      nil,
	"date":   nil,
	"time":   nil,
	"full":   nil,
	"long":   nil,
	"medium": nil,
	"short":  nil,
}

func (m Msg) tplVars(l *Locale, str string, pairs [][]int) string {
	for _, p := range pairs {
		start, end := p[0], p[1]
		varname, fun := zstring.Split2(str[start+2:end], " ")

		var (
			funs []string
			tags = make(map[string]string)
		)
		for _, name := range strings.Fields(fun) {
			f, ok := funcmap[name]
			if !ok { // The special "formatting": layout for date or float
				tags[""] += " " + name
			} else if f == nil {
				tags[name] = ""
			} else {
				if name == "html" && m.bundle.NoHTML {
					continue
				}
				funs = append(funs, name)
			}
		}
		if tags[""] != "" {
			tags[""] = tags[""][1:]
		}

		key := varname
		if m.oneVar {
			key = ""
		}
		val, ok := m.data[key]
		if !ok { // TODO: update CLI to detect this
			str = str[:start] + fmt.Sprintf("!(z18n ERROR: no value for variable %q)", varname) + str[end+1:]
			continue
		}

		v := l10n(l, m.tag, val, tags)

		_, isTagger := val.(Tagger)
		if !isTagger { // TODO: check for template.HTML in general?
			if _, ok := tags["raw"]; !ok && !m.bundle.NoHTML {
				v = template.HTMLEscapeString(v)
			}
		}

		for _, name := range funs {
			v = funcmap[name](v)
		}

		str = str[:start] + v + str[end+1:]
	}
	return str
}

// String displays this string as "other", or the ID if this isn't set.
func (m Msg) Display(l *Locale) string {
	if m.Plural == nil {
		if m.Default != "" {
			return m.tpl(l, m.Default)
		}
		return m.ID
	}

	// Only error failure is on invalid type, so it's safe to ignore.
	op, _ := plural.NewOperands(int(*m.Plural))

	form := m.bundle.pluralRules.Rule(m.tag).PluralFormFunc(op)
	var s string
	switch form {
	case plural.Zero:
		s = m.Zero
	case plural.One:
		s = m.One
	case plural.Two:
		s = m.Two
	case plural.Few:
		s = m.Few
	case plural.Many:
		s = m.Many
	case plural.Other:
		s = m.Default
	}
	if s == "" {
		if form == plural.Other {
			return "unknown"
		}
		return fmt.Sprintf("!(z18n ERROR: plural form %s is empty for %s)", form, m.tag)
	}

	return m.tpl(l, s)
}
