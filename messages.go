package z18n

import (
	"fmt"
	"io/fs"

	"golang.org/x/text/language"
	"zgo.at/errors"
	"zgo.at/z18n/msgfile"
)

// ReadMessagesDir reads all message files from fsys matching the glob pattern.
func (b *Bundle) ReadMessagesDir(fsys fs.FS, glob string) error {
	ls, err := fs.Glob(fsys, glob)
	if err != nil {
		return errors.Wrap(err, "Bundle.ReadMessagesDir")
	}

	for _, f := range ls {
		err := b.ReadMessages(fsys, f)
		if err != nil {
			var tErr tplErr
			if !errors.As(err, &tErr) {
				return errors.Wrapf(err, "Bundle.ReadMessagesDir: %q", f)
			}
		}
	}
	return nil
}

type tplErr string

func (t tplErr) Error() string {
	return fmt.Sprintf("z18n.ReadMessages: can't read messages from a template file (%q)", string(t))
}

// ReadMessages reads a single messages files.
func (b *Bundle) ReadMessages(fsys fs.FS, path string) error {
	file, err := msgfile.ReadFile(fsys, path)
	if err != nil {
		return errors.Wrap(err, "Bundle.ReadMessages: decode")
	}

	// Template file is not intended for translations.
	if file.Template {
		return tplErr(path)
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
				Zero:    v.Zero,
				One:     v.One,
				Two:     v.Two,
				Few:     v.Few,
				Many:    v.Many,
			},
		})
	}

	return nil
}
