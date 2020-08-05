package socks5

import (
	"github.com/bhhbazinga/socks5-go"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/bhhbazinga/socks5-go",

	"OpenLocalTunnel":  socks5.OpenLocalTunnel,
	"OpenRemoteTunnel": socks5.OpenRemoteTunnel,

	"CreateSock": socks5.CreateSock,
}
