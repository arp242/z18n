Documentation for translators
-----------------------------
A translation string looks like:

    Hello, %(name)

The `%(name)` is a variable that will be injected in the string. It can be
anywhere in the string but you *must* keep the same variable name.

You can use functions inside `%(name)` as well:

    Hello, %(name upper_first)

The functions may differ per translation; this can be useful to adhere to things
like capitalisation rules and such.

Functions:

    lower         lower-case the entire string.
    upper         upper-case the entire string.
    upper_first   upper-case the first letter of the string.

You don't need to worry about correct localised formatting of dates, numbers,
currency, etc. This is handled automatically.

TODO: mention date formatting etc.

---

A second type of variable looks like:

    Click %[name here]

The `%[..]` signifies that will be wrapped in some HTML, such as a link or
button, but it may also be just some bold text.

The `name` is the variable name, and must be kept intact. Everything after this
will be translated.

You can put variables inside here too:

    Feel free to %[email email me at %(email)] if you have any questions.
