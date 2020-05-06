package xml

import (
	"encoding/xml"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "encoding/xml",

	"Header": xml.Header,

	"HTMLAutoClose": xml.HTMLAutoClose,
	"HTMLEntity":    xml.HTMLEntity,

	"escape":        xml.Escape,
	"escapeText":    xml.EscapeText,
	"marshal":       xml.Marshal,
	"marshalIndent": xml.MarshalIndent,
	"unmarshal":     xml.Unmarshal,

	"copyToken": xml.CopyToken,

	"Attr":            qlang.StructOf((*xml.Attr)(nil)),
	"decoder":         xml.NewDecoder,
	"newTokenDecoder": xml.NewTokenDecoder,
	"encoder":         xml.NewEncoder,
	"EndElement":      qlang.StructOf((*xml.EndElement)(nil)),
	"Name":            qlang.StructOf((*xml.Name)(nil)),
	"ProcInst":        qlang.StructOf((*xml.ProcInst)(nil)),
	"StartElement":    qlang.StructOf((*xml.StartElement)(nil)),
}
