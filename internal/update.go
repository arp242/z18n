package internal

import (
	"os"
	"path/filepath"

	"zgo.at/z18n/msgfile"
)

func Update(dir string) error {
	ls, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var (
		fsys   = os.DirFS(dir)
		tpl    msgfile.File
		merges = make([]msgfile.File, 0, len(ls)-1)
	)
	for _, l := range ls {
		file, err := msgfile.ReadFile(fsys, l.Name())
		if err != nil {
			return err
		}

		if file.Template {
			tpl = file
		} else if !file.NoUpdate {
			merges = append(merges, file)
		}
	}

	// Find new entries, merge them in.
	newstr, err := find(
		asStringSlice(tpl.Options["paths"]),
		asStringSlice(tpl.Options["fun"]),
		asStringSlice(tpl.Options["tpl-ext"]),
		asStringSlice(tpl.Options["tpl-fun"]))
	if err != nil {
		return err
	}

	// TODO
	// for k, v := range newstr {
	// 	if v.Default != tpl.Strings[k].Default {
	// 		//fmt.Printf("%-20s: %q != %q\n", k, v.Default, tpl.Strings[k].Default)
	// 		v.Updated = msgfile.UpdatedChanged
	// 		newstr[k] = v
	// 	}
	// }
	tpl.Strings = newstr

	for _, m := range merges {
		mergeFile(tpl, &m)
	}

	// Write it out.
	tt, err := tpl.TOML()
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(dir, tpl.Path), []byte(tt), 0644)
	if err != nil {
		return err
	}

	for _, f := range merges {
		tt, err := f.TOML()
		if err != nil {
			return err
		}
		err = os.WriteFile(filepath.Join(dir, f.Path), []byte(tt), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

// We have two sets of translation strings:
//
//  1. The "template file"; that is, the output of a current "z18n find".
//  2. A translation file.
//
// What we want to do is update the translation file so that:
//
//  1. Any translation strings that are no longer used should be marked as such.
//
//  2. Any new translation strings should be added.
//
//  3. Make sure the loc array is up-to-date.
//
//  4. Any changes from the new and old "z18n find" output should be marked.
//     TODO: not implemmented; actually not sure we want/need this, as minor
//     typo corrections don't need a new translation.
//
//     Can just mark them manually if the changes are large enough to warrant a
//     check by translators.
func mergeFile(tpl msgfile.File, merge *msgfile.File) {
	// Add new entries.
	for k, tplV := range tpl.Strings {
		_, ok := merge.Strings[k]
		if !ok {
			def := ""
			if tpl.Language == merge.Language {
				def = tplV.Default
			}

			merge.Strings[k] = msgfile.Entry{
				ID:      tplV.ID,
				Loc:     tplV.Loc,
				Context: tplV.Context,
				Default: def,
			}
		}
	}

	// Remove entries not in tpl, mark changed entries, and update loc.
	for k, mergeV := range merge.Strings {
		tplV, ok := tpl.Strings[k]
		if !ok {
			delete(merge.Strings, k)
			continue
		}

		mergeV.Loc = tplV.Loc
		mergeV.Context = tplV.Context
		mergeV.Updated = ""
		if tplV.Updated == msgfile.UpdatedChanged {
			mergeV.Updated = msgfile.UpdatedChanged
		}
		merge.Strings[k] = mergeV
	}
}

func asStringSlice(i any) []string {
	ii := i.([]any)
	s := make([]string, 0, len(ii))
	for _, ss := range ii {
		s = append(s, ss.(string))
	}
	return s
}
