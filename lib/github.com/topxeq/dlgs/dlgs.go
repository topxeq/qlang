package dlgs

import (
	"github.com/topxeq/dlgs"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/dlgs",

	"ErrNotImplemented": dlgs.ErrNotImplemented,
	"ErrUnsupported":    dlgs.ErrUnsupported,

	"Color":      dlgs.Color,
	"Date":       dlgs.Date,
	"Entry":      dlgs.Entry,
	"Error":      dlgs.Error,
	"File":       dlgs.File,
	"FileMulti":  dlgs.FileMulti,
	"Info":       dlgs.Info,
	"List":       dlgs.List,
	"ListMulti":  dlgs.ListMulti,
	"MessageBox": dlgs.MessageBox,
	"Password":   dlgs.Password,
	"Question":   dlgs.Question,
	"Warning":    dlgs.Warning,
}
