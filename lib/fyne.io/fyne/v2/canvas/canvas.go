package canvas

import (
	"fyne.io/fyne/v2/canvas"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "fyne.io/fyne/v2/canvas",

	"DurationShort":     canvas.DurationShort,
	"DurationStandard":  canvas.DurationStandard,
	"ImageFillContain":  canvas.ImageFillContain,
	"ImageFillOriginal": canvas.ImageFillOriginal,
	"ImageFillStretch":  canvas.ImageFillStretch,
	"ImageScaleFastest": canvas.ImageScaleFastest,
	"ImageScalePixels":  canvas.ImageScalePixels,
	"ImageScaleSmooth":  canvas.ImageScaleSmooth,

	"NewColorRGBAAnimation": canvas.NewColorRGBAAnimation,
	"NewPositionAnimation":  canvas.NewPositionAnimation,
	"NewSizeAnimation":      canvas.NewSizeAnimation,
	"Refresh":               canvas.Refresh,

	"Circle":                qlang.StructOf((*canvas.Circle)(nil)),
	"circle":                canvas.NewCircle,
	"NewCircle":             canvas.NewCircle,
	"Image":                 qlang.StructOf((*canvas.Image)(nil)),
	"NewImageFromFile":      canvas.NewImageFromFile,
	"NewImageFromImage":     canvas.NewImageFromImage,
	"NewImageFromReader":    canvas.NewImageFromReader,
	"NewImageFromResource":  canvas.NewImageFromResource,
	"NewImageFromURI":       canvas.NewImageFromURI,
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
	"Raster":                qlang.StructOf((*canvas.Raster)(nil)),
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
