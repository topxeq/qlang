package filepath

import (
	"path/filepath"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "path/filepath",

	"ListSeparator": filepath.ListSeparator,
	"Separator":     filepath.Separator,

	"ErrBadPattern": filepath.ErrBadPattern,
	"SkipDir":       filepath.SkipDir,

	"abs":          filepath.Abs,
	"base":         filepath.Base,
	"clean":        filepath.Clean,
	"dir":          filepath.Dir,
	"evalSymlinks": filepath.EvalSymlinks,
	"ext":          filepath.Ext,
	"fromSlash":    filepath.FromSlash,
	"glob":         filepath.Glob,
	"hasPrefix":    filepath.HasPrefix,
	"isAbs":        filepath.IsAbs,
	"join":         filepath.Join,
	"Join":         filepath.Join,
	"match":        filepath.Match,
	"rel":          filepath.Rel,
	"split":        filepath.Split,
	"splitList":    filepath.SplitList,
	"toSlash":      filepath.ToSlash,
	"volumeName":   filepath.VolumeName,
	"walk":         filepath.Walk,
}
