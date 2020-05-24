package gg

import (
	"github.com/fogleman/gg"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/fogleman/gg",

	"AlignCenter":     gg.AlignCenter,
	"AlignLeft":       gg.AlignLeft,
	"AlignRight":      gg.AlignRight,
	"FillRuleEvenOdd": gg.FillRuleEvenOdd,
	"FillRuleWinding": gg.FillRuleWinding,
	"LineCapButt":     gg.LineCapButt,
	"LineCapRound":    gg.LineCapRound,
	"LineCapSquare":   gg.LineCapSquare,
	"LineJoinBevel":   gg.LineJoinBevel,
	"LineJoinRound":   gg.LineJoinRound,
	"RepeatBoth":      gg.RepeatBoth,
	"RepeatNone":      gg.RepeatNone,
	"RepeatX":         gg.RepeatX,
	"RepeatY":         gg.RepeatY,

	"Degrees":      gg.Degrees,
	"LoadFontFace": gg.LoadFontFace,
	"LoadImage":    gg.LoadImage,
	"LoadJPG":      gg.LoadJPG,
	"LoadPNG":      gg.LoadPNG,
	"Radians":      gg.Radians,
	"SaveJPG":      gg.SaveJPG,
	"SavePNG":      gg.SavePNG,

	"NewLinearGradient": gg.NewLinearGradient,
	"NewRadialGradient": gg.NewRadialGradient,
	"NewSolidPattern":   gg.NewSolidPattern,
	"NewSurfacePattern": gg.NewSurfacePattern,

	"NewContext":         gg.NewContext,
	"NewContextForImage": gg.NewContextForImage,
	"NewContextForRGBA":  gg.NewContextForRGBA,
	"Matrix":             qlang.StructOf((*gg.Matrix)(nil)),
	"Identity":           gg.Identity,
	"Rotate":             gg.Rotate,
	"Scale":              gg.Scale,
	"Shear":              gg.Shear,
	"Translate":          gg.Translate,
	"Point":              qlang.StructOf((*gg.Point)(nil)),
	"CubicBezier":        gg.CubicBezier,
	"QuadraticBezier":    gg.QuadraticBezier,
}
