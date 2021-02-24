package app

import (
	"fyne.io/fyne/v2/app"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "fyne.io/fyne/v2/app",

	"New":       app.New,
	"NewWithID": app.NewWithID,

	"SettingsSchema": qlang.StructOf((*app.SettingsSchema)(nil)),
}
