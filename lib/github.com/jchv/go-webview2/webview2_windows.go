//go:build windows
// +build windows

package webview2

import (
	"github.com/jchv/go-webview2"
	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
var Exports = map[string]interface{}{
	"_name": "github.com/jchv/go-webview2",

	"HintFixed": webview2.HintFixed,
	"HintMax":   webview2.HintMax,
	"HintMin":   webview2.HintMin,
	"HintNone":  webview2.HintNone,

	"New":            webview2.New,
	"NewWindow":      webview2.NewWindow,
	"NewWithOptions": webview2.NewWithOptions,

	"WindowOptions":  qlang.StructOf((*webview2.WindowOptions)(nil)),
	"WebViewOptions": qlang.StructOf((*webview2.WebViewOptions)(nil)),
}
