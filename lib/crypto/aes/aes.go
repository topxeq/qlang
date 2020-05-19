package aes

import (
	"crypto/aes"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/aes",

	"BlockSize": aes.BlockSize,

	"NewCipher": aes.NewCipher,
}
