package smtp

import (
	"net/smtp"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/smtp",

	"SendMail": smtp.SendMail,

	"CRAMMD5Auth": smtp.CRAMMD5Auth,
	"PlainAuth":   smtp.PlainAuth,

	"client":     smtp.NewClient,
	"Dial":       smtp.Dial,
	"ServerInfo": qlang.StructOf((*smtp.ServerInfo)(nil)),
}
