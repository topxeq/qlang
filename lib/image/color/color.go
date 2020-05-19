package color

import (
	"image/color"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "image/color",

	"Alpha16Model": color.Alpha16Model,
	"AlphaModel":   color.AlphaModel,
	"Black":        color.Black,
	"CMYKModel":    color.CMYKModel,
	"Gray16Model":  color.Gray16Model,
	"GrayModel":    color.GrayModel,
	"NRGBA64Model": color.NRGBA64Model,
	"NRGBAModel":   color.NRGBAModel,
	"NYCbCrAModel": color.NYCbCrAModel,
	"Opaque":       color.Opaque,
	"RGBA64Model":  color.RGBA64Model,
	"RGBAModel":    color.RGBAModel,
	"Transparent":  color.Transparent,
	"White":        color.White,
	"YCbCrModel":   color.YCbCrModel,

	"CMYKToRGB":  color.CMYKToRGB,
	"RGBToCMYK":  color.RGBToCMYK,
	"RGBToYCbCr": color.RGBToYCbCr,
	"YCbCrToRGB": color.YCbCrToRGB,

	"ModelFunc": color.ModelFunc,

	"Alpha":   qlang.StructOf((*color.Alpha)(nil)),
	"Alpha16": qlang.StructOf((*color.Alpha16)(nil)),
	"CMYK":    qlang.StructOf((*color.CMYK)(nil)),
	"Gray":    qlang.StructOf((*color.Gray)(nil)),
	"Gray16":  qlang.StructOf((*color.Gray16)(nil)),
	"NRGBA":   qlang.StructOf((*color.NRGBA)(nil)),
	"NRGBA64": qlang.StructOf((*color.NRGBA64)(nil)),
	"NYCbCrA": qlang.StructOf((*color.NYCbCrA)(nil)),
	"RGBA":    qlang.StructOf((*color.RGBA)(nil)),
	"RGBA64":  qlang.StructOf((*color.RGBA64)(nil)),
	"YCbCr":   qlang.StructOf((*color.YCbCr)(nil)),
}
