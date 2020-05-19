package rpc

import (
	"net/rpc"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/rpc",

	"DefaultDebugPath": rpc.DefaultDebugPath,
	"DefaultRPCPath":   rpc.DefaultRPCPath,

	"DefaultServer": rpc.DefaultServer,
	"ErrShutdown":   rpc.ErrShutdown,

	"Accept":       rpc.Accept,
	"HandleHTTP":   rpc.HandleHTTP,
	"Register":     rpc.Register,
	"RegisterName": rpc.RegisterName,
	"ServeCodec":   rpc.ServeCodec,
	"ServeConn":    rpc.ServeConn,
	"ServeRequest": rpc.ServeRequest,

	"Call":               qlang.StructOf((*rpc.Call)(nil)),
	"NewClient":          rpc.NewClient,
	"NewClientWithCodec": rpc.NewClientWithCodec,
	"Dial":               rpc.Dial,
	"DialHTTP":           rpc.DialHTTP,
	"DialHTTPPath":       rpc.DialHTTPPath,
	"server":             rpc.NewServer,
}
