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
	"QueryDBNSV":         sqltk.QueryDBNSV,
	"QueryDBS":           sqltk.QueryDBS,
	"QueryDBString":      sqltk.QueryDBString,
}
