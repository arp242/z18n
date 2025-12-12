package internal

import (
	"fmt"
	"path/filepath"
	"time"

	"zgo.at/z18n/msgfile"
	"zgo.at/zli"
)

func Find(paths []string, lang, format string, fun, tplFun, tplExt []string) (string, error) {
	if len(paths) == 0 {
		paths = []string{"."}
	}
	found, err := find(paths, fun, tplFun, tplExt)
	if err != nil {
		return "", err
	}

	file := msgfile.File{
		Template: true,
		Language: lang,
		Modified: time.Now().UTC().Round(time.Second),
		Options: map[string]any{
			"paths":    paths,
			"language": lang,
			"fun":      fun,
			"tpl-ext":  tplFun,
			"tpl-fun":  tplExt,
		},
		Strings: found,
	}

	f, ok := map[string]func() (string, error){
		"ls":   file.List,
		"toml": file.TOML,
	}[format]
	if !ok {
		return "", fmt.Errorf("unknown format: %q", format)
	}

	return f()
}

func find(dirs, fun, tplFun, tplExt []string) (msgfile.Entries, error) {
	found := make(msgfile.Entries)
	for _, d := range dirs {
		d, err := filepath.Abs(d)
		zli.F(err)

		f, err := msgfile.FindGo(d, fun...)
		if err != nil {
			return nil, err
		}
		found.Merge(f)

		f, err = msgfile.FindTemplate(d, tplExt, tplFun...)
		if err != nil {
			return nil, err
		}
		found.Merge(f)
	}

	return found, nil
}
