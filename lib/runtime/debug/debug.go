package debug

import (
	"runtime/debug"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "runtime/debug",

	"freeOSMemory":    debug.FreeOSMemory,
	"printStack":      debug.PrintStack,
	"readGCStats":     debug.ReadGCStats,
	"setGCPercent":    debug.SetGCPercent,
	"setMaxStack":     debug.SetMaxStack,
	"setMaxThreads":   debug.SetMaxThreads,
	"setPanicOnFault": debug.SetPanicOnFault,
	"setTraceback":    debug.SetTraceback,
	"stack":           debug.Stack,
	"writeHeapDump":   debug.WriteHeapDump,

	"BuildInfo":     qlang.StructOf((*debug.BuildInfo)(nil)),
	"readBuildInfo": debug.ReadBuildInfo,
	"GCStats":       qlang.StructOf((*debug.GCStats)(nil)),
	"Module":        qlang.StructOf((*debug.Module)(nil)),
}
