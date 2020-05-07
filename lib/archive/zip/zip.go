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

	"registerCompressor":   zip.RegisterCompressor,
	"registerDecompressor": zip.RegisterDecompressor,

	"FileHeader":     qlang.StructOf((*zip.FileHeader)(nil)),
	"fileInfoHeader": zip.FileInfoHeader,
	"openReader":     zip.OpenReader,
	"reader":         zip.NewReader,
	"newReader":      zip.NewReader,
	"writer":         zip.NewWriter,
	"newWriter":      zip.NewWriter,
}
