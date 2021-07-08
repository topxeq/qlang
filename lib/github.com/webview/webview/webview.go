package webview

import (
	"github.com/webview/webview"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/webview/webview",

	"HintFixed": webview.HintFixed,
	"HintMax":   webview.HintMax,
	"HintMin":   webview.HintMin,
	"HintNone":  webview.HintNone,

	"New":       webview.New,
	"NewWindow": webview.NewWindow,
}
