z18n is a i18n library for Go. It can be used to translate Go applications
and/or localize various aspects such as dates and numbers.

Import as `zgo.at/z18n`. API docs: https://godocs.io/zgo.at/z18n

The chief motivation for writing this is that I wanted a nice painless API (as
painless as i18n gets anyway), and none of the existing solutions that I could
find really offered this, not without some extensive wrapping anyway.

It supports pluralisation, named variables with a convenient syntax, placeholder
syntax for HTML tags, and localisation of dates and numbers.

<!-- TODO: not yet..., and includes a web interface to translate messages. -->

[go-i18n]: https://github.com/nicksnyder/go-i18n

README index:

- [Adding it to an application](#adding-it-to-an-application)
- [Translating messages](#translating-messages)
  - [Variables](#variables)
  - [HTML tags](#html-tags)
  - [Localisation of dates, numbers](#localisation-of-dates-numbers)
  - [Plurals](#plurals)
  - [Adding context](#adding-context)
- [Using from templates](#using-from-templats)
- [JavaScript](#javascript)
- [Finding messages and creating translation files](#finding-messages-and-creating-translation-files)
- [Accepting translations and updates from translators](#accepting-translations-and-updates-from-translators)
- [Other tidbits](#other-tidbits)

Adding it to an application
---------------------------
Start by creating a new bundle; a "bundle" is a set of all translations your
application has:

```go
b, err := NewBundle(language.English)
```

This sets the default language set to English. The "default language" is the
"native" language of the application. I will use English here, but there is
nothing stopping you from writing an application in Russian and then translating
that to English or other languages.

You add messages for a language:

```go
b.AddMessages(language.English, map[string]Msg{
    "insult/cow":   Msg{Default: "You fight like a dairy farmer!"},
    "comeback/cow": Msg{Default: "How appropriate. You fight like a cow!"},
})

b.AddMessages(language.Dutch, map[string]Msg{
    "insult/cow":   Msg{Default: "Je vecht als een melkboer!"},
    "comeback/cow": Msg{Default: "Erg toepasselijk. Je vecht als een koe!"},
})
```

You can also load messages from TOML files; see the "Finding messages and
creating translation files" section below.

To get messages you get a locale from the bundle:

```go
l := b.Locale("nl-NL")
```

This accepts multiple languages, in order of preference, and accepts the
contents of the `Accept-Language` HTTP header. In a real-world web app you
usually want to do something like:

```go
l := b.Locale(
    r.Query.Get("lang"),             // Prefer explicit overwrite from query param.
    user.Settings.Language,          // User setting in your application.
    r.Header.Get("Accept-Language")) // Last option is to use the browser setting.
```

*Aside: please do not use the IP address for this. As someone lived abroad for a
few years it's a massive PITA to have things automatically be set to a language
I don't understand. Even in the Netherlands I often just prefer the English
version.*

If you have a CLI or desktop app you can use `b.LocaleFromEnv()` to create a new
locale; this will use the `LANG`, `LANGUAGES`, and `LC_*` environment variables.

You almost certainly want to use codes with a region tag such as `nl-NL` when
creating locales, `nl` being the language code for Dutch, and `NL` being the
region of the Netherlands. It will fall back to the messages for `nl` if there
aren't any for `nl-NL` specifically, and the localisation for dates and such
will be appropriate for this region. American and British people don't write
their dates in the same way for example, even though they both speak English.

---

You can use `Locale.T` to display translated messages:

```go
fmt.Println(l.T("insult/cow"))
// Output: "Erg toepasselijk. Je vecht als een koe!"
```

More details on how to translate messages in the section below.

A bundle only needs to be created once; if you're using this in a long-running
webapp then create a bundle once on startup and a locale for every user/request
based on the user settings, `Accept-Language` header, etc.

There is also a top-level `z18n.T` which takes the Locale object from a context,
which can be created with `z18n.With()`:

```go
ctx := z18n.With(context.Background(), l)
z18n.T(ctx, "insult/cow")
```

Translating messages
--------------------
The `T` function accepts the message ID as the first parameter:

```go
l.T("song/coconuts")
```

This will look up the message with the ID `song/coconuts` in this locale. If no
such message exists in this locale then it will try to get it from the closest
locale such as `nl`.

This means you can translate messages in the `nl` language, which should be
appropriate for most Dutch speakers, but also add a few regional variations for
`nl-BE` (Dutch as spoken in Belgium) if need be.

A message ID can contain any printable character, including whitespace, but
cannot start with a bar (`|`).

You can optionally specify a default message after a `|`:

```go
l.T("song/coconuts|I've got a lovely bunch of coconuts!")
```

By adding both an ID *and* the message in there you get the best of both worlds:
you can still easily find and write code without going back and forth to
translation files, but you can also freely make minor changes to the default
message without invalidating all the translations as the message itself isn't
used as the lookup key.

You can only set the default (unpluralized) message with this; if you want a
pluralized string then you will need to add the other variants as messages
through an `Bundle.AddMessages()` call or message file (more on pluralisation
and messages files later).

The `/` doesn't mean anything special, it's just a convention that might be
useful. You can also use `song-coconuts`, `song#coconuts`, `song coconuts`,
`song/silly/coconuts`, or just not use any prefix at all and use only
`coconuts`. Personally I found using prefixes useful as it adds a bit context
*what* something is (e.g. `btn/accept` clarifies this is a button).

### Variables
Variable interpolation works with `%(varname)`; the `varname` should remain
identical in translated messages as it's used to look up the variable:

```go
l.T("spam|Spam %(email) at your own risk; I will hunt you down!", email)
```

If you have only one variable then you can pass it as just an argument, but if
if you have more than one you will need to use a map:

```go
l.T("spam|Spam %(email) or %(email2) at your own risk; I will hunt you down!", z18n.P{
    "email":  email,
    "email2": email2,
})
```

The reason for this is that the position of the variables might be different in
translated messages. With one variable this isn't an issue and since many
messages will likely have just one variable it's a useful "shortcut" to have.

Variable names can contain any character except whitespace and any of `%()[]|`.
After a space you can add one or more format specifiers:

    lower, upper    Lower or uppercase everything.
    upper_first     Uppercase the first letter, leaving the rest of the case intact.

    html            Escape as HTML string; only has effect if Bundle.NoHTML is set.

    raw             Don't format numbers or dates according to the locale.
                    i.e. will print just "1000" instead of "1,000".

    date, time      Print a time.Time as a date ("18-06-1985") or time ("17:15:30")
                    Default is to print as datetime ("18-06-1985 17:15:30")

    full, long      Print date, time, or datetime in  "full", "long", "medium",
    medium, short   or "short" format. Defaults to "long"

    day, month      Get the day or month name from a time.Time; combine with
                    "short" to get an abbreviated name.

    max             Set the minimum and maximum percision for floats.
    min:max

    [..]            Any other text for a time.Time is taken as a format string
                    for time.Format.

For example:

    %(word upper_first)         Format in upper case.
    %(word lower upper_first)   Format in lower case, and then upper-case the first letter.


See the "Localisation of dates, numbers" section below for more details on
number and date formatting.

Variable values are always HTML-escaped by default unless you set `NoHTML` in
the `Bundle`. You almost certainly want to set this for CLI and desktop apps,
it's enabled by default as forgetting to do so for webapps can be potentially
disastrous, whereas it "only" looks wrong for a CLI or desktop app (but isn't
dangerous).

### HTML tags
Use `%[varname text]` as a placeholder for HTML tags; this is intended to be
used with the `z18n.Tag()` function and removes the need for most – if not all –
HTML inside translation strings, and makes stuff easier to read and HTML easier
to update:

```go
l.T("video/goat|Look at %[cute video] of a farting goat!")
    z18n.Tag("a", `href="/goat.mp4" class="link"`))
```

Everything inside the `%[..]` tag is text that should be translated. The
`z18n.Tag` function controls which HTML gets added; the first parameter (`a`) is
the tag name, and the second whatever you want to put in the opening tag; this
is added without processing and is **not safe** against untrusted input.

As with `%(..)` variables you can pass it as just an argument if you have only
one value, but will need to use a map and add a variable if you use multiple
variables and/or tags:

```go
l.T("video/goat|Look at %[%link cute video] of a farting %[%bold goat]!", z18n.P{
    "link": z18n.Tag("a", `href="/goat.mp4" class="link"`),
    "bold": z18n.Tag("strong", ``),
})
```

If the first word starts with `%` it's taken as the variable name; the `%` is
added to distinguish between a variable name and a regular word.

You can use variables inside `%[..]` tags:

```go
l.T("email/goat|You can %[%link email the goat at %(email)] for requests", z18n.P{
    "link":  z18n.Tag("a", `href="mailto:TheFlatulentGoat@example.com"`),
    "email": email,
})
```

It's not possible to nest these tags: `%[%one %[%two tags]]` won't work. You can
create your own type which implements the `z18n.Tagger` interface if you need
more complex HTML.

In some cases there isn't any text to be translated; in which case you can use
the third parameter to set the innerHTML and use just a variable name:

```go
l.T("email/goat|Or phone the goat at %(phone).",
    z18n.Tag("a", `href="tel:5554242"`, "555-42 42))
```

### Localisation of dates, numbers
Numbers (all `int` and `float` types) and `time.Time` will be formatted
according to the locale, for example:

```go
l.T("test|The genetic test I did on %(t) showed I'm %(f)% platypus; a trait %(n) people share", z18n.P{
    "t": time.Now(),
    "f": 13.42,
    "n": 51341,
})
```

This will print the following for the en-US, en-NZ, and nl-NL locales:

    The genetic test I did on Jun 12, 2021, 7:30:21 AM showed I'm 13.42% platypus; a trait 51,341 people share
    The genetic test I did on 12/06/2021, 7:30:43 AM showed I'm 13.42% platypus; a trait 51,341 people share
    The genetic test I did on 12 jun. 2021 07:32:57 showed I'm 13,42% platypus; a trait 51.341 people share

We can add some format specifiers to the variables to control how it's printed;
for example adding `0` to the float to and `short date` to the time:

```go
l.T("test|The genetic test I did on %(t short date) showed I'm %(f 0)% platypus; a trait %(n) people share", z18n.P{
    "t": time.Now(),
    "f": 13.42,
    "n": 51341,
})
```

Will print:

    The genetic test I did on Saturday, June 12, 2021 showed I'm 13% platypus; a trait 51,341 people share
    The genetic test I did on Saturday, 12 June 2021 showed I'm 13% platypus; a trait 51,341 people share
    The genetic test I did on zaterdag 12 juni 2021 showed I'm 13% platypus; a trait 51.341 people share

Dates can be printed as "datetime" (the default), "date", or "time", all of them
in a "full" "long", "medium", or "short" format. How exactly it's printed
differs per language; a few examples of how it roughly looks:
    
                            en-US                       nl-NL

        %(d)                Jan 02, 2006 2:22 pm        2 Jan 2006 14:22
        %(d short)          02/01/06 2:22 pm            01-02-2006 14:22
        %(d long)           January 2, 2006 2:22 pm     2 January 2006 14:22

    Or print just the date:
        %(d date)           02/01/06                    01-02-2006
        %(d date short)     January 2, 2006             2 January 2006
        %(d date long)      January 2, 2006             2 January 2006

    Just the time
        %(d time)           2:22                        14:22

    Extract specific parts:
        %(d day)            Monday                      Maandag
        %(d month)          March                       Maart

    Or use a custom format:
        %(d 2006-01-02)     2006-01-02                  2006-01-02

Things like ordinals, formatting of bytes, etc. aren't implemented (yet).

Use the `raw` function in variables to prevent formatting and format it as
`fmt.Sprintf("%v")`:

    l.T("id|number: %(n raw); float: %(f raw); time: %(t raw)", z18n.P{
        "n": 1_230_495,
        "f": 6666.42,
        "t": time.Now(),
    })

<!-- TODO: unimplemented.

---

The formats can be overridden for a locale, in case the user set something
different. You usually want to provide some feature for this: there are probably
a number of people who use en-US (more or less the default for many tech peeps)
but aren't used to the "reverse" US date style, and some people prefer may
`2006-01-02` (or maybe some other format) as well. People have all sort of
preferences and there isn't a *true* standard, it's just a reasonable default
that will likely work well for the majority of people.

For CLI and desktop apps this is easy, as you can use the system's locale;
`Bundle.LocaleFromEnv()` does this already and you don't need to do much more.

    TODO: figure this out, maybe it's better to store the LANG and LC_ vars
    instead? Just one tag isn't really enough.

    type Locale struct {
        *localize
        bundle *Bundle
        tags        []language.Tag
        tagNumeric  language.Tag
        tagTime     language.Tag
        tagCollate  language.Tag
    }

    func (b *Bundle) Locale(langs ...string) *Locale {
    }

    l := bundle.Locale(language.MustParse("en-US"))

    // Set dates; datetime format is derives from this.
    l.Date("2006-01-02")
    l.DateMedium("2 Jan 2006")
    l.DateLong("2 January 2006")
    l.Time("15:16")

    // Thousands and fraction separators.
    l.Thousands('.')
    l.Fraction('.')
-->

### Plurals
Thus far we've only set `Msg.Default`; this is the message to use if there are
no pluralisations to apply; there are five other messages:

    One, Zero, Two, Few, Many

z18n will use one of these (or the `Default`) automatically when supplied with a
`Plural` value. Leaving the appropriate value empty in the message file will
result in an error.

The exact meaning of these vary per language, and most languages don't have all
of them. The logic for all of this can actually be quite complex and often
includes exceptions – as languages do. Plurals in English (and most Germanic
languages) are usually fairly easy with just "one" and "more than one", and you
only need to set `One` and `Default`. Many Asian languages like Indonesian have
it even easier by just not having plural forms at all, and Polish people [must
have a Ph.D. in math embedded in their DNA][pl].

To add Plurals to the messages use the appropriate field(s):

```go
b.AddMessages(language.BritishEnglish, map[string]Msg{
    "ants!": Msg{
        One:     "Help, I've got an ant in my trousers!"
        Default: "Help, I've got %(n) ants in my trousers!"
    },
})

b.AddMessages(language.AmericanEnglish, map[string]Msg{
    "ants!": Msg{
        One:     "Help, I've got an ant in my pants!",
        Default: "Help, I've got %(n) ants in my pants!",
    },
})

b.AddMessages(language.Indonesian, map[string]Msg{
    "ants!": Msg{
        // You can probably get away with just "Default" here; there is no
        // grammatical difference other than not putting the %(n) in there.
        One:     "Tolong, saya punya semut di celana saya! ",
        Default: "Tolong, saya punya %(n) semut di celana saya!",
    },
})

b.AddMessages(language.Polish, map[string]Msg{
    "ants!": Msg{
        One:  "Pomocy, mam mrówkę w spodniach!",
        Two:  "Pomocy, mam %(n) mrówki w spodniach!",:
        Few:  "Pomocy, mam %(n) mrówek w spodniach!",:
        Many: "Pomocy, mam w spodniach %(n) mrówek!",:
    },
})
```

Pass a `z18n.Plural` to `z18n.T` to tell z18n which form to use; `z18n.N()`
conveniently creates this without too much noise:

```go
l.T("ants!", z18n.N(5))
```

This can be in any position and will automatically be made available as the
variable `n`, and can of course be combined with other variables:

```go
b.AddMessages(language.BritishEnglish, map[string]Msg{
    "marketers": Msg{
        One:     "I emailed %(email) only once with my stupid marketing offer, so better send 5 more emails"
        Default: "I emailed %(email) %(n) times with my stupid marketing offers",
    },
})

l.T("marketers", z18n.N(5), email)
```

Note that in CLDR "Default" is named "Other". I renamed this as I thought it
made more sense, especially since most messages don't have any
plurals/translations setting "Other" seems kinda weird.

[pl]: plural/rule_gen.go#L344

### Adding context
It's often useful for translators to have some clue what exactly a string refers
to. Generally speaking, the shorter the string, the more useful adding context
is. 

There are two ways to add context; the first is in the message ID; for example:

    l.T("button/get-quote|Get")

This makes it clear that "get" is used as a form button to get a quote. This may
be important, because words like "get" can sometimes be translated in multiple
ways, and not all of them may be appropriate in this context. For example one
product I worked one had "Get [product-name]" translated in such a way where
"get" meant "to contract" (as in, "contract the flu"). Funny? Yes. A good
translation? Not really.

This is one reason I like using prefixes, because now it's pretty clear that
this is a button that *does* something (but you can also use `get-quote-button`
or some other variation, if you prefer).

A second way is to use special comments:

```go
// z18n: Context
// Context continues.
l.T("button/get-quote|Get")
```

Or:

```go
/* z18n: Context
   Context continues. */
l.T("button/get-quote|Get")
```

Or:

```go
l.T("button/get-quote|Get") // z18n: some context
l.T("button/get-quote|Get") /* z18n: some context */
```

I would recommend avoiding this unless necessary; good ids are better. But
sometimes even with descriptive IDs it's useful to add some extra context.

The comments *need* to be prefixed with `z18n:` and immediately precede the `T`
call. This **won't** work:

```go
// z18n: there is a blank like.

l.T("button/get-quote|Get")

// z18n There is no ":" after z18n.
l.T("button/get-quote|Get")

//z18n: there is no space after //
l.T("button/get-quote|Get")

l.T("button/get-quote|Get")
// z18n: This is on the next line.
```

Using from templates
--------------------
Using z18n from {text,html}/template is a first-class use case; you'll need to
add a few functions:

```go
tplfunc := template.FuncMap{
    "t":      z18n.Thtml,
    "tag":    z18n.Tag,
    "plural": z18n.N,
}
```

And then use:

    {{t .Context "message/swedish|The Swedish prime minister has a massage for you."}}

Whitespace after at the start and end will be stripped and all other whitespace
will be collapsed to a single space, so multi-line messages work well:

    {{t .Context `message/swedish|
        The Swedish prime minister has a massage for you.
    `}}

The downside of this is that you need to pass the context every time. You can
create a "scoped" version by assigning a variable:

```go
b := NewBundle(...)
l := b.Locale()

tpl.ExecuteTemplate("foo.gohtml", struct {
    T func(string, ...interface{}) template.HTML
}{l.T})
```

And then use it like:

    {{.T "message/swedish|The Swedish prime minister has a massage for you."}}

Or if there is only one user for the application (i.e. a CLI or desktop app)
then you can use a regular `FuncMap` with a closure to achieve the same.

You can add context with `{{/* z18n: ... */}}` with the same rules as Go files.

JavaScript
----------
z18n is a Go i18n tool, not a JavaScript one; there isn't great support for
JavaScript right now.

That said, it's not uncommon to have an application where almost all of the text
is in the backend with just a few messages in the frontend. The general strategy
would be to render the messages you need server-side and then load them in JS. A
simple example:

```html
<span id="z18n" style="display: none"
    data-msg-one="{{.T "id/msg"}}
    data-msg-two="{{.T "id/msg"}}
></span>
```

And then get it from `#z18n` on init or when needed.

Or you can add a JSON endpoint and load it from there.

Support for some better/more convenient integration is something I plan to add
later. There are actually quite a few i18n JS libraries out there already, and I
need to see if any of them can be integrated.

You can, by the way, also use the [`Intl`][intl] API to get people's locale
preferences, which is more fine-grained than `Accept-Language`.

[intl]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl


Finding messages and creating translation files
-----------------------------------------------
The `./cmd/z18n` tool can find messages in Go and template files.

To find all messages in the current directory and all subdirectories:

    % z18n init i18n

This will scan for Go files and templates. See `z18n help` for various options.

    % z18n find > i18n/base.toml

To translate something, copy `i18n/base.toml` to `i18n/file.toml`, edit the
information in the `__meta__` key, and start translating!

To load messages from `ReadMessages()`:

```go
err := b.ReadMessages(os.DirFS("i18n"), "file.toml")

// Or all files in a directory:
err := b.ReadMessagesFromDir(os.DirFS("i18n"), "*.toml")
```

---

For updating use the `update` command:

    % z18n update i18n/base.toml i18n/*.toml

The first file is the "base" file, and all other files are translation files.
This will re-scan the project, add new translation entries, comment out entries
that no longer exist, and mark entries where the default has changed.

The options used in "z18n find" are stored in the base file, so you don't need
to add them again.

<!--
Other tidbits
-------------

TODO: mention log here.

TODO: mention marking variables.
-->
