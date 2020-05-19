package jsonrpc

import (
	"net/rpc/jsonrpc"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/rpc/jsonrpc",

	"Dial":           jsonrpc.Dial,
	"NewClient":      jsonrpc.NewClient,
	"NewClientCodec": jsonrpc.NewClientCodec,
	"NewServerCodec": jsonrpc.NewServerCodec,
	"ServeConn":      jsonrpc.ServeConn,
}
