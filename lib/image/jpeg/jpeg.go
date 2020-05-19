package jpeg

import (
	"image/jpeg"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "image/jpeg",

	"DefaultQuality": jpeg.DefaultQuality,

	"Decode":       jpeg.Decode,
	"DecodeConfig": jpeg.DecodeConfig,
	"Encode":       jpeg.Encode,

	"Options": qlang.StructOf((*jpeg.Options)(nil)),
}
