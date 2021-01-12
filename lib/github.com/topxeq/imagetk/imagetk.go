package imagetk

import (
	"github.com/topxeq/imagetk"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/imagetk",

	"BOTTOM":            imagetk.BOTTOM,
	"BOTTOM_LEFT":       imagetk.BOTTOM_LEFT,
	"BOTTOM_RIGHT":      imagetk.BOTTOM_RIGHT,
	"Bicubic":           imagetk.Bicubic,
	"Bilinear":          imagetk.Bilinear,
	"CENTER":            imagetk.CENTER,
	"LEFT":              imagetk.LEFT,
	"Lanczos2":          imagetk.Lanczos2,
	"Lanczos3":          imagetk.Lanczos3,
	"MitchellNetravali": imagetk.MitchellNetravali,
	"NearestNeighbor":   imagetk.NearestNeighbor,
	"RIGHT":             imagetk.RIGHT,
	"TOP":               imagetk.TOP,
	"TOP_LEFT":          imagetk.TOP_LEFT,
	"TOP_RIGHT":         imagetk.TOP_RIGHT,

	"ITKX": imagetk.ITKX,

	"ImageTK":    qlang.StructOf((*imagetk.ImageTK)(nil)),
	"imagetk":    imagetk.NewImageTK,
	"NewImageTK": imagetk.NewImageTK,
}
