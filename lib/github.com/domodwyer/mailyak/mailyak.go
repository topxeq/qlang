package mailyak

import (
	"github.com/domodwyer/mailyak"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/domodwyer/mailyak",

	"BodyPart": qlang.StructOf((*mailyak.BodyPart)(nil)),
	"New":      mailyak.New,
}
