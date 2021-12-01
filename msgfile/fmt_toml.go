package msgfile

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

// Custom unmarshal to treat __meta__ as special.
func (f *File) UnmarshalTOML(d interface{}) error {
	m := d.(map[string]interface{})
	meta := m["__meta__"]
	delete(m, "__meta__")

	if meta == nil {
		return errors.New("no __meta__ table")
	}

	for k, v := range meta.(map[string]interface{}) {
		switch k {
		case "template":
			f.Template = v.(bool)
		case "no-update":
			f.NoUpdate = v.(bool)
		case "generated":
			f.Generated = v.(time.Time)
		case "language":
			f.Language = v.(string)
		case "comments":
			f.Comments = v.(string)
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
		e := Entry{
			ID:      k,
			Loc:     asStringSlice(vv["loc"]),
			Context: asString(vv["context"]),
			Updated: asString(vv["updated"]),
			Default: asString(vv["default"]),
			Zero:    asString(vv["zero"]),
			One:     asString(vv["one"]),
			Two:     asString(vv["two"]),
			Few:     asString(vv["few"]),
			Many:    asString(vv["many"]),
		}
		f.Strings[k] = e
	}

	return nil
}

func asString(i interface{}) string {
	s, ok := i.(string)
	if ok {
		return s
	}
	return ""
}
func asBool(i interface{}) bool {
	b, ok := i.(bool)
	if ok {
		return false
	}
	return b
}
func asStringSlice(i interface{}) []string {
	ii, ok := i.([]interface{})
	if !ok {
		return nil
	}
	s := make([]string, 0, len(ii))
	for _, ss := range ii {
		s = append(s, ss.(string))
	}
	return s
}

// TOML formats all entries as TOML.
//
// TODO: Implement MarshalTOML instead.
func (f File) TOML() (string, error) {
	meta := map[string]interface{}{
		"generated": f.Generated,
		"language":  f.Language,
	}
	if f.Template {
		meta["template"] = f.Template
		meta["options"] = f.Options
	} else {
		meta["maintainers"] = f.Maintainers
	}
	if f.NoUpdate {
		meta["no-update"] = true
	}
	if f.Comments != "" {
		meta["comments"] = f.Comments
	}

	b := new(bytes.Buffer)
	err := toml.NewEncoder(b).Encode(map[string]interface{}{"__meta__": meta})
	if err != nil {
		return "", err
	}

	t, err := f.Strings.toml(f.Template)
	if err != nil {
		return "", err
	}

	return b.String() + "\n" + t, nil
}

// TODO: implement MarshalTOML instead.
func (e Entries) toml(isTpl bool) (string, error) {
	b := new(strings.Builder)
	for _, x := range e.Sorted() {
		fmt.Fprintf(b, "[%s]\n", tomlString(x.ID))
		if x.Updated != "" {
			fmt.Fprintf(b, "  updated = %s\n", tomlString(x.Updated))
		}

		if isTpl {
			switch len(x.Loc) {
			case 0: /// Should never happen, but just in case.
				fmt.Fprintf(b, "  loc     = []\n")
			case 1:
				fmt.Fprintf(b, "  loc     = [%s]\n", tomlString(x.Loc[0]))
			default:
				fmt.Fprintf(b, "  loc = [\n")
				for _, l := range x.Loc {
					fmt.Fprintf(b, "    %s,\n", tomlString(l))
				}
				fmt.Fprintf(b, "  ]\n")
			}
			if x.Context != "" {
				fmt.Fprintf(b, "  context = %s\n", tomlString(x.Context))
			}
		}
		if x.Default != "" {
			fmt.Fprintf(b, "  default = %s\n", tomlString(x.Default))
		}
		if x.Zero != "" {
			fmt.Fprintf(b, "  zero    = %s\n", tomlString(x.Zero))
		}
		if x.One != "" {
			fmt.Fprintf(b, "  one     = %s\n", tomlString(x.One))
		}
		if x.Two != "" {
			fmt.Fprintf(b, "  two     = %s\n", tomlString(x.Two))
		}
		if x.Few != "" {
			fmt.Fprintf(b, "  few     = %s\n", tomlString(x.Few))
		}
		if x.Many != "" {
			fmt.Fprintf(b, "  many    = %s\n", tomlString(x.Many))
		}
		b.WriteByte('\n')
	}
	return b.String(), nil
}

var tomlDblQuote = strings.NewReplacer(
	"\t", "\\t",
	//"\n", "\\n", // We print this as a multi-line message, so don't escape
	"\r", "\\r",
	"\"", "\\\"",
	"\\", "\\\\")

func tomlString(s string) string {
	s = strings.TrimSpace(s)
	var (
		hasNL   = strings.IndexByte(s, '\n') > -1
		hasDblQ = strings.IndexByte(s, '"') > -1
		hasLitQ = strings.IndexByte(s, '\'') > -1
	)
	switch {
	case hasNL || (hasDblQ && hasLitQ):
		return `"""` + "\n" + tomlDblQuote.Replace(s) + "\n" + `"""`
	case hasDblQ:
		return `'` + s + `'`
	default:
		return `"` + tomlDblQuote.Replace(s) + `"`
	}
}
