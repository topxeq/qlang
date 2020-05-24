package imagetk

import (
	"github.com/topxeq/imagetk"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/imagetk",

	"ImageTK":    qlang.StructOf((*imagetk.ImageTK)(nil)),
	"imagetk":    imagetk.NewImageTK,
	"NewImageTK": imagetk.NewImageTK,
}
