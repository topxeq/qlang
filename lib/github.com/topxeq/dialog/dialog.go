package dialog

import (
	"github.com/topxeq/dialog"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/dialog",

	"Cancelled":    dialog.Cancelled,
	"ErrCancelled": dialog.ErrCancelled,

	"DirectoryBuilder": qlang.StructOf((*dialog.DirectoryBuilder)(nil)),
	"Directory":        dialog.Directory,
	"Dlg":              qlang.StructOf((*dialog.Dlg)(nil)),
	"FileBuilder":      qlang.StructOf((*dialog.FileBuilder)(nil)),
	"File":             dialog.File,
	"FileFilter":       qlang.StructOf((*dialog.FileFilter)(nil)),
	"MsgBuilder":       qlang.StructOf((*dialog.MsgBuilder)(nil)),
	"Message":          dialog.Message,
}
