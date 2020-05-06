package rand

import (
	"crypto/rand"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/rand",

	"Reader": rand.Reader,

	"int":   rand.Int,
	"prime": rand.Prime,
	"read":  rand.Read,
}
