package msgfile

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template/parse"

	"zgo.at/zstd/zfilepath"
	"zgo.at/zstd/zstring"
	"zgo.at/ztpl"
)

// FindTemplate finds all strings in Go templates.
func FindTemplate(dir string, ext []string, funs ...string) (Entries, error) {
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

		tree, err := ztpl.Parse("", string(data), parse.SkipFuncCheck|parse.ParseComments, "{{", "}}")
		if err != nil {
			return fmt.Errorf("parsing %q: %w", path, err)
		}

		comment := ""
		ztpl.Visit(tree[""].Root, func(n parse.Node, _ int) bool {
			if _, ok := n.(*parse.TextNode); ok {
				return true
			}

			//fmt.Printf("    %-20T → %q\n", n, n.String())

			cmt, ok := n.(*parse.CommentNode)
			if ok {
				if !strings.HasPrefix(cmt.Text, "/* z18n: ") {
					comment = ""
					return true
				}

				comment = strings.TrimSuffix(strings.TrimPrefix(cmt.Text, "/* z18n: "), " */")
				return true
			}

			nn, ok := n.(*parse.ActionNode)
			if !ok {
				comment = ""
				return true
			}

			if nn.Pipe == nil || len(nn.Pipe.Cmds) == 0 || len(nn.Pipe.Cmds[0].Args) < 2 {
				comment = ""
				return false
			}

			name := ""
			switch f := nn.Pipe.Cmds[0].Args[0].(type) {
			case *parse.IdentifierNode:
				name = f.Ident
			case *parse.FieldNode:
				name = "." + strings.Join(f.Ident, ".")
			case *parse.VariableNode:
				name = strings.Join(f.Ident, ".")
				if strings.HasPrefix(name, "$.") {
					name = name[1:]
				}
			}
			if !zstring.Contains(funs, name) {
				comment = ""
				return false
			}

			idlit, ok := nn.Pipe.Cmds[0].Args[1].(*parse.StringNode)
			if !ok || idlit.Text == "" {
				idlit, ok = nn.Pipe.Cmds[0].Args[2].(*parse.StringNode)
				if !ok || idlit.Text == "" {
					comment = ""
					return false
				}
			}
			id, def := zstring.Split2(idlit.Text, "|")

			// if def != "" && errors.Append(typeCheck(id, def)) {
			//    return false
			// }

			e := Entry{ID: id, Default: normalizeMessage(def), Context: comment}
			f, ok := found[id]
			if ok {
				e.Loc = f.Loc
			}

			e.Loc = append(e.Loc, fmt.Sprintf("%s:%d", zfilepath.TrimPrefix(path, dir), nn.Line))
			found[id] = e
			comment = ""
			return false
		})

		return nil
	})

	return found, err
}
