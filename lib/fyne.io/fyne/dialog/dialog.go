package dialog

import (
	"fyne.io/fyne/dialog"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "fyne.io/fyne/dialog",

	"NewFileIcon":       dialog.NewFileIcon,
	"ShowConfirm":       dialog.ShowConfirm,
	"ShowCustom":        dialog.ShowCustom,
	"ShowCustomConfirm": dialog.ShowCustomConfirm,
	"ShowError":         dialog.ShowError,
	"ShowFileOpen":      dialog.ShowFileOpen,
	"ShowFileSave":      dialog.ShowFileSave,
	"ShowInformation":   dialog.ShowInformation,

	"NewCustom":        dialog.NewCustom,
	"NewCustomConfirm": dialog.NewCustomConfirm,
	"NewError":         dialog.NewError,
	"NewInformation":   dialog.NewInformation,

	"NewConfirm":          dialog.NewConfirm,
	"NewFileOpen":         dialog.NewFileOpen,
	"NewFileSave":         dialog.NewFileSave,
	"NewProgress":         dialog.NewProgress,
	"NewProgressInfinite": dialog.NewProgressInfinite,
}
