package sqltk

import (
	"github.com/topxeq/sqltk"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/sqltk",

	"ConnectDB":               sqltk.ConnectDB,
	"ConnectDBNoPing":         sqltk.ConnectDBNoPing,
	"ExecV":                   sqltk.ExecV,
	"FormatSQLValue":          sqltk.FormatSQLValue,
	"GetVersion":              sqltk.GetVersion,
	"NewSqlTK":                sqltk.NewSqlTK,
	"OneLineRecordToMap":      sqltk.OneLineRecordToMap,
	"OneColumnRecordsToArray": sqltk.OneColumnRecordsToArray,
	"QueryDBCount":            sqltk.QueryDBCount,
	"QueryDBFloat":            sqltk.QueryDBFloat,
	"QueryDBI":                sqltk.QueryDBI,
	"QueryDBNS":               sqltk.QueryDBNS,
	"QueryDBNSS":              sqltk.QueryDBNSS,
	"QueryDBNSSF":             sqltk.QueryDBNSSF,
	"QueryDBNSV":              sqltk.QueryDBNSV,
	"QueryDBS":                sqltk.QueryDBS,
	"QueryDBString":           sqltk.QueryDBString,
	"RecordsToMapArray":       sqltk.RecordsToMapArray,
	"RecordsToMapArrayMap":    sqltk.RecordsToMapArrayMap,

	"ConnectDBX":       sqltk.ConnectDBX,
	"ExecDBX":          sqltk.ExecDBX,
	"QueryDBX":         sqltk.QueryDBX,
	"QueryDBRecsX":     sqltk.QueryDBRecsX,
	"QueryDBMapX":      sqltk.QueryDBMapX,
	"QueryDBMapArrayX": sqltk.QueryDBMapArrayX,
	"QueryCountX":      sqltk.QueryCountX,
	"QueryFloatX":      sqltk.QueryFloatX,
	"QueryStringX":     sqltk.QueryStringX,
	"CloseDBX":         sqltk.CloseDBX,

	"BeginTransX": sqltk.BeginTransX,
	"PrepareX":    sqltk.PrepareX,
	"CommitX":     sqltk.CommitX,

	"SqlTKX": sqltk.SqlTKX,

	"SqlTK": qlang.StructOf((*sqltk.SqlTK)(nil)),
}
