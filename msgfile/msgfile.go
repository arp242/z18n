// Package msgfile manages message files.
package msgfile

import (
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"zgo.at/errors"
	"zgo.at/zstd/zfilepath"
	"zgo.at/zstd/zstring"
)

const (
	UpdatedRemoved = "removed"
	UpdatedChanged = "changed"
)

type (
	// File is a translation file.
	File struct {
		// Is this a template file?
		Template bool `toml:"template,omitempty" json:"template,omitempty"`

		// Don't update with "z18n update"
		NoUpdate bool `toml:"no-update,omitempty" json:"no_update,omitempty"`

		// Updated date.
		Generated time.Time `toml:"generated,omitempty" json:"generated,omitempty"`

		// Language this is for.
		Language string `toml:"language" json:"language"`

		// Maintainer(s) of this file.
		Maintainers []string `toml:"maintainers" json:"maintainers"`

		// Comments for translators.
		Comments string `toml:"comments" json:"comments"`

		// CLI options; only if template=true.
		Options map[string]interface{} `toml:"options,omitempty" json:"options,omitempty"`

		// Translate strings.
		Strings Entries `toml:"strings" json:"strings"`

		// File path; used internally.
		Path string `toml:"-" json:"-"`
	}

	// Entries are a list of translatable entries.
	Entries map[string]Entry

	// Entry is a single translatable entry.
	Entry struct {
		ID      string   `toml:"id" json:"id"`                               // Translation ID.
		Loc     []string `toml:"loc" json:"loc"`                             // Locations in code.
		Context string   `toml:"context,omitempty" json:"context,omitempty"` // Translation context.
		Updated string   `toml:"updated,omitempty" json:"updated,omitempty"` // Updated?

		// Messages.
		Default string `toml:"default,omitempty" json:"default,omitempty"`
		Zero    string `toml:"zero,omitempty" json:"zero,omitempty"`
		One     string `toml:"one,omitempty" json:"one,omitempty"`
		Two     string `toml:"two,omitempty" json:"two,omitempty"`
		Few     string `toml:"few,omitempty" json:"few,omitempty"`
		Many    string `toml:"many,omitempty" json:"many,omitempty"`
	}
)

// ReadFile reads a message file.
func ReadFile(fsys fs.FS, path string) (File, error) {
	fp, err := fsys.Open(path)
	if err != nil {
		return File{}, errors.Wrap(err, "msgfile.ReadFile")
	}
	defer fp.Close()

	_, ext := zfilepath.SplitExt(path)
	var f File
	switch ext {
	default:
		return File{}, errors.Errorf("msgfile.ReadFile: unknown file type %q in %q", ext, path)
	case "toml":
		_, err = toml.DecodeReader(fp, &f)
	}
	if err != nil {
		return f, errors.Wrap(err, "msgfile.ReadFile")
	}

	f.Path = path
	return f, nil
}

// FromTOML reads a message file from a TOML string.
func FromTOML(t string) (File, error) {
	var f File
	_, err := toml.Decode(t, &f)
	return f, errors.Wrap(err, "msgfile.ReadFile")
}

// WriteTo writes out a message file.
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

// Equal reports if all entries in cmp are equal.
func (e Entries) Equal(cmp Entries) bool {
	return reflect.DeepEqual(e, cmp)
}

// Untranslated reports the number of untranslated messages.
func (e Entries) Untranslated() int {
	var n int
	for _, ee := range e {
		if ee.Empty() {
			n++
		}
	}
	return n
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

// Empty reports if all messages are empty.
func (e Entry) Empty() bool {
	return e.Default == "" && e.Zero == "" && e.One == "" && e.Two == "" && e.Few == "" && e.Many == ""
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
