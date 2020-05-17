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

	"Abs":          filepath.Abs,
	"Base":         filepath.Base,
	"Clean":        filepath.Clean,
	"Dir":          filepath.Dir,
	"EvalSymlinks": filepath.EvalSymlinks,
	"Ext":          filepath.Ext,
	"FromSlash":    filepath.FromSlash,
	"Glob":         filepath.Glob,
	"HasPrefix":    filepath.HasPrefix,
	"IsAbs":        filepath.IsAbs,
	"join":         filepath.Join,
	"Join":         filepath.Join,
	"Match":        filepath.Match,
	"Rel":          filepath.Rel,
	"Split":        filepath.Split,
	"SplitList":    filepath.SplitList,
	"ToSlash":      filepath.ToSlash,
	"VolumeName":   filepath.VolumeName,
	"Walk":         filepath.Walk,
}
