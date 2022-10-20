package main

import (
	"errors"
	"fmt"
	"os"

	"zgo.at/z18n/internal"
	"zgo.at/zli"
)

const usage = `z18n is an i18n tool: https://github.com/zgoat/z18n

See the README for a more complete introduction.

Quick start:

    The "find" command finds translatable strings:

        % z18n find -language en-GB >i18n/template.toml

    The template.toml file contains all strings found in the source code; the
    -language tags records which language is this application's "native"
    language.

    After this, you can translate it with:

        % z18n new nl-NL i18n

    This will create i18n/nl-NL.toml, which is a copy of template.toml with some
    things modified. There is the assumption that all files will be in a single
    directory. You can also do this manually if you want and know what to
    modify.

    If you use only message ids without a default string or want to use plurals
    then you need to create a new "translation" in the application's native
    language; the template file is never read for translations.

    Now you translate things and everything is good. At some point you'll
    probably add, remove, or change translation strings; you can update the
    translations with "z18n update":

        % z18n update i18n

    Since the flags are recorded in the template file this should be enough. You
    usually only need to run "find" once to create the template file, and after
    that "update" will update it and all translation files.

Commands:

    help    Show this help.

    find    Find message strings from the source files; accepts one or more
            source directories. Defaults to the current directory if omitted.

            Flags:

                -format     Format to write; possible values:
                                toml        TOML file; this is the default.
                                ls          List in concise syntax, useful
                                            for verifying/auditing strings.

                -language   The application's "native" language.

                -fun        Functions considered "translation functions" where
                            to find strings in. The signature needs to be one
                            of:
                                func(string, [..])
                                func(context.Context, string, [..])

                            Can be given multiple times. Defaults to "T" if not
                            given.

                -tpl-fun    Template functions, default if not given is "t"
                            and ".T".

                -tpl-ext    Extensions for templates; can be given more than
                            one. Default is "gohtml" and "gotxt".

    new     Create a new translation; this accepts two positional arguments:

                language    Language to use, as just the language ("nl") or
                            language with region tag ("nl-NL").

                dir         Directory to use; requires a template.toml file
                            (i.e. a file with "template = true").

    update  Update message files; this requires a directory with translation
            files as the first positional argument.

            This will update the template.toml file with all current message
            strings, and for every translation it will add new entries, comment
            out entries that no longer exist, and mark entries where the default
            has changed.
`

func main() {
	f := zli.NewFlags(os.Args)
	var (
		help   = f.Bool(false, "help", "h")
		lang   = f.String("", "language")
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

	switch cmd, err := f.ShiftCommand("help", "find", "new", "update", "check"); cmd {
	case "":
		if errors.Is(err, zli.ErrCommandNoneGiven{}) {
			fmt.Print(usage)
			return
		}

		zli.Fatalf("unknown command: %q", cmd)
	case "help":
		fmt.Print(usage)

	case "find":
		if lang.String() == "" {
			zli.Fatalf("-language must be set")
		}
		out, err := internal.Find(f.Args, lang.String(), format.String(),
			fun.Strings(), tplFun.Strings(), tplExt.Strings())
		zli.F(err)
		fmt.Print(out)

	case "new":
		if len(f.Args) != 2 {
			zli.Fatalf("requires exactly two positional arguments: [language] [dir]")
		}
		err := internal.New(f.Args[0], f.Args[1])
		zli.F(err)

	case "update":
		if len(f.Args) != 1 {
			zli.Fatalf("requires exactly one positional argument: [dir]")
		}
		err := internal.Update(f.Args[0])
		zli.F(err)

	case "check":
		if len(f.Args) != 1 {
			zli.Fatalf("requires exactly one positional argument: [dir]")
		}
		notes, err := internal.Check(f.Args[0])
		zli.F(err)

		for _, n := range notes {
			fmt.Println(n)
		}
		if len(notes) > 0 {
			os.Exit(2)
		}
	}
}
