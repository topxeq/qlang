package sync

import (
	qlang "github.com/topxeq/qlang/spec"
	"sync"
)

// -----------------------------------------------------------------------------

func newMutex() *sync.Mutex {
	return new(sync.Mutex)
}

func newRWMutex() *sync.RWMutex {
	return new(sync.RWMutex)
}

func newWaitGroup() *sync.WaitGroup {
	return new(sync.WaitGroup)
}

// -----------------------------------------------------------------------------

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name":     "sync",
	"cond":      sync.NewCond,
	"mutex":     newMutex,
	"waitGroup": newWaitGroup,

	"WaitGroup":    qlang.StructOf((*sync.WaitGroup)(nil)),
	"Mutex":        qlang.StructOf((*sync.Mutex)(nil)),
	"RWMutex":      qlang.StructOf((*sync.RWMutex)(nil)),
	"Cond":         qlang.StructOf((*sync.Cond)(nil)),
	"Once":         qlang.StructOf((*sync.Once)(nil)),
	"NewCond":      sync.NewCond,
	"NewMutex":     newMutex,
	"NewWaitGroup": newWaitGroup,
}

// -----------------------------------------------------------------------------
