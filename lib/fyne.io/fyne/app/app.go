package app

import (
	"fyne.io/fyne/app"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "fyne.io/fyne/app",

	"New":              app.New,
	"NewAppWithDriver": app.NewAppWithDriver,
	"NewWithID":        app.NewWithID,

	"SettingsSchema": qlang.StructOf((*app.SettingsSchema)(nil)),
}
