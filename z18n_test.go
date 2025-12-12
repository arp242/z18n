package z18n

import (
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"golang.org/x/text/language"
)

func mkbundle() *Bundle {
	b := NewBundle(language.MustParse("en-NZ"))
	b.AddMessages(language.MustParse("en-NZ"), map[string]Msg{
		"hello":     Msg{Default: "Hello"},
		"hello-loc": Msg{Default: "Hello, %(loc)!"},
		"btn":       Msg{Default: "%[Button]"},
		"btn2":      Msg{Default: "Hello %[%btn1 Button], %[%btn2 another] XX"},
		"btn3":      Msg{Default: "%[%btn Button %(var)]"},
		"btn4":      Msg{Default: "%[%btn Send]"},
		"btn5":      Msg{Default: "phone me: %[phone]"},
		"btn6":      Msg{Default: "phone me: %(phone)"},
		"ants!": Msg{
			One:     "Help, I've got an ant in my trousers!",
			Default: "Help, I've got %(n) ants in my trousers!",
		},
	})
	b.AddMessages(language.MustParse("nl-NL"), map[string]Msg{
		"hello":     Msg{Default: "Hallo"},
		"hello-loc": Msg{Default: "Hallo, %(loc)!"},
		"btn":       Msg{Default: "%[Knop]"},
		"btn2":      Msg{Default: "Hallo %[%btn1 Knop], %[%btn2 nog een] XX"},
		"btn3":      Msg{Default: "%[%btn Knop %(var)]"},
		"btn4":      Msg{Default: "%[%btn Verstuur]"},
		"btn5":      Msg{Default: "bel me: %[phone]"},
		"btn6":      Msg{Default: "bel me: %(phone)"},
		"ants!": Msg{
			One:     "Help, ik heb een mier in m'n broek!",
			Default: "Help, ik heb %(n) mieren in m'n broek!",
		},
	})
	return b
}

func TestT(t *testing.T) {
	nz, err := time.LoadLocation("Pacific/Auckland")
	if err != nil {
		panic(err)
	}
	tUTC := time.Date(1985, 6, 18, 15, 42, 30, 123, time.UTC)
	tNZ := time.Date(1985, 6, 18, 15, 42, 30, 123, nz)

	tests := []struct {
		name   string
		id     string
		data   []any
		wantEN string
		wantNL string
	}{
		{"empty string",
			"", nil, "", ""},
		{"basic string",
			"hello", nil, "Hello", "Hallo"},
		{"default message",
			"id|Default msg", nil, "Default msg", "Default msg"},
		{"unknown key",
			"unknown", nil, "unknown", "unknown"},

		// Variables
		{"variable",
			"hello-loc", []any{"z18n"}, "Hello, z18n!", "Hallo, z18n!"},
		{"variable in default msg",
			"|%(var)", []any{"xx"}, "xx", "xx"},
		{"variable in id",
			"%(var)", nil, "%(var)", "%(var)"},
		{"two variables",
			"|%(var) %(bar)", []any{P{"var": "xx", "bar": "yy"}}, "xx yy", "xx yy"},

		// Functions, HTML escape
		{"upper", "|a%(var upper asd)!", []any{"xx"}, "aXX!", "aXX!"},
		{"html", "|%(var)", []any{"<script>"}, "&lt;script&gt;", "&lt;script&gt;"},
		{"html-raw", "|%(var raw)", []any{"<script>"}, "<script>", "<script>"},

		// Ints
		{"int", "|%(n)", []any{2}, "2", "2"},
		{"int", "|%(n)", []any{2000}, "2,000", "2.000"},
		{"int", "|%(n)", []any{2000000}, "2,000,000", "2.000.000"},
		{"int", "|%(n raw)", []any{2000000}, "2000000", "2000000"},

		// Floats
		{"float", "|%(n)", []any{2000.1565}, "2,000.157", "2.000,157"},
		{"float", "|%(n raw)", []any{2000.1565}, "2000.157", "2000,157"},
		{"float", "|%(n 1)", []any{2000.1565}, "2,000.2", "2.000,2"},
		{"float", "|%(n 2 raw)", []any{2000.1565}, "2000.16", "2000,16"},
		{"float", "|%(n 10)", []any{2000.1}, "2,000.1", "2.000,1"},

		// Time
		{"time", "|%(t)", []any{tUTC}, "18/06/1985, 3:42:30 PM", "18 jun. 1985 15:42:30"},
		{"time", "|%(t)", []any{tNZ}, "18/06/1985, 3:42:30 PM", "18 jun. 1985 15:42:30"},
		{"time", "|%(t full)", []any{tUTC},
			"Wednesday, 18 June 1985 at 3:42:30 PM UTC +0000",
			"woensdag 18 juni 1985 om 15:42:30 UTC +0000"},
		{"time", "|%(t full)", []any{tNZ},
			"Wednesday, 18 June 1985 at 3:42:30 PM NZST +1200",
			"woensdag 18 juni 1985 om 15:42:30 NZST +1200"},

		{"time", "|%(t date)", []any{tUTC}, "18/06/1985", "18 jun. 1985"},
		{"time", "|%(t date short)", []any{tUTC}, "18/06/85", "18-06-1985"},
		{"time", "|%(t time)", []any{tUTC}, "3:42:30 PM", "15:42:30"},
		{"time", "|%(t time short)", []any{tUTC}, "3:42 PM", "15:42"},

		{"time", "|%(t time Mon Jan)", []any{tUTC}, "Wed Jun", "wo jun."},

		// Tags
		{"tags",
			"btn|%[Button]", []any{Tag("a", `href="/foo"`)},
			`<a href="/foo">Button</a>`, `<a href="/foo">Knop</a>`},
		{"tags", "btn2", []any{P{
			"btn1": Tag("a", `href="/btn1"`),
			"btn2": Tag("a", `href="/btn2"`),
		}},
			`Hello <a href="/btn1">Button</a>, <a href="/btn2">another</a> XX`,
			`Hallo <a href="/btn1">Knop</a>, <a href="/btn2">nog een</a> XX`,
		},
		{"var inside html", "btn3|%[%btn Button %(var)]", []any{P{
			"var": "X",
			"btn": Tag("a", `href="/foo"`),
		}},
			`<a href="/foo">Button X</a>`,
			`<a href="/foo">Knop X</a>`},

		// Tag with innerHTML
		{"tag with content", "btn5|phone me: %[phone]", []any{Tag("a", `href="tel:555"`, `5-5-5`)},
			`phone me: <a href="tel:555">5-5-5</a>`,
			`bel me: <a href="tel:555">5-5-5</a>`},
		{"tag with content", "btn6|phone me: %(phone)", []any{Tag("a", `href="tel:555"`, `5-5-5`)},
			`phone me: <a href="tel:555">5-5-5</a>`,
			`bel me: <a href="tel:555">5-5-5</a>`},

		// Plurals
		{"ants!", "ants!", []any{N(0)}, "Help, I've got 0 ants in my trousers!", "Help, ik heb 0 mieren in m'n broek!"},
		{"ants!", "ants!", []any{N(1)}, "Help, I've got an ant in my trousers!", "Help, ik heb een mier in m'n broek!"},
		{"ants!", "ants!", []any{N(2)}, "Help, I've got 2 ants in my trousers!", "Help, ik heb 2 mieren in m'n broek!"},

		// TODO: error reporting could be a bit better on this. Also have z18n
		// CLI catch this.
		{"ants!", "ants!", []any{},
			`Help, I've got !(z18n ERROR: no value for variable "n") ants in my trousers!`,
			`Help, ik heb !(z18n ERROR: no value for variable "n") mieren in m'n broek!`},
		{"plural", "hello", []any{N(1)},
			"!(z18n ERROR: plural form one is empty for en-NZ)",
			"!(z18n ERROR: plural form one is empty for nl-NL)"},

		// Errors
		{"extra params", "hello", []any{"X"}, "Hello", "Hallo"},
		{"no vars", "hello-loc", nil,
			`Hello, !(z18n ERROR: no value for variable "loc")!`,
			`Hallo, !(z18n ERROR: no value for variable "loc")!`},
	}

	b := mkbundle()
	en := b.Locale("en-NZ")
	nl := b.Locale("nl-NL")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			haveEN := en.T(tt.id, tt.data...)
			if haveEN != tt.wantEN {
				t.Errorf("English wrong\nhave: %s\nwant: %s", haveEN, tt.wantEN)
			}
			haveNL := nl.T(tt.id, tt.data...)
			if haveNL != tt.wantNL {
				t.Errorf("Dutch wrong\nhave: %s\nwant: %s", haveNL, tt.wantNL)
			}
		})
	}
}

func TestL10n(t *testing.T) {
	tests := []struct {
		in string
	}{
		{""},
	}

	b := NewBundle(language.MustParse("en-US"))
	b.AddMessages(language.MustParse("en-US"), map[string]Msg{
		"order": Msg{
			Default: "Your order of %(n) gophers will arrive on %(d).",
		},
	})
	b.AddMessages(language.MustParse("en-NZ"), map[string]Msg{
		"order": Msg{
			Default: "Your order of %(n) gophers will arrive on %(d).",
		},
	})
	b.AddMessages(language.MustParse("nl-NL"), map[string]Msg{
		"order": Msg{
			Default: "Uw bestelling van %(n) gophers wordt op %(d) geleverd",
		},
	})

	lUS := b.Locale("en-US")
	lNZ := b.Locale("en-NZ")
	lNL := b.Locale("nl-NL")

	data := P{
		"n": 153_416_164_1,
		"d": time.Date(2021, 06, 18, 0, 0, 0, 0, time.UTC),
	}
	_, _, _, _ = lUS, lNZ, lNL, data
	// fmt.Println(lUS.T("order", data))
	// fmt.Println(lNZ.T("order", data))
	// fmt.Println(lNL.T("order", data))

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			_ = tt

			// if have != tt.want {
			// 	t.Errorf("\nhave: %q\nwant: %q", have, tt.want)
			// }
			// if !reflect.DeepEqual(have, tt.want) {
			// 	t.Errorf("\nhave: %#v\nwant: %#v", have, tt.want)
			// }
			// if d := ztest.Diff(have, tt.want); d != "" {
			// 	t.Errorf(d)
			// }
		})
	}
}

// Cascade rules:
//
//  1. User has one or more prefered languages, check every one of thise.
//     First match wins.
//     e.g. "nl-NL" "en-US", "en-GB"
//
//  2. No matches? Check the parents.
//     e.g. "nl" "en"
func TestCascade(t *testing.T) {
	en := language.MustParse("en")
	enUS := language.MustParse("en-US")
	enNZ := language.MustParse("en-NZ")
	enGB := language.MustParse("en-GB")

	b := NewBundle(en)
	b.AddMessages(en, map[string]Msg{
		"en": {Default: "en from en"},
	})
	b.AddMessages(enNZ, map[string]Msg{
		"en":   {Default: "en from en-NZ"},
		"enNZ": {Default: "enNZ from en-NZ"},
	})
	// TODO: if I don't do this than it will print numbers etc. in "en".
	b.AddMessages(language.Dutch, map[string]Msg{})

	lEN := b.Locale("en")
	lNZ := b.Locale("en-NZ")
	lUS := b.Locale("en-US", "en-NZ") // TODO: Shouldn't this get en-NZ msg?

	//fmt.Println(
	//	lUS.T("test|The genetic test I did on %(t date full) showed I'm %(f 0)% platypus; a trait %(n) people share", P{
	//		"n": 51341,
	//		"f": 13.42,
	//		"t": time.Now(),
	//	}))

	// fmt.Printf("en → %q -> %q\n", lEN.T("en"), lEN.T("enNZ"))
	// fmt.Printf("nz → %q -> %q\n", lNZ.T("en"), lNZ.T("enNZ"))
	// fmt.Printf("us → %q -> %q\n", lUS.T("en"), lUS.T("enNZ"))

	_, _ = enUS, enGB
	_, _, _ = lEN, lNZ, lUS

	// 2021 ጁን 12, ቅዳሜ 3:46:44 PM WITA +0800
	//
	// TODO: not entirely sure how to get ethiopic numerals out of here; it
	// looks like x/text doesn't support this (?) I also can't easily seem to
	// find this in the CLDR data either(?)
	// Doesn't seem to be use a huge issue, as e.g. the Ethiopian Wikipedia uses
	// Latin/Arabic numerals throughout (as do the Hebrew, Greek, Arabic
	// Wikipedia), but would be nice to support other numbering systems too.
	//
	// tag := language.MustParse("am-ET")
	// b.AddMessages(tag, map[string]Msg{})
	// l := b.Locale("am-ET")
	// fmt.Println(l.T("|%(t full)", time.Now()))

	// 123, in latin
	// message.NewPrinter(tag).Printf("%f\n", number.Decimal(123))
}

func TestLocaleFromEnv(t *testing.T) {
	b := mkbundle()
	env := os.Environ()
	restore := func() {
		os.Clearenv()
		for _, e := range env {
			k, v, _ := strings.Cut(e, "=")
			err := os.Setenv(k, v)
			if err != nil {
				panic(err)
			}
		}
	}

	nz, nl, id := language.MustParse("en-NZ"), language.MustParse("nl-NL"), language.MustParse("id-ID")

	tests := []struct {
		env  map[string]string
		want []language.Tag
	}{
		{map[string]string{}, []language.Tag{nz}},
		{map[string]string{"LANG": "nl-NL"}, []language.Tag{nl, nz}},
		{map[string]string{"LC_ALL": "nl-NL"}, []language.Tag{nl, nz}},
		{map[string]string{"LC_ALL": "id-ID", "LC_MESSAGES": "nl-NL"}, []language.Tag{id, nz}},
		{map[string]string{"LANGUAGES": "nl-NL:id-ID"}, []language.Tag{nl, id, nz}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			defer restore()

			os.Clearenv()
			for k, v := range tt.env {
				os.Setenv(k, v)
			}

			l := b.LocaleFromEnv()
			if !reflect.DeepEqual(l.tags, tt.want) {
				t.Errorf("\nhave: %s\nwant: %s", l.tags, tt.want)
			}
		})
	}
}

func BenchmarkNew(b *testing.B) {
	tag := language.MustParse("en-NZ")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = NewBundle(tag)
	}
}

func BenchmarkLocale(b *testing.B) {
	bundle := mkbundle()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = bundle.Locale("nl-NL")
	}
}

func BenchmarkT(b *testing.B) {
	bundle := mkbundle()
	nl := bundle.Locale("nl-NL")
	b.ResetTimer()

	b.Run("string", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = nl.T("hello", "a")
		}
	})
	b.Run("one-variable", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = nl.T("hello-loc", "a")
		}
	})
	b.Run("one-tag", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = nl.T("btn", Tag("a", `href="/foo"`))
		}
	})
}
