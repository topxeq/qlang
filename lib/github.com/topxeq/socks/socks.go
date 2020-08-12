package socks

import (
	"github.com/topxeq/socks"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/socks",

	"OpenLocalTunnel":  socks.OpenLocalTunnel,
	"OpenRemoteTunnel": socks.OpenRemoteTunnel,

	"CreateSock": socks.CreateSock,
}
