package template

import (
	"html/template"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "html/template",

	"ErrAmbigContext":      template.ErrAmbigContext,
	"ErrBadHTML":           template.ErrBadHTML,
	"ErrBranchEnd":         template.ErrBranchEnd,
	"ErrEndContext":        template.ErrEndContext,
	"ErrNoSuchTemplate":    template.ErrNoSuchTemplate,
	"ErrOutputContext":     template.ErrOutputContext,
	"ErrPartialCharset":    template.ErrPartialCharset,
	"ErrPartialEscape":     template.ErrPartialEscape,
	"ErrPredefinedEscaper": template.ErrPredefinedEscaper,
	"ErrRangeLoopReentry":  template.ErrRangeLoopReentry,
	"ErrSlashAmbig":        template.ErrSlashAmbig,
	"OK":                   template.OK,

	"HTMLEscape":       template.HTMLEscape,
	"HTMLEscapeString": template.HTMLEscapeString,
	"HTMLEscaper":      template.HTMLEscaper,
	"IsTrue":           template.IsTrue,
	"JSEscape":         template.JSEscape,
	"JSEscapeString":   template.JSEscapeString,
	"JSEscaper":        template.JSEscaper,
	"URLQueryEscaper":  template.URLQueryEscaper,

	"Must":       template.Must,
	"New":        template.New,
	"ParseFiles": template.ParseFiles,
	"ParseGlob":  template.ParseGlob,
}
