package multipart

import (
	"mime/multipart"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name":     "mime/multipart",
	"NewWriter": multipart.NewWriter,
	"NewReader": multipart.NewReader,

	"File":       qlang.StructOf((*multipart.File)(nil)),
	"FileHeader": qlang.StructOf((*multipart.FileHeader)(nil)),
	"Form":       qlang.StructOf((*multipart.Form)(nil)),
	"Part":       qlang.StructOf((*multipart.Part)(nil)),
}

// -----------------------------------------------------------------------------
