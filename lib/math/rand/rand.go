package rand

import (
	"math/rand"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "math/rand",

	"ExpFloat64":  rand.ExpFloat64,
	"Float32":     rand.Float32,
	"Float64":     rand.Float64,
	"Int":         rand.Int,
	"Int31":       rand.Int31,
	"Int31n":      rand.Int31n,
	"Int63":       rand.Int63,
	"Int63n":      rand.Int63n,
	"Intn":        rand.Intn,
	"NormFloat64": rand.NormFloat64,
	"Perm":        rand.Perm,
	"Read":        rand.Read,
	"Seed":        rand.Seed,
	"Shuffle":     rand.Shuffle,
	"Uint32":      rand.Uint32,
	"Uint64":      rand.Uint64,

	"source":    rand.NewSource,
	"NewSource": rand.NewSource,

	"New":     rand.New,
	"zipf":    rand.NewZipf,
	"NewZipf": rand.NewZipf,
}
