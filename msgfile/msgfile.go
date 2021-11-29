// Package msgfile manages message files.
package msgfile

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"zgo.at/errors"
	"zgo.at/zstd/zstring"
)

const (
	UpdatedRemoved = "removed"
	UpdatedChanged = "changed"
)

type (
	// File is a translation file.
	File struct {
		Template    bool                   `toml:"template"`    // Is this a template file?
		NoUpdate    bool                   `toml:"no-update"`   // Don't update with "z18n update"
		Generated   time.Time              `toml:"generated"`   // Updated date.
		Language    string                 `toml:"language"`    // Language this is for.
		Maintainers []string               `toml:"maintainers"` // Maintainer(s) of this file.
		Comments    string                 `toml:"comments"`    // Comments for translators.
		Options     map[string]interface{} `toml:"options"`     // CLI options; only if template=true.

		Strings Entries `toml:"-"` // Translate strings.
		Path    string  `toml:"-"` // File path; used internally.
	}

	// Entries are a list of translatable entries.
	Entries map[string]Entry

	// Entry is a single translatable entry.
	Entry struct {
		ID      string   `toml:"-"`       // Translation ID.
		Loc     []string `toml:"loc"`     // Locations in code.
		Context string   `toml:"context"` // Translation context.
		Updated string   `toml:"updated"` // Updated?

		// Messages.
		Default string `toml:"default"`
		Zero    string `toml:"zero"`
		One     string `toml:"one"`
		Two     string `toml:"two"`
		Few     string `toml:"few"`
		Many    string `toml:"many"`
	}
)

func ReadFile(path string) (File, error) {
	var f File
	_, err := toml.DecodeFile(path, &f)
	f.Path = path
	return f, errors.Wrap(err, "msgfile.ReadFile")
}

func (f File) WriteTo(path string) error {
	t, err := f.TOML()
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(t), 0644)
}

// Merge strings from src in to this Entries.
func (e Entries) Merge(src Entries) {
	for k, v := range src {
		have, ok := e[k]
		if ok {
			have.Loc = append(have.Loc, v.Loc...)
		} else {
			have = v
		}
		e[k] = v
	}
}

// Sorted returns a list of entries sorted by ID.
func (e Entries) Sorted() []Entry {
	ord := make([]Entry, 0, len(e))
	for _, f := range e {
		ord = append(ord, f)
	}
	sort.Slice(ord, func(i, j int) bool { return ord[i].ID < ord[j].ID })
	return ord
}

var reIndent = regexp.MustCompile(`\n[ \t]+`)

// Normalize the indent: remove leading/trailing whitespace and remove all
// whitespace at the start of every line. This allows doing stuff like:
//
//    z18n.T(`
//          Some message that's indented nicely in Go
//          But you don't really want this indent to show up.
//    `)
func normalizeMessage(msg string) string {
	msg = strings.TrimSpace(msg)
	msg = reIndent.ReplaceAllString(msg, "\n")
	return msg
}

type param struct {
	kind string   // map, plural, literal, tag
	keys []string // only for map
}

// TODO: need to parse the params from both Go and Template and convert to
// something we can pass here, so we don't need to duplicate the logic.
//
//   typeCheck("foo/bar", "hello %(world)", MapParam{keys: []string{...}, types: LiteralParam{}, TagParam{}}, PluralParam{})
//   typeCheck("foo/bar", "hello %(world)", LiteralParam{})
//   typeCheck("foo/bar", "hello %[world asd]", TagParam{})
//
// 1. Grep the %() and %[] out of there
// 2. Check syntax of %[word text]
// 3. Check if it matches with that we expect in params
func typeCheck(id, msg string, params ...param) error {
	if msg == "" {
		return nil
	}

	var (
		tags                      = zstring.IndexPairs(msg, "%[", "]")
		vars                      = zstring.IndexPairs(msg, "%(", ")")
		mapKeys                   []string
		nPlural, nMap, nLit, nTag int
		//total                     = len(tags) + len(vars)
	)
	for _, p := range params {
		switch p.kind {
		case "plural":
			nPlural++
		case "map":
			nMap++
			mapKeys = p.keys
		case "literal":
			nLit++
		case "tag":
			nTag++
		}
	}

	errs := errors.NewGroup(10)
	if nPlural > 1 {
		errs.Append(errors.New("more than one plural parameter"))
	}
	if nMap > 1 {
		errs.Append(errors.New("more than one map parameter"))
	}
	if nLit > 1 {
		errs.Append(errors.New("more than one literal parameter"))
	}
	if nMap > 0 && nLit > 0 {
		errs.Append(errors.New("both literal and map parameter"))
	}
	if nTag != len(tags) {
		errs.Append(fmt.Errorf("wrong number of tag parameters; %d in string but %d parameters", len(tags), nTag))
	}

	// TODO: also check if there's stuff in the map that's not used.
	if len(vars) > 1 {
		for _, p := range vars {
			start, end := p[0], p[1]
			varname := msg[start+2 : end]
			if !zstring.Contains(mapKeys, varname) {
				errs.Append(fmt.Errorf("not in map: %q", varname))
			}
		}
	}

	// TODO: check tag names
	//for _, p := range tags {
	//	start, end := p[0], p[1]
	//	text := str[start+2 : end]
	//	varname, text := zstring.Split2(text, " ")

	return errs.ErrorOrNil()
}
