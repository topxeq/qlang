package mail

import (
	"net/mail"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/mail",

	"ErrHeaderNotPresent": mail.ErrHeaderNotPresent,

	"ParseDate": mail.ParseDate,

	"Address":          qlang.StructOf((*mail.Address)(nil)),
	"ParseAddress":     mail.ParseAddress,
	"ParseAddressList": mail.ParseAddressList,
	"AddressParser":    qlang.StructOf((*mail.AddressParser)(nil)),
	"Message":          qlang.StructOf((*mail.Message)(nil)),
	"ReadMessage":      mail.ReadMessage,
}
