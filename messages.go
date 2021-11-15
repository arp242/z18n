package z18n

import (
	"io/fs"
	"os"

	"github.com/BurntSushi/toml"
	"golang.org/x/text/language"
	"zgo.at/errors"
	"zgo.at/z18n/finder"
	"zgo.at/zstd/zfilepath"
)

func (b *Bundle) ReadMessagesFromFile(lang language.Tag, file string) error {
	return b.ReadMessages(os.DirFS("."), lang, file)
}

func (b *Bundle) ReadMessages(fsys fs.FS, lang language.Tag, file string) error {
	_, ext := zfilepath.SplitExt(file)
	var (
		m   map[string]finder.Entry
		err error
	)
	switch ext {
	default:
		return errors.Errorf("unknown file type: %q", ext)
	case "toml":
		_, err = toml.DecodeFS(fsys, file, &m)
	case "json":
		// TODO
	case "yaml":
		// TODO
	case "go":
		// TODO
	case "po":
		// TODO
	}
	if err != nil {
		return err
	}

	// TODO: is this really needed?
	if m != nil {
		for k, v := range m {
			v.ID = k
			m[k] = v
		}
	}

	// TODO
	for k, v := range m {
		_, _ = k, v
		b.AddMessages(lang, nil)
	}

	return nil
}
