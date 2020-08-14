package regexpx

import (
	"github.com/topxeq/regexpx"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/regexpx",

	"Match":       regexpx.Match,
	"MatchString": regexpx.MatchString,
	"QuoteMeta":   regexpx.QuoteMeta,

	"Compile":                regexpx.Compile,
	"CompileFreeSpacing":     regexpx.CompileFreeSpacing,
	"MustCompile":            regexpx.MustCompile,
	"MustCompileFreeSpacing": regexpx.MustCompileFreeSpacing,
}
