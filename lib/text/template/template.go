package template

import (
	"text/template"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "text/template",

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
