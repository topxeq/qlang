package afero

import (
	"github.com/topxeq/afero"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/afero",

	"BADFD":             afero.BADFD,
	"FilePathSeparator": afero.FilePathSeparator,

	"ErrDestinationExists": afero.ErrDestinationExists,
	"ErrFileClosed":        afero.ErrFileClosed,
	"ErrFileExists":        afero.ErrFileExists,
	"ErrFileNotFound":      afero.ErrFileNotFound,
	"ErrNoReadlink":        afero.ErrNoReadlink,
	"ErrNoSymlink":         afero.ErrNoSymlink,
	"ErrOutOfRange":        afero.ErrOutOfRange,
	"ErrTooLarge":          afero.ErrTooLarge,

	"DirExists":            afero.DirExists,
	"Exists":               afero.Exists,
	"FileContainsAnyBytes": afero.FileContainsAnyBytes,
	"FileContainsBytes":    afero.FileContainsBytes,
	"FullBaseFsPath":       afero.FullBaseFsPath,
	"GetTempDir":           afero.GetTempDir,
	"Glob":                 afero.Glob,
	"IsDir":                afero.IsDir,
	"IsEmpty":              afero.IsEmpty,
	"NeuterAccents":        afero.NeuterAccents,
	"ReadAll":              afero.ReadAll,
	"ReadDir":              afero.ReadDir,
	"ReadFile":             afero.ReadFile,
	"SafeWriteReader":      afero.SafeWriteReader,
	"TempDir":              afero.TempDir,
	"UnicodeSanitize":      afero.UnicodeSanitize,
	"Walk":                 afero.Walk,
	"WriteFile":            afero.WriteFile,
	"WriteReader":          afero.WriteReader,

	"TempFile":         afero.TempFile,
	"NewBasePathFs":    afero.NewBasePathFs,
	"NewCacheOnReadFs": afero.NewCacheOnReadFs,
	"NewCopyOnWriteFs": afero.NewCopyOnWriteFs,
	"NewMemMapFs":      afero.NewMemMapFs,
	"NewOsFs":          afero.NewOsFs,
	"NewReadOnlyFs":    afero.NewReadOnlyFs,
	"NewRegexpFs":      afero.NewRegexpFs,

	"Afero":     qlang.StructOf((*afero.Afero)(nil)),
	"httpfs":    afero.NewHttpFs,
	"NewHttpFs": afero.NewHttpFs,
	"OsFs":      qlang.StructOf((*afero.OsFs)(nil)),
}
