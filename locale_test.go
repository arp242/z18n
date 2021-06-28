package z18n

import (
	"context"
	"html/template"
	"testing"
)

func TestContext(t *testing.T) {
	b := mkbundle()
	ctx := With(context.Background(), b.Locale("nl-NL"))

	want := "Hallo, &lt;&gt;!"
	t.Run("T", func(t *testing.T) {
		have := T(ctx, "hello-loc", "<>")
		if have != want {
			t.Errorf("\nhave: %q\nwant: %q", have, want)
		}
	})
	t.Run("T", func(t *testing.T) {
		have := Thtml(ctx, "hello-loc", "<>")
		if have != template.HTML(want) {
			t.Errorf("\nhave: %q\nwant: %q", have, want)
		}
	})
	t.Run("nil", func(t *testing.T) {
		if g := Get(context.Background()); g != nil {
			t.Errorf("not nil: %#v", g)
		}

		want := "%(z18n ERROR: T: no context)"
		have := T(context.Background(), "hello-loc", "<>")
		if have != want {
			t.Errorf("\nhave: %q\nwant: %q", have, want)
		}
	})
}
