package png

import (
	"image/png"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "image/png",

	"BestCompression":    png.BestCompression,
	"BestSpeed":          png.BestSpeed,
	"DefaultCompression": png.DefaultCompression,
	"NoCompression":      png.NoCompression,

	"Decode":       png.Decode,
	"DecodeConfig": png.DecodeConfig,
	"Encode":       png.Encode,

	"Encoder": qlang.StructOf((*png.Encoder)(nil)),
}
