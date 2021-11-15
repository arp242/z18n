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

	for k, v := range m {
		b.AddMessages(lang, map[string]Msg{
			k: Msg{
				ID:      k,
				Default: v.Default,
			},
		})
	}

	return nil
}
