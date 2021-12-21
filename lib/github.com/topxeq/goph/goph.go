package goph

import (
	"github.com/topxeq/goph"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/goph",

	"DefaultTimeout": goph.DefaultTimeout,

	"AddKnownHost":          goph.AddKnownHost,
	"CheckKnownHost":        goph.CheckKnownHost,
	"DefaultKnownHosts":     goph.DefaultKnownHosts,
	"DefaultKnownHostsPath": goph.DefaultKnownHostsPath,
	"Dial":                  goph.Dial,
	"GetSigner":             goph.GetSigner,
	"HasAgent":              goph.HasAgent,
	"KnownHosts":            goph.KnownHosts,

	"Client":     qlang.StructOf((*goph.Client)(nil)),
	"New":        goph.New,
	"NewConn":    goph.NewConn,
	"NewUnknown": goph.NewUnknown,
	"Cmd":        qlang.StructOf((*goph.Cmd)(nil)),
	"Config":     qlang.StructOf((*goph.Config)(nil)),
}
