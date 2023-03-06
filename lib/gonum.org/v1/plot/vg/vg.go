package vg

import (
	"gonum.org/v1/plot/vg"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
var Exports = map[string]interface{}{
	"_name": "gonum.org/v1/plot/vg",

	"ArcComp":    vg.ArcComp,
	"Centimeter": vg.Centimeter,
	"CloseComp":  vg.CloseComp,
	"CurveComp":  vg.CurveComp,
	"Inch":       vg.Inch,
	"LineComp":   vg.LineComp,
	"Millimeter": vg.Millimeter,
	"MoveComp":   vg.MoveComp,

	// "FontDirs": vg.FontDirs,
	// "FontMap":  vg.FontMap,

	// "AddFont":    vg.AddFont,
	"Initialize": vg.Initialize,

	// "MakeFont":    vg.MakeFont,
	// "FontExtents": qlang.StructOf((*vg.FontExtents)(nil)),
	"PathComp":  qlang.StructOf((*vg.PathComp)(nil)),
	"Point":     qlang.StructOf((*vg.Point)(nil)),
	"Rectangle": qlang.StructOf((*vg.Rectangle)(nil)),
}
