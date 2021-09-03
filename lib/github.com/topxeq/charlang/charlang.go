package charlang

import (
	"github.com/topxeq/charlang"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/charlang",

	"NewChar":         charlang.NewChar,
	"RunChar":         charlang.RunChar,
	"RunCharCode":     charlang.NewChar,
	"ConvertToObject": charlang.ConvertToObject,

	"NewAny":      charlang.NewAny,
	"NewAnyValue": charlang.NewAnyValue,
	"MssToMap":    charlang.MssToMap,
	"MsiToMap":    charlang.MsiToMap,

	"Map":   qlang.StructOf((*charlang.Map)(nil)),
	"Array": qlang.StructOf((*charlang.Array)(nil)),
}
