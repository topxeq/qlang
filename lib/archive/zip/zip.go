package zip

import (
	"archive/zip"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "archive/zip",

	"Deflate": zip.Deflate,
	"Store":   zip.Store,

	"ErrAlgorithm": zip.ErrAlgorithm,
	"ErrChecksum":  zip.ErrChecksum,
	"ErrFormat":    zip.ErrFormat,

	"RegisterCompressor":   zip.RegisterCompressor,
	"RegisterDecompressor": zip.RegisterDecompressor,

	"FileHeader":     qlang.StructOf((*zip.FileHeader)(nil)),
	"FileInfoHeader": zip.FileInfoHeader,
	"OpenReader":     zip.OpenReader,
	"reader":         zip.NewReader,
	"writer":         zip.NewWriter,
}
