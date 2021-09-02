package charlang

import (
	"github.com/topxeq/charlang"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/charlang",

	"NewChar":     charlang.NewChar,
	"RunChar":     charlang.RunChar,
	"RunCharCode": charlang.NewChar,

	"NewAny":      charlang.NewAny,
	"NewAnyValue": charlang.NewAnyValue,
	"MssToMap":    charlang.MssToMap,
	"MsiToMap":    charlang.MsiToMap,

	// "SqlTK": qlang.StructOf((*sqltk.SqlTK)(nil)),
}
