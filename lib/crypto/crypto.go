package crypto

import (
	"crypto"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto",

	"BLAKE2b_256": crypto.BLAKE2b_256,
	"BLAKE2b_384": crypto.BLAKE2b_384,
	"BLAKE2b_512": crypto.BLAKE2b_512,
	"BLAKE2s_256": crypto.BLAKE2s_256,
	"MD4":         crypto.MD4,
	"MD5":         crypto.MD5,
	"MD5SHA1":     crypto.MD5SHA1,
	"RIPEMD160":   crypto.RIPEMD160,
	"SHA1":        crypto.SHA1,
	"SHA224":      crypto.SHA224,
	"SHA256":      crypto.SHA256,
	"SHA384":      crypto.SHA384,
	"SHA3_224":    crypto.SHA3_224,
	"SHA3_256":    crypto.SHA3_256,
	"SHA3_384":    crypto.SHA3_384,
	"SHA3_512":    crypto.SHA3_512,
	"SHA512":      crypto.SHA512,
	"SHA512_224":  crypto.SHA512_224,
	"SHA512_256":  crypto.SHA512_256,

	"registerHash": crypto.RegisterHash,
}
