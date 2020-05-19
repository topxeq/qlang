package palette

import (
	"image/color/palette"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "image/color/palette",

	"Plan9":   palette.Plan9,
	"WebSafe": palette.WebSafe,
}
