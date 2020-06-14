package blink

import (
	"github.com/topxeq/blink"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/blink",

	"TempPath": blink.TempPath,

	"InitBlink":            blink.InitBlink,
	"LoadIconFromBytes":    blink.LoadIconFromBytes,
	"LoadIconFromFile":     blink.LoadIconFromFile,
	"RegisterFileSystem":   blink.RegisterFileSystem,
	"SetDebugMode":         blink.SetDebugMode,
	"UnregisterFileSystem": blink.UnregisterFileSystem,

	"webview":    blink.NewWebView,
	"NewWebView": blink.NewWebView,
}
