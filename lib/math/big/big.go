package big

import (
	"math/big"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "math/big",

	"Above":         big.Above,
	"AwayFromZero":  big.AwayFromZero,
	"Below":         big.Below,
	"Exact":         big.Exact,
	"MaxBase":       big.MaxBase,
	"MaxExp":        big.MaxExp,
	"MaxPrec":       big.MaxPrec,
	"MinExp":        big.MinExp,
	"ToNearestAway": big.ToNearestAway,
	"ToNearestEven": big.ToNearestEven,
	"ToNegativeInf": big.ToNegativeInf,
	"ToPositiveInf": big.ToPositiveInf,
	"ToZero":        big.ToZero,

	"Jacobi":     big.Jacobi,
	"ParseFloat": big.ParseFloat,

	"float": big.NewFloat,
	"int":   big.NewInt,
	"rat":   big.NewRat,

	"NewFloat": big.NewFloat,
	"NewInt":   big.NewInt,
	"NewRat":   big.NewRat,
}
