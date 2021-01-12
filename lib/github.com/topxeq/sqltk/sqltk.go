package sqltk

import (
	"github.com/topxeq/sqltk"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/sqltk",

	"ConnectDB":          sqltk.ConnectDB,
	"ConnectDBNoPing":    sqltk.ConnectDBNoPing,
	"ExecV":              sqltk.ExecV,
	"FormatSQLValue":     sqltk.FormatSQLValue,
	"GetVersion":         sqltk.GetVersion,
	"NewSqlTK":           sqltk.NewSqlTK,
	"OneLineRecordToMap": sqltk.OneLineRecordToMap,
	"QueryDBCount":       sqltk.QueryDBCount,
	"QueryDBI":           sqltk.QueryDBI,
	"QueryDBNS":          sqltk.QueryDBNS,
	"QueryDBNSS":         sqltk.QueryDBNSS,
	"QueryDBNSSF":        sqltk.QueryDBNSSF,
	"QueryDBNSV":         sqltk.QueryDBNSV,
	"QueryDBS":           sqltk.QueryDBS,
	"QueryDBString":      sqltk.QueryDBString,
	"RecordsToMapArray":  sqltk.RecordsToMapArray,
	"SqlTKX":             sqltk.SqlTKX,

	"SqlTK": qlang.StructOf((*sqltk.SqlTK)(nil)),
}
