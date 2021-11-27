package main

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"zgo.at/z18n/finder"
	"zgo.at/zli"
)

const usage = `z18n scans files for translatable strings.

https://github.com/zgoat/z18n

Commands:

    help    Show this help.

    find    Find message strings

                Find message strings in a source directory; defauts to the
                current directory if omitted.

                Flags:

                    -format     Format to write; possible values:
                                    toml        TOML file; this is the default.
                                    ls          List in concise syntax, useful
                                                for verifying/auditing strings.

                    -fun        Function names, separated by a comma.
                                Default is "T".

                    -tpl-ext    Extensions for templates; default is "gohtml"
                                and "gotxt". Can be given more than once.

                    -tpl-fun    Template functions; default is "t" and ".T".

    update  Update message files

                Update existing translations; this requires a [base-file] as
                the first positional argument. This is the result of a previous
                "z18n find" command.

                Any further positional arguments are translations; for every
                translation it will add new entries, comment out entries that
                no longer exist, and mark entries where the default has changed.

Typical usage is:

    # Create a base file.
    % z18n find ./.. >i18n/base.toml

    # Translate!
    % cp i18n/base.toml i18n/nl_NL.toml
    % cp i18n/base.toml i18n/id_ID.toml

Now you commit things and everything is translated and everything is good. At
some point you'll probably add, remove, or change translation strings.

    % z18n update i18n/base.toml i18n/*.toml

The options used in "z18n find" are stored in the base file, so you don't need
to add them again.
`

func main() {
	f := zli.NewFlags(os.Args)
	var (
		help   = f.Bool(false, "help", "h")
		format = f.String("toml", "format")
		fun    = f.StringList([]string{"z18n.T", "z18n.Locale.T", "T"}, "fun")
		tplExt = f.StringList([]string{"gohtml", "gotxt"}, "tpl-ext")
		tplFun = f.StringList([]string{".T", "t"}, "tpl-fun")
	)
	zli.F(f.Parse())

	if help.Bool() {
		fmt.Print(usage)
		return
	}

	switch f.ShiftCommand("help", "find", "update") {
	case zli.CommandAmbiguous, zli.CommandUnknown:
		zli.Fatalf("unknown command")
	case "help", zli.CommandNoneGiven:
		fmt.Print(usage)
	case "find":
		paths := []string{"."}
		if len(f.Args) > 1 {
			paths = f.Args[1:]
		}
		found, err := find(paths, fun.Strings(), tplFun.Strings(), tplExt.Strings())
		zli.F(err)

		file := finder.File{
			BaseFile:  true,
			Generated: time.Now().UTC().Round(time.Second),
			Options: map[string]interface{}{
				"paths":   paths,
				"fun":     fun.Strings(),
				"tpl-ext": tplFun.Strings(),
				"tpl-fun": tplExt.Strings(),
			},
			Strings: found,
		}

		f, ok := map[string]func() (string, error){
			"ls":   file.List,
			"toml": file.TOML,
		}[format.String()]
		if !ok {
			zli.Fatalf("unknown format: %q", format)
		}

		out, err := f()
		zli.F(err)
		fmt.Print(out)

	case "update":
		if len(f.Args) < 1 {
			zli.Fatalf("need base file as first argument")
		}
		update(f.Args[0], f.Args[1:])
	}
}

func find(dirs, fun, tplFun, tplExt []string) (finder.Entries, error) {
	found := make(finder.Entries)
	for _, d := range dirs {
		f, err := finder.Go(d, fun...)
		if err != nil {
			return nil, err
		}
		found.Merge(f)

		f, err = finder.Template(d, tplExt, tplFun...)
		if err != nil {
			return nil, err
		}
		found.Merge(f)
	}
	return found, nil
}

func strSlice(i interface{}) []string {
	ii := i.([]interface{})
	s := make([]string, 0, len(ii))
	for _, ss := range ii {
		s = append(s, ss.(string))
	}
	return s
}

func update(baseFile string, mergeFiles []string) {
	var baseMap map[string]interface{}
	_, err := toml.DecodeFile(baseFile, &baseMap)
	zli.F(err)

	meta, ok := baseMap["__meta__"].(map[string]interface{})
	if !ok {
		zli.Fatalf("no '__meta__' table in %q", baseFile)
	}
	if !meta["base-file"].(bool) {
		zli.Fatalf("basefile does not have __meta__.base-file?")
	}
	opts := meta["options"].(map[string]interface{})
	baseEntries, err := find(
		strSlice(opts["paths"]),
		strSlice(opts["fun"]),
		strSlice(opts["tpl-ext"]),
		strSlice(opts["tpl-fun"]))
	zli.F(err)

	base := finder.File{
		Path:      baseFile,
		BaseFile:  meta["base-file"].(bool),
		Generated: meta["generated"].(time.Time),
		Options:   opts,
		Strings:   baseEntries,
	}

	// for k, v := range baseMap {
	// 	vv := v.(map[string]interface{})
	// 	ll := vv["loc"].([]interface{})
	// 	loc := make([]string, 0, len(ll))
	// 	for _, s := range ll {
	// 		loc = append(loc, s.(string))
	// 	}

	// 	base.Strings[k] = finder.Entry{
	// 		ID:      k,
	// 		Default: vv["default"].(string),
	// 		Loc:     loc,
	// 	}
	// }

	merges := make([]finder.File, 0, len(mergeFiles))
	for _, f := range mergeFiles {
		if f == baseFile {
			continue
		}

		m, err := os.ReadFile(f)
		zli.F(err)
		var file finder.File
		_, err = toml.Decode(string(m), &file)
		zli.F(err)

		if file.BaseFile {
			zli.Fatalf("%q is a base-file", f)
		}

		file.Path = f
		mergeFile(&base, &file)
		merges = append(merges, file)
	}

	tt, err := base.TOML()
	zli.F(err)
	zli.F(os.WriteFile(base.Path, []byte(tt), 0644))

	for _, f := range merges {
		tt, err := f.TOML()
		zli.F(err)
		zli.F(os.WriteFile(f.Path, []byte(tt), 0644))
	}
}

// We have two sets of translation strings:
//
//   1. The "base file"; that is, the output of a current "z18n find".
//   2. A translation file.
//
// What we want to do is update the translation file so that:
//
//   1. Any changes from the new and old "z18n find" output should be marked.
//   2. Any translation strings that are no longer used should be marked as such.
//   3. Any new translation strings should be added.
//   4. Make sure the loc array is up-to-date.
func mergeFile(base, merge *finder.File) {
	// Add entries in "base" that are not in "merge".

	// Add new entries.
	for k, foundV := range base.Strings {
		_, ok := merge.Strings[k]
		if !ok {
			merge.Strings[k] = foundV
		}
	}

	// Remove entries not in found.
	for k, mergeV := range merge.Strings {
		baseV, ok := base.Strings[k]
		if ok {
			mergeV.Loc = baseV.Loc
		} else {
			mergeV.Removed = true
		}
		merge.Strings[k] = mergeV
	}
}
