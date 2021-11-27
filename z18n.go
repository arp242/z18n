//go:generate sh -c "go run ./cmd/gen > ./gen_data.go"

// Package z18n provides translations ("internationalisation", or "i18n").
//
// A detailed guide to get started is available in the project README at:
// https://github.com/arp242/z18n
package z18n

import (
	"os"
	"strings"

	"golang.org/x/text/language"
	"zgo.at/z18n/plural"
)

// Bundle is a "bundle" of all translations and localisations.
type Bundle struct {
	// NoHTML disabled automatic HTML escaping.
	//
	// You almost certainly *DON'T* want to set this if you're using a
	// web-app or the like or you may risk XSS security problems.
	//
	// This is provided for cases where you're not using it from a web-app,
	// in which case you probably do want to disable it. It's enabled by
	// default because printing escape codes is ugly but safe, and not
	// escaping in a web app can be disastrous.
	NoHTML bool

	// Mark messages with «msg»; useful mostly when adding z18n to an existing
	// application to visually check for untranslated strings.
	Mark bool

	// Log errors. If an unrecoverable error occurs z18n will add
	// "%(z18n ERROR [msg])" to the string. If this is non-nil then it will
	// also be called with the same string.
	Log func(string)

	defaultLang language.Tag
	tags        []language.Tag
	matcher     language.Matcher
	pluralRules plural.Rules

	// We need to store the messages on the Bundle; we can't store that on
	// the Locale as a single Locale may have multiple languages associated
	// with it.
	//
	// For example, someone may prefer Dutch, falling back to German
	// (instead of the default English). This is also how we can do
	// "inheritance" by the way: e.g. have [en-UK en-US] and only have
	// messages in en-US that need to be overridden.
	messages map[language.Tag]map[string]Msg
}

// NewBundle creates a new bundle of languages, falling back to defaultLang if a
// chosen language doesn't exist.
func NewBundle(defaultLang language.Tag) *Bundle {
	b := &Bundle{
		defaultLang: defaultLang,
		tags:        make([]language.Tag, 0, 8),
		messages:    make(map[language.Tag]map[string]Msg),
		pluralRules: plural.DefaultRules(),
	}
	b.addTag(defaultLang)
	return b
}

// AddMessages adds new messages for this language.
func (b *Bundle) AddMessages(l language.Tag, msg map[string]Msg) {
	if b.messages[l] == nil {
		b.addTag(l)
		b.messages[l] = msg
		return
	}

	for k, v := range msg {
		b.messages[l][k] = v
	}
}

// Locale gets the first matching locale for the list of languages.
func (b *Bundle) Locale(langs ...string) *Locale {
	l := &Locale{bundle: b, tags: make([]language.Tag, 0, len(langs)+1)}
	for _, lang := range langs {
		t, _, err := language.ParseAcceptLanguage(lang)
		if err != nil {
			continue
		}
		l.tags = append(l.tags, t...)
	}
	l.tags = append(l.tags, b.defaultLang)

	for _, t := range l.tags {
		if mapping[t] != nil {
			l.localize = mapping[t]
			break
		}
		// Try again with parent
		if t = t.Parent(); mapping[t] != nil {
			l.localize = mapping[t]
			break
		}
	}
	if l.localize == nil {
		panic("bug: l.localize is nil")
	}

	return l
}

// LocaleFromEnv creates a new Locale from the LANG and LC_* environment
// variables.
//
// Variables currently used:
//
//   LANGUAGES       Colon-separated list of locales
//   LANG            Default language (one value)
//   LC_COLLATE      Sorting collation
//   LC_MESSAGES     Which messages to use
//   LC_NUMERIC      Number formating
//   LC_TIME         Date/time formatting
//   LC_ALL          All of the above.
//
// Lookup order is LC_ALL, LC_*, LANGUAGES (for LC_MESSAGES), and then LANG. The
// first match wins. So if LC_ALL is set that will just override everything and
// none of the other values matter. Generally speaking you want to set only
// LANG, and maybe one of the LC_* variables if you want to override something
// specific.
//
// TODO: LC_* variables outside of LC_ALL don't do anything; we need to add
// options to set these values individually.
func (b *Bundle) LocaleFromEnv() *Locale {
	if e, ok := getEnv("LC_ALL"); ok { // LC_ALL overrides everything.
		return b.Locale(e)
	}

	var l *Locale
	if e, ok := getEnv("LANGUAGES"); ok {
		l = b.Locale(strings.Split(e, ":")...)
	} else if e, ok := getEnv("LANG"); ok {
		l = b.Locale(e)
	} else {
		l = b.Locale()
	}

	// if e, ok := getEnv("LC_COLLATE"); ok {
	// }
	// if e, ok := getEnv("LC_NUMERIC"); ok {
	// }
	// if e, ok := getEnv("LC_TIME"); ok {
	// }
	// if e, ok := getEnv("LC_MESSAGES"); ok {
	// }

	return l
}

func getEnv(name string) (string, bool) {
	l, ok := os.LookupEnv(name)
	if !ok {
		return "", false
	}
	l = strings.SplitN(l, ".", 2)[0] // Remove ".UTF-8" or ".ASCII" encoding
	if l == "" || l == "C" {         // We can't do anything with this.
		return "", false
	}
	return l, true
}

func (b *Bundle) addTag(tag language.Tag) {
	for _, t := range b.tags {
		if t == tag {
			return
		}
	}
	b.tags = append(b.tags, tag)
	b.matcher = language.NewMatcher(b.tags)
}
