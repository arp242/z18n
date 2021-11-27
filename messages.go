package z18n

import (
	"io/fs"

	"github.com/BurntSushi/toml"
	"golang.org/x/text/language"
	"zgo.at/errors"
	"zgo.at/z18n/finder"
	"zgo.at/zstd/zfilepath"
)

func (b *Bundle) ReadMessagesFromDir(fsys fs.FS, glob string) error {
	ls, err := fs.Glob(fsys, glob)
	if err != nil {
		return errors.Wrap(err, "Bundle.ReadMessagesFromDir")
	}

	for _, f := range ls {
		err := b.ReadMessages(fsys, f)
		if err != nil {
			return errors.Wrapf(err, "Bundle.ReadMessagesFromDir: %q", f)
		}
	}
	return nil
}

func (b *Bundle) ReadMessages(fsys fs.FS, path string) error {
	_, ext := zfilepath.SplitExt(path)
	var (
		file finder.File
		err  error
	)
	switch ext {
	default:
		return errors.Errorf("Bundle.ReadMessages: unknown file type: %q", ext)
	case "toml":
		_, err = toml.DecodeFS(fsys, path, &file)
	}
	if err != nil {
		return errors.Wrap(err, "Bundle.ReadMessages: decode")
	}

	// Basefile is not intended for translations.
	if file.BaseFile {
		return nil
	}

	lang, err := language.Parse(file.Language)
	if err != nil {
		return errors.Wrapf(err, "Bundle.ReadMessages: parse language %q", file.Language)
	}

	for k, v := range file.Strings {
		b.AddMessages(lang, map[string]Msg{
			k: Msg{
				ID:      k,
				Default: v.Default,
			},
		})
	}

	return nil
}
