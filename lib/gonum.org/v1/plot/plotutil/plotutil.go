package plotutil

import (
	"gonum.org/v1/plot/plotutil"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "gonum.org/v1/plot/plotutil",

	"DarkColors":         plotutil.DarkColors,
	"DefaultColors":      plotutil.DefaultColors,
	"DefaultDashes":      plotutil.DefaultDashes,
	"DefaultGlyphShapes": plotutil.DefaultGlyphShapes,
	"SoftColors":         plotutil.SoftColors,

	"AddBoxPlots":         plotutil.AddBoxPlots,
	"AddErrorBars":        plotutil.AddErrorBars,
	"AddLinePoints":       plotutil.AddLinePoints,
	"AddLines":            plotutil.AddLines,
	"AddScatters":         plotutil.AddScatters,
	"AddStackedAreaPlots": plotutil.AddStackedAreaPlots,
	"AddXErrorBars":       plotutil.AddXErrorBars,
	"AddYErrorBars":       plotutil.AddYErrorBars,
	"Color":               plotutil.Color,
	"Dashes":              plotutil.Dashes,
	"MeanAndConf95":       plotutil.MeanAndConf95,
	"MedianAndMinMax":     plotutil.MedianAndMinMax,
	"Shape":               plotutil.Shape,

	"ErrorPoints":    qlang.StructOf((*plotutil.ErrorPoints)(nil)),
	"errorpoints":    plotutil.NewErrorPoints,
	"NewErrorPoints": plotutil.NewErrorPoints,
}
