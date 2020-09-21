package canvas

import (
	"fyne.io/fyne/canvas"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "fyne.io/fyne/canvas",

	"ImageFillContain":  canvas.ImageFillContain,
	"ImageFillOriginal": canvas.ImageFillOriginal,
	"ImageFillStretch":  canvas.ImageFillStretch,
	"ImageScalePixels":  canvas.ImageScalePixels,
	"ImageScaleSmooth":  canvas.ImageScaleSmooth,

	"Refresh": canvas.Refresh,

	"Circle":                qlang.StructOf((*canvas.Circle)(nil)),
	"circle":                canvas.NewCircle,
	"NewCircle":             canvas.NewCircle,
	"Image":                 qlang.StructOf((*canvas.Image)(nil)),
	"NewImageFromFile":      canvas.NewImageFromFile,
	"NewImageFromImage":     canvas.NewImageFromImage,
	"NewImageFromResource":  canvas.NewImageFromResource,
	"Line":                  qlang.StructOf((*canvas.Line)(nil)),
	"line":                  canvas.NewLine,
	"NewLine":               canvas.NewLine,
	"LinearGradient":        qlang.StructOf((*canvas.LinearGradient)(nil)),
	"lineargradient":        canvas.NewLinearGradient,
	"NewLinearGradient":     canvas.NewLinearGradient,
	"NewHorizontalGradient": canvas.NewHorizontalGradient,
	"NewVerticalGradient":   canvas.NewVerticalGradient,
	"RadialGradient":        qlang.StructOf((*canvas.RadialGradient)(nil)),
	"radialgradient":        canvas.NewRadialGradient,
	"NewRadialGradient":     canvas.NewRadialGradient,
	"NewRaster":             canvas.NewRaster,
	"NewRasterFromImage":    canvas.NewRasterFromImage,
	"NewRasterWithPixels":   canvas.NewRasterWithPixels,
	"Rectangle":             qlang.StructOf((*canvas.Rectangle)(nil)),
	"rectangle":             canvas.NewRectangle,
	"NewRectangle":          canvas.NewRectangle,
	"Text":                  qlang.StructOf((*canvas.Text)(nil)),
	"text":                  canvas.NewText,
	"NewText":               canvas.NewText,
}
