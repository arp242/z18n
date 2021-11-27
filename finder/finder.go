// Package finder finds translatable strings.
package finder

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/BurntSushi/toml"
	"golang.org/x/tools/go/packages"
	"zgo.at/errors"
	"zgo.at/zstd/zstring"
	"zgo.at/ztpl/parse"
)

type (
	// Entry is a single translatable entry.
	Entry struct {
		ID      string   `toml:"-"`
		Loc     []string `toml:"loc"`
		Default string   `toml:"default"`
		Removed bool     `toml:"-"`
	}

	// Entries are a list of translatable entries.
	Entries map[string]Entry

	// File is a translation file.
	File struct {
		BaseFile    bool                   `toml:"base-file"`
		Generated   time.Time              `toml:"generated"`
		Language    string                 `toml:"language"`
		Maintainers []string               `toml:"maintainers"`
		Options     map[string]interface{} `toml:"options"`
		Strings     Entries                `toml:"-"`
		Path        string                 `toml:"-"`
	}
)

// List all entries in this file.
func (f File) List() (string, error) {
	return f.Strings.list()
}

func (f *File) UnmarshalTOML(d interface{}) error {
	m := d.(map[string]interface{})
	meta := m["__meta__"]
	delete(m, "__meta__")

	for k, v := range meta.(map[string]interface{}) {
		switch k {
		case "base-file":
			f.BaseFile = v.(bool)
		case "generated":
			f.Generated = v.(time.Time)
		case "language":
			f.Language = v.(string)
		case "maintainers":
			vv := v.([]interface{})
			set := make([]string, 0, len(vv))
			for _, s := range vv {
				set = append(set, s.(string))
			}
			f.Maintainers = set
		case "options":
			f.Options = v.(map[string]interface{})
		}
	}

	f.Strings = make(Entries)
	for k, v := range m {
		vv := v.(map[string]interface{})

		var loc []string
		if vv["loc"] != nil {
			ll := vv["loc"].([]interface{})
			loc = make([]string, 0, len(ll))
			for _, s := range ll {
				loc = append(loc, s.(string))
			}
		}

		e := Entry{
			ID:      k,
			Default: vv["default"].(string),
			Loc:     loc,
		}
		f.Strings[k] = e
	}

	return nil
}

// TOML formats all entries as TOML.
//
// TODO: using the TOML encoder would be better, but it doesn't support the
// options we want for formatting things (yet).
func (f File) TOML() (string, error) {
	meta := map[string]interface{}{
		"generated": f.Generated,
	}
	if f.BaseFile {
		meta["base-file"] = f.BaseFile
		meta["options"] = f.Options
	} else {
		meta["language"] = f.Language
		meta["maintainers"] = f.Maintainers
	}

	b := new(bytes.Buffer)
	err := toml.NewEncoder(b).Encode(map[string]interface{}{"__meta__": meta})
	if err != nil {
		return "", err
	}

	t, err := f.Strings.toml()
	if err != nil {
		return "", err
	}

	return b.String() + "\n" + t, nil
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

// List all messages, one per line, without context.
func (e Entries) list() (string, error) {
	max := 0
	for _, x := range e {
		if l := utf8.RuneCountInString(x.ID); l > max {
			max = l
		}
	}

	b := new(strings.Builder)
	for _, x := range e.Sorted() {
		fmt.Fprintf(b, "%s %s %q\n", x.ID, strings.Repeat(" ", max-utf8.RuneCountInString(x.ID)), x.Default)
	}
	return b.String(), nil
}

func (e Entries) toml() (string, error) {
	b := new(strings.Builder)
	// err := toml.NewEncoder(b).Encode(e)
	// return b.String(), err

	for _, x := range e.Sorted() {
		// for _, l := range x.Loc {
		// 	b.WriteString("# ")
		// 	b.WriteString(l)
		// 	b.WriteByte('\n')
		// }

		cmt := ""
		if x.Removed {
			cmt = "# "
			fmt.Fprintf(b, "# No longer used (%s)\n", time.Now().UTC().Round(time.Second))
		}

		fmt.Fprintf(b, "%s[%s]\n", cmt, tomlString(x.ID))
		switch len(x.Loc) {
		case 0: /// Should never happen, but just in case.
			fmt.Fprintf(b, "%s  loc     = []\n", cmt)
		case 1:
			fmt.Fprintf(b, "%s  loc     = [%s]\n", cmt, tomlString(x.Loc[0]))
		default:
			fmt.Fprintf(b, "%s  loc     = [\n", cmt)
			for _, l := range x.Loc {
				fmt.Fprintf(b, "%s    %s,\n", cmt, tomlString(l))
			}
			fmt.Fprintf(b, "%s  ]\n", cmt)
		}
		fmt.Fprintf(b, "%s  default = %s\n\n", cmt, tomlString(x.Default))
	}
	return b.String(), nil
}

// Go finds all translatable strings in dir.
//
// TODO: find context from comments too.
func Go(dir string, funs ...string) (Entries, error) {
	p := dir
	if p == "" || p == "." {
		p = "./..."
	}

	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedTypes | packages.NeedSyntax,
		Dir:  dir,
	}, p)
	if err != nil {
		return nil, errors.Wrap(err, "finder.Go")
	}

	var (
		found = make(Entries)
		errs  = errors.NewGroup(20)
	)
	for _, p := range pkgs {
		if len(p.Errors) > 0 {
			for _, e := range p.Errors {
				errs.Append(e)
			}
			return nil, errors.Wrap(errs.ErrorOrNil(), "finder.Go")
		}
		for _, f := range p.Syntax {
			ast.Inspect(f, func(n ast.Node) bool {
				if n == nil {
					return true
				}

				c, ok := n.(*ast.CallExpr)
				if !ok {
					return true
				}

				var name string
				switch f := c.Fun.(type) {
				case *ast.Ident:
					name = f.Name
				case *ast.SelectorExpr:
					name = f.Sel.Name
				case *ast.FuncLit, *ast.ArrayType, *ast.ParenExpr, *ast.CallExpr:
					return true // No work.
				default:
					pos := p.Fset.Position(c.Pos())
					panic(errors.Errorf("finder.Go: unhandled c.Fun type %T at %s:%d\n    %#[1]v", c.Fun, pos.Filename, pos.Line))
				}

				if !zstring.Contains(funs, name) {
					return true
				}

				if len(c.Args) < 1 {
					pos := p.Fset.Position(c.Fun.Pos())
					errs.Append(errors.Errorf("%s:%d: not enough arguments to %q",
						pos.Filename, pos.Line, name))
				}

				// Get the first string argument.
				idlit, ok := c.Args[0].(*ast.BasicLit)
				if !ok || idlit.Kind != token.STRING {
					if len(c.Args) > 1 {
						idlit, ok = c.Args[1].(*ast.BasicLit)
					}
					if !ok || idlit.Kind != token.STRING {
						return true
					}
				}

				id, def := zstring.Split2(strings.Trim(idlit.Value, quotes), "|")

				// if def != "" && errors.Append(typeCheck(id, def)) {
				//    return false
				// }

				e := Entry{ID: id, Default: normalizeMessage(def)}
				if f, ok := found[id]; ok {
					e.Loc = f.Loc
				}

				pos := p.Fset.Position(c.Pos())
				e.Loc = append(e.Loc, fmt.Sprintf("%s:%d", strings.TrimPrefix(pos.Filename, dir+"/"), pos.Line))
				found[id] = e

				return true
			})
		}
	}

	return found, errors.Wrap(errs.ErrorOrNil(), "finder.Go")
}

// Template finds all strings in Go templates.
//
// TODO: find context from comments too.
func Template(dir string, ext []string, funs ...string) (Entries, error) {
	found := make(map[string]Entry)
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !zstring.HasSuffixes(d.Name(), ext...) {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		tree, err := parse.Parse("", string(data), parse.ParseRelaxFunctions, "{{", "}}")
		if err != nil {
			return fmt.Errorf("parsing %q: %w", path, err)
		}

		parse.Visit(tree[""].Root, func(n parse.Node, _ int) bool {
			nn, ok := n.(*parse.ActionNode)
			if !ok {
				return true
			}

			if nn.Pipe == nil || len(nn.Pipe.Cmds) == 0 || len(nn.Pipe.Cmds[0].Args) < 2 {
				return false
			}

			name := ""
			switch f := nn.Pipe.Cmds[0].Args[0].(type) {
			case *parse.IdentifierNode:
				name = f.Ident
			case *parse.FieldNode:
				name = "." + strings.Join(f.Ident, ".")
			}
			if !zstring.Contains(funs, name) {
				return false
			}

			idlit, ok := nn.Pipe.Cmds[0].Args[1].(*parse.StringNode)
			if !ok || idlit.Text == "" {
				idlit, ok = nn.Pipe.Cmds[0].Args[2].(*parse.StringNode)
				if !ok || idlit.Text == "" {
					return false
				}
			}
			id, def := zstring.Split2(idlit.Text, "|")

			// if def != "" && errors.Append(typeCheck(id, def)) {
			//    return false
			// }

			e := Entry{ID: id, Default: normalizeMessage(def)}
			f, ok := found[id]
			if ok {
				e.Loc = f.Loc
			}

			e.Loc = append(e.Loc, fmt.Sprintf("%s:%d", strings.TrimPrefix(path, dir+"/"), nn.Line))
			found[id] = e
			return false
		})

		return nil
	})

	return found, err
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

var tomlQuote = strings.NewReplacer(
	"\t", "\\t",
	//"\n", "\\n", // We print this as a multi-line message, so don't escape
	"\r", "\\r",
	"\"", "\\\"",
	"\\", "\\\\")

func tomlString(s string) string {
	s = strings.TrimSpace(s)
	if strings.IndexByte(s, '\n') > -1 {
		return `"""` + "\n" + tomlQuote.Replace(s) + "\n" + `"""`
	}
	return `"` + tomlQuote.Replace(s) + `"`
}

var goRawQuote = strings.NewReplacer("`", "` + \"`\" + `")

// This has a tendency to break Vim syntax highlights so put it here... Yeah, I
// know :-/
const quotes = "\"`"
