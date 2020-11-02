package grab

import (
	"github.com/cavaliercoder/grab"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/cavaliercoder/grab",

	"DefaultClient":  grab.DefaultClient,
	"ErrBadChecksum": grab.ErrBadChecksum,
	"ErrBadLength":   grab.ErrBadLength,
	"ErrFileExists":  grab.ErrFileExists,
	"ErrNoFilename":  grab.ErrNoFilename,
	"ErrNoTimestamp": grab.ErrNoTimestamp,

	"GetBatch":          grab.GetBatch,
	"IsStatusCodeError": grab.IsStatusCodeError,

	"Client":     qlang.StructOf((*grab.Client)(nil)),
	"client":     grab.NewClient,
	"NewClient":  grab.NewClient,
	"request":    grab.NewRequest,
	"NewRequest": grab.NewRequest,
	"Get":        grab.Get,
}
