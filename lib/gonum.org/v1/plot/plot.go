package plot

import (
	"gonum.org/v1/plot"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "gonum.org/v1/plot",

	"DefaultFont": plot.DefaultFont,
	"UTCUnixTime": plot.UTCUnixTime,

	"Align":      plot.Align,
	"UnixTimeIn": plot.UnixTimeIn,
	"Version":    plot.Version,

	"Axis":          qlang.StructOf((*plot.Axis)(nil)),
	"DefaultTicks":  qlang.StructOf((*plot.DefaultTicks)(nil)),
	"GlyphBox":      qlang.StructOf((*plot.GlyphBox)(nil)),
	"InvertedScale": qlang.StructOf((*plot.InvertedScale)(nil)),
	"legend":        plot.NewLegend,
	"LinearScale":   qlang.StructOf((*plot.LinearScale)(nil)),
	"LogScale":      qlang.StructOf((*plot.LogScale)(nil)),
	"LogTicks":      qlang.StructOf((*plot.LogTicks)(nil)),
	"New":           plot.New,
	"Tick":          qlang.StructOf((*plot.Tick)(nil)),
	"TimeTicks":     qlang.StructOf((*plot.TimeTicks)(nil)),
}
