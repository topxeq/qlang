package fs

import (
	"io/fs"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "io/fs",

	"ModeAppend":     fs.ModeAppend,
	"ModeCharDevice": fs.ModeCharDevice,
	"ModeDevice":     fs.ModeDevice,
	"ModeDir":        fs.ModeDir,
	"ModeExclusive":  fs.ModeExclusive,
	"ModeIrregular":  fs.ModeIrregular,
	"ModeNamedPipe":  fs.ModeNamedPipe,
	"ModePerm":       fs.ModePerm,
	"ModeSetgid":     fs.ModeSetgid,
	"ModeSetuid":     fs.ModeSetuid,
	"ModeSocket":     fs.ModeSocket,
	"ModeSticky":     fs.ModeSticky,
	"ModeSymlink":    fs.ModeSymlink,
	"ModeTemporary":  fs.ModeTemporary,
	"ModeType":       fs.ModeType,

	"ErrClosed":     fs.ErrClosed,
	"ErrExist":      fs.ErrExist,
	"ErrInvalid":    fs.ErrInvalid,
	"ErrNotExist":   fs.ErrNotExist,
	"ErrPermission": fs.ErrPermission,
	"SkipDir":       fs.SkipDir,

	"Glob":      fs.Glob,
	"ReadFile":  fs.ReadFile,
	"ValidPath": fs.ValidPath,
	"WalkDir":   fs.WalkDir,

	"ReadDir": fs.ReadDir,
	"Sub":     fs.Sub,
	"Stat":    fs.Stat,
}
