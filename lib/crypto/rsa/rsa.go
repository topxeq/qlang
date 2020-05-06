package rsa

import (
	"crypto/rsa"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/rsa",

	"PSSSaltLengthAuto":       rsa.PSSSaltLengthAuto,
	"PSSSaltLengthEqualsHash": rsa.PSSSaltLengthEqualsHash,

	"ErrDecryption":     rsa.ErrDecryption,
	"ErrMessageTooLong": rsa.ErrMessageTooLong,
	"ErrVerification":   rsa.ErrVerification,

	"decryptOAEP":               rsa.DecryptOAEP,
	"decryptPKCS1v15":           rsa.DecryptPKCS1v15,
	"decryptPKCS1v15SessionKey": rsa.DecryptPKCS1v15SessionKey,
	"encryptOAEP":               rsa.EncryptOAEP,
	"encryptPKCS1v15":           rsa.EncryptPKCS1v15,
	"signPKCS1v15":              rsa.SignPKCS1v15,
	"signPSS":                   rsa.SignPSS,
	"verifyPKCS1v15":            rsa.VerifyPKCS1v15,
	"verifyPSS":                 rsa.VerifyPSS,

	"CRTValue":               qlang.StructOf((*rsa.CRTValue)(nil)),
	"OAEPOptions":            qlang.StructOf((*rsa.OAEPOptions)(nil)),
	"PKCS1v15DecryptOptions": qlang.StructOf((*rsa.PKCS1v15DecryptOptions)(nil)),
	"PSSOptions":             qlang.StructOf((*rsa.PSSOptions)(nil)),
	"PrecomputedValues":      qlang.StructOf((*rsa.PrecomputedValues)(nil)),
	"PrivateKey":             qlang.StructOf((*rsa.PrivateKey)(nil)),
	"generateKey":            rsa.GenerateKey,
	"generateMultiPrimeKey":  rsa.GenerateMultiPrimeKey,
	"PublicKey":              qlang.StructOf((*rsa.PublicKey)(nil)),
}
