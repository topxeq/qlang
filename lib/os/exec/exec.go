package exec

import (
	"os/exec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "os/exec",

	"ErrNotFound": exec.ErrNotFound,

	"LookPath": exec.LookPath,

	"Command":        exec.Command,
	"CommandContext": exec.CommandContext,
}
