package debug

import (
	"runtime/debug"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "runtime/debug",

	"FreeOSMemory":    debug.FreeOSMemory,
	"PrintStack":      debug.PrintStack,
	"ReadGCStats":     debug.ReadGCStats,
	"SetGCPercent":    debug.SetGCPercent,
	"SetMaxStack":     debug.SetMaxStack,
	"SetMaxThreads":   debug.SetMaxThreads,
	"SetPanicOnFault": debug.SetPanicOnFault,
	"SetTraceback":    debug.SetTraceback,
	"Stack":           debug.Stack,
	"WriteHeapDump":   debug.WriteHeapDump,

	"BuildInfo":     qlang.StructOf((*debug.BuildInfo)(nil)),
	"ReadBuildInfo": debug.ReadBuildInfo,
	"GCStats":       qlang.StructOf((*debug.GCStats)(nil)),
	"Module":        qlang.StructOf((*debug.Module)(nil)),
}
