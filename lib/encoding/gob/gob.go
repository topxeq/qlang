package gob

import (
	"encoding/gob"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "encoding/gob",

	"Register":     gob.Register,
	"RegisterName": gob.RegisterName,

	"CommonType": qlang.StructOf((*gob.CommonType)(nil)),
	"decoder":    gob.NewDecoder,
	"encoder":    gob.NewEncoder,
}
