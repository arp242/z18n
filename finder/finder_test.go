package finder

import (
	"os"
	"testing"

	"zgo.at/zstd/ztest"
)

func TestGo(t *testing.T) {
	tmp := ztest.TempFiles(t,
		"go.mod", `module test

go 1.17

require (
        golang.org/x/text v0.3.7
        zgo.at/z18n v0.0.0-20211127015941-59d647da7743
)

require (
        github.com/BurntSushi/toml v0.4.1 // indirect
        golang.org/x/mod v0.4.2 // indirect
        golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
        golang.org/x/tools v0.1.7 // indirect
        golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
        zgo.at/errors v1.0.1-0.20210313142254-4e0fb19b1249 // indirect
        zgo.at/zstd v0.0.0-20211017205211-017273d7d29c // indirect
        zgo.at/ztpl v0.0.0-20211017232908-7dce3dc79277 // indirect
)
`,
		"go.sum", `
github.com/BurntSushi/toml v0.4.1 h1:GaI7EiDXDRfa8VshkTj7Fym7ha+y8/XxIgD2okUIjLw=
github.com/BurntSushi/toml v0.4.1/go.mod h1:CxXYINrC8qIiEnFrOxCa7Jy5BFHlXnUU2pbicEuybxQ=
github.com/yuin/goldmark v1.4.0/go.mod h1:mwnBkeHKe2W/ZEtQ+71ViKU8L12m81fl3OWwC1Zlc8k=
golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2/go.mod h1:djNgcEr1/C05ACkg1iLfiJU5Ep61QUkGW8qpdssI0+w=
golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550/go.mod h1:yigFU9vqHzYiE8UmvKecakEJjdnWj3jj499lnFckfCI=
golang.org/x/mod v0.4.2 h1:Gz96sIWK3OalVv/I/qNygP42zyoKp3xptRVCWRFEBvo=
golang.org/x/mod v0.4.2/go.mod h1:s0Qsj1ACt9ePp/hMypM3fl4fZqREWJwdYDEqhRiZZUA=
golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
golang.org/x/net v0.0.0-20190620200207-3b0461eec859/go.mod h1:z5CRVTTTmAJ677TzLLGU+0bjPO0LkuOLi4/5GtJWs/s=
golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d/go.mod h1:9nx3DQGgdP8bBQD5qxJ1jj9UTztislL4KSBs9R2vV5Y=
golang.org/x/sync v0.0.0-20190423024810-112230192c58/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20210220032951-036812b2e83c/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20190412213103-97732733099d/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20201119102817-f84b799fce68/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210423082822-04245dca01da/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e h1:WUoyKPm6nCo1BnNUvPGnFG3T5DUVem42yDJZZ4CNxMA=
golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
golang.org/x/term v0.0.0-20210317153231-de623e64d2a6/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
golang.org/x/text v0.3.0/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
golang.org/x/text v0.3.6/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.7 h1:olpwvP2KacW1ZWvsR7uQhoyTYvKAupfQrRGBFM352Gk=
golang.org/x/text v0.3.7/go.mod h1:u+2+/6zg+i71rQMx5EYifcz6MCKuco9NR6JIITiCfzQ=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/tools v0.0.0-20191119224855-298f0cb1881e/go.mod h1:b+2E5dAYhXwXZwtnZ6UAqBI28+e2cm9otk0dWdXHAEo=
golang.org/x/tools v0.1.7 h1:6j8CgantCy3yc8JGBqkDLMKWqZ0RDU2g1HVgacojGWQ=
golang.org/x/tools v0.1.7/go.mod h1:LGqMHiF4EqQNHR1JncWGqT5BVaXmza+X+BDGol+dOxo=
golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=
golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
zgo.at/errors v1.0.1-0.20210313142254-4e0fb19b1249 h1:2omJRHHcR1x7itExvYb2+qmBjtMtEuzCLac5ISYolYI=
zgo.at/errors v1.0.1-0.20210313142254-4e0fb19b1249/go.mod h1:POfgvh1LafF2NZJk6buGYCIhcHWuR/miB3nndyf3ozs=
zgo.at/json v0.0.0-20200627042140-d5025253667f/go.mod h1:DUuNIRpNC7/cup+Gy0qhwQEjjlLFXXRZ04VnVv9bf3E=
zgo.at/z18n v0.0.0-20211127015941-59d647da7743 h1:FcTCJO6r/OyXWmZO2plOLkktcESbonj1wDH0C1UtWPQ=
zgo.at/z18n v0.0.0-20211127015941-59d647da7743/go.mod h1:nN4hcWbII1wrdUYF2sI2Y1zu/vuoRxx99xlsvOVuVyw=
zgo.at/zli v0.0.0-20211017231103-84f8e371c324/go.mod h1:pWkBHZ5qoEmvSDbFYdQMAKaRMEuAtdUKb21k/OOYOJw=
zgo.at/zstd v0.0.0-20210512041107-8951517febd3/go.mod h1:sQqrTxBwKW0nlwcOg9RxXB8ikY+NBciJnJRPOq/gEuY=
zgo.at/zstd v0.0.0-20211017205211-017273d7d29c h1:Tn/vQ0YyXVbf6RWMuUQdHLNo0WWxbfHjIXlQb7x2xwQ=
zgo.at/zstd v0.0.0-20211017205211-017273d7d29c/go.mod h1:sQqrTxBwKW0nlwcOg9RxXB8ikY+NBciJnJRPOq/gEuY=
zgo.at/ztpl v0.0.0-20211017232908-7dce3dc79277 h1:FicTpGqTvbeeRMA4tp+hmQ3vthoNoTr8KFq04l/uDec=
zgo.at/ztpl v0.0.0-20211017232908-7dce3dc79277/go.mod h1:mhjSF7+pfwMjNhDgmw9Ywr1CixzeVWuvMjF7gLJT8Uw=
`,
		"test.go", `package test

import (
    "context"

    "golang.org/x/text/language"
    "zgo.at/z18n"
)

var (
	loc = z18n.NewBundle(language.BritishEnglish).Locale()
	ctx = z18n.With(context.Background(), loc)
)

var a = loc.T("a|aaa")

var T = z18n.T

func test() {
	ctx = context.Background()

	var b = loc.T("b|bbb")

	var (
		c = loc.T("c|ccc")
		d = z18n.T(ctx, "d|ddd")
	)

	// z18n: context for e
	e := T(ctx, "e|eee")

	_, _, _, _ = b, c, d, e
}
`, "test.gotxt", `
	{{t "tpl|tpl str"}}
`)

	wd, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(wd)

	entries, err := Go(tmp, "T", "z18n.T")
	if err != nil {
		t.Fatal(err)
	}

	tpl, err := Template(tmp, []string{"gotxt"}, "t")
	if err != nil {
		t.Fatal(err)
	}

	entries.Merge(tpl)

	have, err := entries.toml()
	if err != nil {
		t.Fatal(err)
	}

	want := `["a"]
  loc     = ["test.go:15"]
  default = "aaa"

["b"]
  loc     = ["test.go:22"]
  default = "bbb"

["c"]
  loc     = ["test.go:25"]
  default = "ccc"

["d"]
  loc     = ["test.go:26"]
  default = "ddd"

["e"]
  loc     = ["test.go:30"]
  default = "eee"

["tpl"]
  loc     = ["test.gotxt:2"]
  default = "tpl str"

`

	if d := ztest.Diff(have, want); d != "" {
		t.Error(d)
	}
}
