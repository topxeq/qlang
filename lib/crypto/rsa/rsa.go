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

	"DecryptOAEP":               rsa.DecryptOAEP,
	"DecryptPKCS1v15":           rsa.DecryptPKCS1v15,
	"DecryptPKCS1v15SessionKey": rsa.DecryptPKCS1v15SessionKey,
	"EncryptOAEP":               rsa.EncryptOAEP,
	"EncryptPKCS1v15":           rsa.EncryptPKCS1v15,
	"SignPKCS1v15":              rsa.SignPKCS1v15,
	"SignPSS":                   rsa.SignPSS,
	"VerifyPKCS1v15":            rsa.VerifyPKCS1v15,
	"VerifyPSS":                 rsa.VerifyPSS,

	"CRTValue":               qlang.StructOf((*rsa.CRTValue)(nil)),
	"OAEPOptions":            qlang.StructOf((*rsa.OAEPOptions)(nil)),
	"PKCS1v15DecryptOptions": qlang.StructOf((*rsa.PKCS1v15DecryptOptions)(nil)),
	"PSSOptions":             qlang.StructOf((*rsa.PSSOptions)(nil)),
	"PrecomputedValues":      qlang.StructOf((*rsa.PrecomputedValues)(nil)),
	"PrivateKey":             qlang.StructOf((*rsa.PrivateKey)(nil)),
	"GenerateKey":            rsa.GenerateKey,
	"GenerateMultiPrimeKey":  rsa.GenerateMultiPrimeKey,
	"PublicKey":              qlang.StructOf((*rsa.PublicKey)(nil)),
}
