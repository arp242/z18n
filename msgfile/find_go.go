package msgfile

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
	"zgo.at/errors"
	"zgo.at/zstd/zfilepath"
	"zgo.at/zstd/zstring"
)

// Go finds all translatable strings in dir.
func FindGo(dir string, funs ...string) (Entries, error) {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedTypes | packages.NeedSyntax,
		Dir:  dir,
	}, filepath.Join(dir, "/..."))
	if err != nil {
		return nil, errors.Wrap(err, "msgfile.FindGo")
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
			return nil, errors.Wrap(errs.ErrorOrNil(), "msgfile.FindGo")
		}

		var comment string
		for _, f := range p.Syntax {
			ast.Inspect(f, func(n ast.Node) bool {
				if n == nil {
					return true
				}
				{
					// text := new(bytes.Buffer)
					// printer.Fprint(text, p.Fset, n)
					// fmt.Printf("    %-20T → %q\n", n, text.String())
				}

				cmt, ok := n.(*ast.CommentGroup)
				if ok {
					t := cmt.Text()
					if !strings.HasPrefix(t, "z18n: ") {
						comment = ""
						return true
					}

					comment = strings.TrimPrefix(strings.TrimRight(t, "\n"), "z18n: ")
					return true
				}

				c, ok := n.(*ast.CallExpr)
				if !ok {
					switch n.(type) {
					case *ast.ValueSpec, *ast.Ident, *ast.Comment:
					default:
						comment = ""
					}
					return true
				}

				var name string
				switch f := c.Fun.(type) {
				case *ast.FuncLit, *ast.ArrayType, *ast.ParenExpr, *ast.CallExpr:
					comment = ""
					return true // No work.

				case *ast.Ident: // T(..)
					name = f.Name
				case *ast.SelectorExpr: // pkg.T(..)
					name = f.Sel.Name
				case *ast.IndexExpr: // someMap[T(..)]
					// TODO
					return true
				default:
					pos := p.Fset.Position(c.Pos())
					text := new(bytes.Buffer)
					printer.Fprint(text, p.Fset, c.Fun)
					panic(errors.Errorf("msgfile.FindGo: unhandled c.Fun type %[1]T at %[2]s:%[3]d\n       type → %#[1]v\n       text → %[4]s\n\n",
						c.Fun, pos.Filename, pos.Line, text.String()))
				}

				if !zstring.Contains(funs, name) {
					comment = ""
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
						comment = ""
						return true
					}
				}

				id, def := zstring.Split2(strings.Trim(idlit.Value, quotes), "|")

				// if def != "" && errors.Append(typeCheck(id, def)) {
				//    return false
				// }

				e := Entry{ID: id, Default: normalizeMessage(def), Context: comment}
				if f, ok := found[id]; ok {
					e.Loc = f.Loc
				}

				pos := p.Fset.Position(c.Pos())
				e.Loc = append(e.Loc, fmt.Sprintf("%s:%d", zfilepath.TrimPrefix(pos.Filename, dir), pos.Line))
				found[id] = e

				comment = ""
				return true
			})
		}
	}

	return found, errors.Wrap(errs.ErrorOrNil(), "msgfile.FindGo")
}
