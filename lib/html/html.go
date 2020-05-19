package html

import (
	"html"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "html",

	"EscapeString":   html.EscapeString,
	"UnescapeString": html.UnescapeString,
}
