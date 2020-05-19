package utf8

import (
	"unicode/utf8"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "unicode/utf8",

	"MaxRune":   utf8.MaxRune,
	"RuneError": utf8.RuneError,
	"RuneSelf":  utf8.RuneSelf,
	"UTFMax":    utf8.UTFMax,

	"DecodeLastRune":         utf8.DecodeLastRune,
	"DecodeLastRuneInString": utf8.DecodeLastRuneInString,
	"DecodeRune":             utf8.DecodeRune,
	"DecodeRuneInString":     utf8.DecodeRuneInString,
	"EncodeRune":             utf8.EncodeRune,
	"FullRune":               utf8.FullRune,
	"FullRuneInString":       utf8.FullRuneInString,
	"RuneCount":              utf8.RuneCount,
	"RuneCountInString":      utf8.RuneCountInString,
	"RuneLen":                utf8.RuneLen,
	"RuneStart":              utf8.RuneStart,
	"Valid":                  utf8.Valid,
	"ValidRune":              utf8.ValidRune,
	"ValidString":            utf8.ValidString,
}
