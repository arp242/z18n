package testapp

func T(string) string { return "" }

var _ = T("package-var")

var (
	_ = T("package-var-in-block-1")
	_ = T("package-var-in-block-1")
	_ = T("package-var-in-block-2")
)

var m = map[string]string{
	"a": T("package-var-in-map-1"),
	"b": T("package-var-in-map-2"),
}

func test() {
	T("message-id")
}

// z18n: context for 1.
var _ = T("with-context-1")

func x() {
	// z18n: context for 2.
	var _ = T("with-context-2")
}
