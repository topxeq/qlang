package sql

import (
	"database/sql"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "database/sql",

	"LevelDefault":         sql.LevelDefault,
	"LevelLinearizable":    sql.LevelLinearizable,
	"LevelReadCommitted":   sql.LevelReadCommitted,
	"LevelReadUncommitted": sql.LevelReadUncommitted,
	"LevelRepeatableRead":  sql.LevelRepeatableRead,
	"LevelSerializable":    sql.LevelSerializable,
	"LevelSnapshot":        sql.LevelSnapshot,
	"LevelWriteCommitted":  sql.LevelWriteCommitted,

	"ErrConnDone": sql.ErrConnDone,
	"ErrNoRows":   sql.ErrNoRows,
	"ErrTxDone":   sql.ErrTxDone,

	"Drivers":  sql.Drivers,
	"Register": sql.Register,

	"Open":        sql.Open,
	"OpenDB":      sql.OpenDB,
	"DBStats":     qlang.StructOf((*sql.DBStats)(nil)),
	"Named":       sql.Named,
	"NullBool":    qlang.StructOf((*sql.NullBool)(nil)),
	"NullFloat64": qlang.StructOf((*sql.NullFloat64)(nil)),
	"NullInt32":   qlang.StructOf((*sql.NullInt32)(nil)),
	"NullInt64":   qlang.StructOf((*sql.NullInt64)(nil)),
	"NullString":  qlang.StructOf((*sql.NullString)(nil)),
	"NullTime":    qlang.StructOf((*sql.NullTime)(nil)),
	"TxOptions":   qlang.StructOf((*sql.TxOptions)(nil)),
}
