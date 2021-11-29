package msgfile

import "strings"

var goRawQuote = strings.NewReplacer("`", "` + \"`\" + `")

// This has a tendency to break Vim syntax highlights so put it here... Yeah, I
// know :-/
const quotes = "\"`"
