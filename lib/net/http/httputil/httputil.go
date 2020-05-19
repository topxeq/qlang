package httputil

import (
	"net/http/httputil"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/http/httputil",

	"ErrClosed":      httputil.ErrClosed,
	"ErrLineTooLong": httputil.ErrLineTooLong,
	"ErrPersistEOF":  httputil.ErrPersistEOF,
	"ErrPipeline":    httputil.ErrPipeline,

	"DumpRequest":      httputil.DumpRequest,
	"DumpRequestOut":   httputil.DumpRequestOut,
	"DumpResponse":     httputil.DumpResponse,
	"NewChunkedReader": httputil.NewChunkedReader,
	"NewChunkedWriter": httputil.NewChunkedWriter,

	"clientconn":                httputil.NewClientConn,
	"NewProxyClientConn":        httputil.NewProxyClientConn,
	"ReverseProxy":              qlang.StructOf((*httputil.ReverseProxy)(nil)),
	"NewSingleHostReverseProxy": httputil.NewSingleHostReverseProxy,
	"serverconn":                httputil.NewServerConn,
}
