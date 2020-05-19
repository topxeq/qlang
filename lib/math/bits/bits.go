package bits

import (
	"math/bits"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "math/bits",

	"UintSize": bits.UintSize,

	"Add":             bits.Add,
	"Add32":           bits.Add32,
	"Add64":           bits.Add64,
	"Div":             bits.Div,
	"Div32":           bits.Div32,
	"Div64":           bits.Div64,
	"LeadingZeros":    bits.LeadingZeros,
	"LeadingZeros16":  bits.LeadingZeros16,
	"LeadingZeros32":  bits.LeadingZeros32,
	"LeadingZeros64":  bits.LeadingZeros64,
	"LeadingZeros8":   bits.LeadingZeros8,
	"Len":             bits.Len,
	"Len16":           bits.Len16,
	"Len32":           bits.Len32,
	"Len64":           bits.Len64,
	"Len8":            bits.Len8,
	"Mul":             bits.Mul,
	"Mul32":           bits.Mul32,
	"Mul64":           bits.Mul64,
	"OnesCount":       bits.OnesCount,
	"OnesCount16":     bits.OnesCount16,
	"OnesCount32":     bits.OnesCount32,
	"OnesCount64":     bits.OnesCount64,
	"OnesCount8":      bits.OnesCount8,
	"Rem":             bits.Rem,
	"Rem32":           bits.Rem32,
	"Rem64":           bits.Rem64,
	"Reverse":         bits.Reverse,
	"Reverse16":       bits.Reverse16,
	"Reverse32":       bits.Reverse32,
	"Reverse64":       bits.Reverse64,
	"Reverse8":        bits.Reverse8,
	"ReverseBytes":    bits.ReverseBytes,
	"ReverseBytes16":  bits.ReverseBytes16,
	"ReverseBytes32":  bits.ReverseBytes32,
	"ReverseBytes64":  bits.ReverseBytes64,
	"RotateLeft":      bits.RotateLeft,
	"RotateLeft16":    bits.RotateLeft16,
	"RotateLeft32":    bits.RotateLeft32,
	"RotateLeft64":    bits.RotateLeft64,
	"RotateLeft8":     bits.RotateLeft8,
	"Sub":             bits.Sub,
	"Sub32":           bits.Sub32,
	"Sub64":           bits.Sub64,
	"TrailingZeros":   bits.TrailingZeros,
	"TrailingZeros16": bits.TrailingZeros16,
	"TrailingZeros32": bits.TrailingZeros32,
	"TrailingZeros64": bits.TrailingZeros64,
	"TrailingZeros8":  bits.TrailingZeros8,
}
