package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"zgo.at/z18n/msgfile"
)

func New(language, dir string) error {
	dst := filepath.Join(dir, language+".toml")
	if _, err := os.Stat(dst); err == nil {
		return fmt.Errorf("%q already exists", dst)
	}

	ls, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, l := range ls {
		if !strings.HasSuffix(l.Name(), ".toml") {
			continue
		}

		var f msgfile.File
		_, err := toml.DecodeFile(filepath.Join(dir, l.Name()), &f)
		if err != nil {
			return fmt.Errorf("reading %q: %w", l.Name(), err)
		}

		if f.Template {
			f.Template = false
			f.Language = language

			for k, v := range f.Strings {
				f.Strings[k] = msgfile.Entry{
					ID:  v.ID,
					Loc: v.Loc,
				}
			}

			t, err := f.TOML()
			if err != nil {
				return err
			}

			err = os.WriteFile(dst, []byte(t), 0644)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return fmt.Errorf("no template in %q", dir)
}
