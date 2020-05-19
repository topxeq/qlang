package image

import (
	"image"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "image",

	"YCbCrSubsampleRatio410": image.YCbCrSubsampleRatio410,
	"YCbCrSubsampleRatio411": image.YCbCrSubsampleRatio411,
	"YCbCrSubsampleRatio420": image.YCbCrSubsampleRatio420,
	"YCbCrSubsampleRatio422": image.YCbCrSubsampleRatio422,
	"YCbCrSubsampleRatio440": image.YCbCrSubsampleRatio440,
	"YCbCrSubsampleRatio444": image.YCbCrSubsampleRatio444,

	"Black":       image.Black,
	"ErrFormat":   image.ErrFormat,
	"Opaque":      image.Opaque,
	"Transparent": image.Transparent,
	"White":       image.White,
	"ZP":          &image.ZP,
	"ZR":          &image.ZR,

	"RegisterFormat": image.RegisterFormat,

	"Decode": image.Decode,

	"Alpha":        qlang.StructOf((*image.Alpha)(nil)),
	"alpha":        image.NewAlpha,
	"Alpha16":      qlang.StructOf((*image.Alpha16)(nil)),
	"alpha16":      image.NewAlpha16,
	"CMYK":         qlang.StructOf((*image.CMYK)(nil)),
	"cmyk":         image.NewCMYK,
	"Config":       qlang.StructOf((*image.Config)(nil)),
	"DecodeConfig": image.DecodeConfig,
	"Gray":         qlang.StructOf((*image.Gray)(nil)),
	"gray":         image.NewGray,
	"Gray16":       qlang.StructOf((*image.Gray16)(nil)),
	"gray16":       image.NewGray16,
	"NRGBA":        qlang.StructOf((*image.NRGBA)(nil)),
	"nrgba":        image.NewNRGBA,
	"NRGBA64":      qlang.StructOf((*image.NRGBA64)(nil)),
	"nrgba64":      image.NewNRGBA64,
	"NYCbCrA":      qlang.StructOf((*image.NYCbCrA)(nil)),
	"nycbcra":      image.NewNYCbCrA,
	"Paletted":     qlang.StructOf((*image.Paletted)(nil)),
	"paletted":     image.NewPaletted,
	"Point":        qlang.StructOf((*image.Point)(nil)),
	"Pt":           image.Pt,
	"RGBA":         qlang.StructOf((*image.RGBA)(nil)),
	"rgba":         image.NewRGBA,
	"RGBA64":       qlang.StructOf((*image.RGBA64)(nil)),
	"rgba64":       image.NewRGBA64,
	"Rectangle":    qlang.StructOf((*image.Rectangle)(nil)),
	"Rect":         image.Rect,
	"Uniform":      qlang.StructOf((*image.Uniform)(nil)),
	"uniform":      image.NewUniform,
	"YCbCr":        qlang.StructOf((*image.YCbCr)(nil)),
	"ycbcr":        image.NewYCbCr,
}
