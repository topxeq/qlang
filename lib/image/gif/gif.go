package gif

import (
	"image/gif"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "image/gif",

	"DisposalBackground": gif.DisposalBackground,
	"DisposalNone":       gif.DisposalNone,
	"DisposalPrevious":   gif.DisposalPrevious,

	"Decode":       gif.Decode,
	"DecodeConfig": gif.DecodeConfig,
	"Encode":       gif.Encode,
	"EncodeAll":    gif.EncodeAll,

	"GIF":       qlang.StructOf((*gif.GIF)(nil)),
	"DecodeAll": gif.DecodeAll,
	"Options":   qlang.StructOf((*gif.Options)(nil)),
}
