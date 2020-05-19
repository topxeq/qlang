package fnv

import (
	"hash/fnv"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "hash/fnv",

	"New128":  fnv.New128,
	"New128a": fnv.New128a,
	"New32":   fnv.New32,
	"New32a":  fnv.New32a,
	"New64":   fnv.New64,
	"New64a":  fnv.New64a,
}
