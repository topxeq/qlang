package regexp

import (
	"regexp"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "regexp",

	"Match":       regexp.Match,
	"MatchReader": regexp.MatchReader,
	"MatchString": regexp.MatchString,
	"QuoteMeta":   regexp.QuoteMeta,

	"Compile":          regexp.Compile,
	"CompilePOSIX":     regexp.CompilePOSIX,
	"MustCompile":      regexp.MustCompile,
	"MustCompilePOSIX": regexp.MustCompilePOSIX,
}
