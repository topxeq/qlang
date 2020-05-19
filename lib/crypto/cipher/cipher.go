package cipher

import (
	"crypto/cipher"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/cipher",

	"NewGCM":              cipher.NewGCM,
	"NewGCMWithNonceSize": cipher.NewGCMWithNonceSize,
	"NewGCMWithTagSize":   cipher.NewGCMWithTagSize,
	"NewCBCDecrypter":     cipher.NewCBCDecrypter,
	"NewCBCEncrypter":     cipher.NewCBCEncrypter,
	"NewCFBDecrypter":     cipher.NewCFBDecrypter,
	"NewCFBEncrypter":     cipher.NewCFBEncrypter,
	"NewCTR":              cipher.NewCTR,
	"NewOFB":              cipher.NewOFB,

	"StreamReader": qlang.StructOf((*cipher.StreamReader)(nil)),
	"StreamWriter": qlang.StructOf((*cipher.StreamWriter)(nil)),
}
