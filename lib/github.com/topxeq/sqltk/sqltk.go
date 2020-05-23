package sqltk

import (
	"github.com/topxeq/sqltk"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/sqltk",

	"ConnectDB":          sqltk.ConnectDB,
	"ConnectDBNoPing":    sqltk.ConnectDBNoPing,
	"ExecV":              sqltk.ExecV,
	"OneLineRecordToMap": sqltk.OneLineRecordToMap,
	"QueryDBCount":       sqltk.QueryDBCount,
	"QueryDBI":           sqltk.QueryDBI,
	"QueryDBNS":          sqltk.QueryDBNS,
	"QueryDBNSS":         sqltk.QueryDBNSS,
	"QueryDBS":           sqltk.QueryDBS,
	"QueryDBString":      sqltk.QueryDBString,

	"connectDB":          sqltk.ConnectDB,
	"connectDBNoPing":    sqltk.ConnectDBNoPing,
	"execV":              sqltk.ExecV,
	"oneLineRecordToMap": sqltk.OneLineRecordToMap,
	"queryDBCount":       sqltk.QueryDBCount,
	"queryDBI":           sqltk.QueryDBI,
	"queryDBNS":          sqltk.QueryDBNS,
	"queryDBNSS":         sqltk.QueryDBNSS,
	"queryDBS":           sqltk.QueryDBS,
	"queryDBString":      sqltk.QueryDBString,
}
