package msgfile

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// List all entries in this file.
func (f File) List() (string, error) {
	return f.Strings.list()
}

// List all messages, one per line, without context.
func (e Entries) list() (string, error) {
	max := 0
	for _, x := range e {
		if l := utf8.RuneCountInString(x.ID); l > max {
			max = l
		}
	}

	b := new(strings.Builder)
	for _, x := range e.Sorted() {
		// TODO: add plurals (or at least mark they exist).
		fmt.Fprintf(b, "%s %s %q\n", x.ID, strings.Repeat(" ", max-utf8.RuneCountInString(x.ID)), x.Default)
	}
	return b.String(), nil
}
