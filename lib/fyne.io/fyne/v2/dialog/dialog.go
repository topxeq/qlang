package dialog

import (
	"fyne.io/fyne/v2/dialog"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "fyne.io/fyne/v2/dialog",

	"ShowColorPicker":   dialog.ShowColorPicker,
	"ShowConfirm":       dialog.ShowConfirm,
	"ShowCustom":        dialog.ShowCustom,
	"ShowCustomConfirm": dialog.ShowCustomConfirm,
	"ShowEntryDialog":   dialog.ShowEntryDialog,
	"ShowError":         dialog.ShowError,
	"ShowFileOpen":      dialog.ShowFileOpen,
	"ShowFileSave":      dialog.ShowFileSave,
	"ShowFolderOpen":    dialog.ShowFolderOpen,
	"ShowForm":          dialog.ShowForm,
	"ShowInformation":   dialog.ShowInformation,

	"NewCustom":        dialog.NewCustom,
	"NewCustomConfirm": dialog.NewCustomConfirm,
	"NewError":         dialog.NewError,
	"NewForm":          dialog.NewForm,
	"NewInformation":   dialog.NewInformation,

	"NewColorPicker":      dialog.NewColorPicker,
	"NewConfirm":          dialog.NewConfirm,
	"entrydialog":         dialog.NewEntryDialog,
	"NewEntryDialog":      dialog.NewEntryDialog,
	"NewFileOpen":         dialog.NewFileOpen,
	"NewFileSave":         dialog.NewFileSave,
	"NewFolderOpen":       dialog.NewFolderOpen,
	"NewProgress":         dialog.NewProgress,
	"NewProgressInfinite": dialog.NewProgressInfinite,
}
