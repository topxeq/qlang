package draw

import (
	"image/draw"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "image/draw",

	"Over": draw.Over,
	"Src":  draw.Src,

	"FloydSteinberg": draw.FloydSteinberg,

	"Draw":     draw.Draw,
	"DrawMask": draw.DrawMask,
}
