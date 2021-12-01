package msgfile

import (
	"errors"
	"io/fs"
)

// TODO: what we really want is a new "Dir" type; would come in handy in a few
// other places too.
func New(fsys fs.FS, language string) (File, error) {
	ls, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return File{}, err
	}

	for _, l := range ls {
		f, err := ReadFile(fsys, l.Name())
		if err != nil {
			return File{}, err
		}

		if f.Template {
			f.Template = false
			f.Language = language
			for k, v := range f.Strings {
				f.Strings[k] = Entry{
					ID:  v.ID,
					Loc: v.Loc,
				}
			}

			return f, nil
		}
	}

	return File{}, errors.New("no template file")
}
