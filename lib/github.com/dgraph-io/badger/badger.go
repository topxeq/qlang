package badger

import (
	"github.com/dgraph-io/badger"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/dgraph-io/badger",

	"ManifestFilename":    badger.ManifestFilename,
	"ValueThresholdLimit": badger.ValueThresholdLimit,

	"DefaultIteratorOptions": badger.DefaultIteratorOptions,
	"ErrBlockedWrites":       badger.ErrBlockedWrites,
	"ErrConflict":            badger.ErrConflict,
	"ErrDiscardedTxn":        badger.ErrDiscardedTxn,
	"ErrEmptyKey":            badger.ErrEmptyKey,
	"ErrInvalidDump":         badger.ErrInvalidDump,
	"ErrInvalidKey":          badger.ErrInvalidKey,
	"ErrInvalidLoadingMode":  badger.ErrInvalidLoadingMode,
	"ErrInvalidRequest":      badger.ErrInvalidRequest,
	"ErrKeyNotFound":         badger.ErrKeyNotFound,
	"ErrManagedTxn":          badger.ErrManagedTxn,
	"ErrNilCallback":         badger.ErrNilCallback,
	"ErrNoRewrite":           badger.ErrNoRewrite,
	"ErrReadOnlyTxn":         badger.ErrReadOnlyTxn,
	"ErrRejected":            badger.ErrRejected,
	"ErrReplayNeeded":        badger.ErrReplayNeeded,
	"ErrRetry":               badger.ErrRetry,
	"ErrThresholdZero":       badger.ErrThresholdZero,
	"ErrTruncateNeeded":      badger.ErrTruncateNeeded,
	"ErrTxnTooBig":           badger.ErrTxnTooBig,
	// "ErrUnsortedKey":         badger.ErrUnsortedKey,
	"ErrValueLogSize": badger.ErrValueLogSize,
	// "ErrValueThreshold":      badger.ErrValueThreshold,
	"ErrWindowsNotSupported": badger.ErrWindowsNotSupported,
	"ErrZeroBandwidth":       badger.ErrZeroBandwidth,

	"Open":               badger.Open,
	"OpenManaged":        badger.OpenManaged,
	"entry":              badger.NewEntry,
	"Manifest":           qlang.StructOf((*badger.Manifest)(nil)),
	"ReplayManifestFile": badger.ReplayManifestFile,
	"DefaultOptions":     badger.DefaultOptions,
	"LSMOnlyOptions":     badger.LSMOnlyOptions,
	"TableInfo":          qlang.StructOf((*badger.TableInfo)(nil)),
	"TableManifest":      qlang.StructOf((*badger.TableManifest)(nil)),
	"IteratorOptions":    qlang.StructOf((*badger.IteratorOptions)(nil)),
}
