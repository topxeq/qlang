package charlang

import (
	"github.com/topxeq/charlang"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/charlang",

	"NewChar":           charlang.NewChar,
	"RunChar":           charlang.RunChar,
	"RunCharCode":       charlang.NewChar,
	"QuickCompile":      charlang.QuickCompile,
	"QuickRun":          charlang.QuickRun,
	"ConvertToObject":   charlang.ConvertToObject,
	"ConvertFromObject": charlang.ConvertFromObject,

	"NewAny":      charlang.NewAny,
	"NewAnyValue": charlang.NewAnyValue,
	"MssToMap":    charlang.MssToMap,
	"MsiToMap":    charlang.MsiToMap,
	"ObjectsToS":  charlang.ObjectsToS,
	"ObjectsToI":  charlang.ObjectsToI,

	"TkFunction": charlang.TkFunction,

	"VersionG": charlang.VersionG,

	"Byte":  qlang.StructOf((*charlang.Byte)(nil)),
	"Any":   qlang.StructOf((*charlang.Any)(nil)),
	"Map":   qlang.StructOf((*charlang.Map)(nil)),
	"Array": qlang.StructOf((*charlang.Array)(nil)),
}
