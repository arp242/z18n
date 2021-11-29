package z18n

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"golang.org/x/text/language"
	"zgo.at/z18n/internal"
	"zgo.at/z18n/msgfile"
	"zgo.at/zstd/zio"
	"zgo.at/zstd/ztest"
)

func TestWorkflow(t *testing.T) {
	// Setup.
	tmp := t.TempDir()
	if err := zio.CopyTree("./testdata/testapp", tmp, nil); err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(tmp, "i18n")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create i18n dir contents
	{
		tpl, err := internal.Find([]string{tmp}, "en-GB", "toml",
			[]string{"T"}, []string{"t", ".T"}, []string{"gotxt"})
		if err != nil {
			t.Fatalf("%s: %s", err, tpl)
		}
		if err := os.WriteFile(filepath.Join(dir, "/template.toml"), []byte(tpl), 0644); err != nil {
			t.Fatal(err)
		}
		if err := internal.New("en-GB", dir); err != nil {
			t.Fatal(err)
		}
		if err := internal.New("nl-NL", dir); err != nil {
			t.Fatal(err)
		}
		if err := internal.New("nl-NL", dir); err == nil {
			t.Fatal("creating translation file second time didn't give an error")
		}
	}

	// Set actual translation strings.
	tpl, err := msgfile.ReadFile(filepath.Join(dir, "template.toml"))
	if err != nil {
		t.Fatal(err)
	}
	en, err := msgfile.ReadFile(filepath.Join(dir, "en-GB.toml"))
	if err != nil {
		t.Fatal(err)
	}
	nl, err := msgfile.ReadFile(filepath.Join(dir, "nl-NL.toml"))
	if err != nil {
		t.Fatal(err)
	}

	{
		for k, v := range en.Strings {
			v.Default = "English translation of " + k
			en.Strings[k] = v
		}
		for k, v := range nl.Strings {
			v.Default = "Nederlandse vertaling van " + k
			nl.Strings[k] = v
		}

		if err := en.WriteTo(en.Path); err != nil {
			t.Fatal(err)
		}
		if err := nl.WriteTo(nl.Path); err != nil {
			t.Fatal(err)
		}
	}

	// Make some changes.
	{
		err := os.WriteFile(filepath.Join(tmp, "new.go"), []byte(`package testapp
			var (
				newStr1 = T("newstr1")
				newStr2 = T("newstr2|newstr2 with default")
			)
		`), 0644)
		if err != nil {
			t.Fatal(err)
		}

		subFile(t, filepath.Join(tmp, "testapp.go"),
			`var _ = T("package-var")`+"\n", ``,
			`T("message-id")`, `T("message-id|now with a default")`,
		)

		err = internal.Update(dir)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Test resulting TOML files match.
	{
		wantTpl := setID(msgfile.Entries{
			"tpl-1":                  {Loc: []string{"testapp.gotxt:1"}},
			"tpl-2":                  {Loc: []string{"testapp.gotxt:4"}, Context: "context for tpl-2."},
			"newstr1":                {Loc: []string{"new.go:3"}},
			"package-var-in-block-1": {Loc: []string{"testapp.go:7", "testapp.go:8"}},
			"package-var-in-block-2": {Loc: []string{"testapp.go:9"}},
			"package-var-in-map-1":   {Loc: []string{"testapp.go:13"}},
			"package-var-in-map-2":   {Loc: []string{"testapp.go:14"}},
			"with-context-1":         {Loc: []string{"testapp.go:22"}, Context: "context for 1."},
			"with-context-2":         {Loc: []string{"testapp.go:26"}, Context: "context for 2."},
			"message-id": {
				Loc:     []string{"testapp.go:18"},
				Default: "now with a default",
			},
			"newstr2": {
				Loc:     []string{"new.go:4"},
				Default: "newstr2 with default",
			},
		})

		wantF := map[string]msgfile.File{
			"template.toml": {
				Template:  true,
				Language:  "en-GB",
				Generated: tpl.Generated,
				Options:   tpl.Options,
				Strings:   wantTpl,
			},
			"en-GB.toml": {
				Language:  "en-GB",
				Generated: en.Generated,
				Strings: merge(wantTpl, map[string]map[string]string{
					"tpl-1":                  {"default": "English translation of tpl-1"},
					"tpl-2":                  {"default": "English translation of tpl-2"},
					"message-id":             {"default": "English translation of message-id"},
					"newstr2":                {"default": "newstr2 with default"},
					"package-var-in-block-1": {"default": "English translation of package-var-in-block-1"},
					"package-var-in-block-2": {"default": "English translation of package-var-in-block-2"},
					"package-var-in-map-1":   {"default": "English translation of package-var-in-map-1"},
					"package-var-in-map-2":   {"default": "English translation of package-var-in-map-2"},
					"with-context-1":         {"default": "English translation of with-context-1"},
					"with-context-2":         {"default": "English translation of with-context-2"},
					"package-var": {
						"id":      "package-var",
						"unused":  "",
						"loc":     "testapp.go:5",
						"default": "English translation of package-var",
					},
				}),
			},
			"nl-NL.toml": {
				Language:  "nl-NL",
				Generated: nl.Generated,
				Strings: merge(wantTpl, map[string]map[string]string{
					"tpl-1":                  {"default": "Nederlandse vertaling van tpl-1"},
					"tpl-2":                  {"default": "Nederlandse vertaling van tpl-2"},
					"message-id":             {"default": "Nederlandse vertaling van message-id"},
					"newstr2":                {"default": ""},
					"package-var-in-block-1": {"default": "Nederlandse vertaling van package-var-in-block-1"},
					"package-var-in-block-2": {"default": "Nederlandse vertaling van package-var-in-block-2"},
					"package-var-in-map-1":   {"default": "Nederlandse vertaling van package-var-in-map-1"},
					"package-var-in-map-2":   {"default": "Nederlandse vertaling van package-var-in-map-2"},
					"with-context-1":         {"default": "Nederlandse vertaling van with-context-1"},
					"with-context-2":         {"default": "Nederlandse vertaling van with-context-2"},
					"package-var": {
						"id":      "package-var",
						"unused":  "",
						"loc":     "testapp.go:5",
						"default": "Nederlandse vertaling van package-var",
					},
				}),
			},
		}
		want := make([][2]string, 0, len(wantF))
		for k, v := range wantF {
			data, err := v.TOML()
			if err != nil {
				t.Fatal(err)
			}
			want = append(want, [2]string{k, data})
		}
		sort.Slice(want, func(i, j int) bool { return want[i][0] < want[j][0] })

		for i, v := range readDir(t, dir) {
			h := v[0] + "\n" +
				"\t" + strings.ReplaceAll(strings.TrimSpace(v[1]), "\n", "\n\t") + "\n\n"
			w := want[i][0] + "\n" +
				"\t" + strings.ReplaceAll(strings.TrimSpace(want[i][1]), "\n", "\n\t") + "\n\n"

			if d := ztest.Diff(h, w); d != "" {
				t.Errorf("\n%q differs:%s", v[0], d)
			}

		}
	}

	// Test if the locale works and matches.
	// TODO
	{
		b := NewBundle(language.BritishEnglish)
		err := b.ReadMessagesDir(os.DirFS(dir), "*.toml")
		if err != nil {
			t.Fatal(err)
		}
		//
		// en := b.Locale("en"), nl  b.Locale("nl")
		// fmt.Println(en.T("package-var"))
		// fmt.Println(nl.T("package-var"))
	}
}

func merge(dst msgfile.Entries, src map[string]map[string]string) msgfile.Entries {
	cp := make(msgfile.Entries)
	for k, v := range dst {
		cp[k] = v
	}

	for k, v := range src {
		if v == nil {
			delete(cp, k)
			continue
		}

		d := cp[k]
		for kk, vv := range v {
			switch kk {
			case "id":
				d.ID = vv
			case "loc":
				d.Loc = []string{vv}
			case "default":
				d.Default = vv
			case "unused":
				d.Updated = msgfile.UpdatedRemoved
			case "changed":
				d.Updated = msgfile.UpdatedChanged
			}
		}
		cp[k] = d
	}
	return cp
}

func setID(m msgfile.Entries) msgfile.Entries {
	for k, v := range m {
		v.ID = k
		m[k] = v
	}
	return m
}

func subFile(t *testing.T, path string, subs ...string) {
	if len(subs)%2 != 0 {
		t.Fatal("wrong args")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(subs); i += 2 {
		data = bytes.ReplaceAll(data, []byte(subs[i]), []byte(subs[i+1]))
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		t.Fatal(err)
	}
}

func readDir(t *testing.T, path string) [][2]string {
	ls, err := os.ReadDir(path)
	if err != nil {
		t.Fatal(err)
	}

	r := make([][2]string, 0, len(ls))
	for _, l := range ls {
		if l.IsDir() {
			continue
		}

		d, err := os.ReadFile(filepath.Join(path, l.Name()))
		if err != nil {
			t.Fatal(err)
		}
		r = append(r, [2]string{l.Name(), string(d)})
	}
	return r
}
