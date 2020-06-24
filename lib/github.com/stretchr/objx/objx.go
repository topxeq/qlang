package objx

import (
	"github.com/stretchr/objx"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/stretchr/objx",

	"PathSeparator":                objx.PathSeparator,
	"SignatureSeparator":           objx.SignatureSeparator,
	"URLValuesSliceKeySuffixArray": objx.URLValuesSliceKeySuffixArray,
	"URLValuesSliceKeySuffixEmpty": objx.URLValuesSliceKeySuffixEmpty,
	"URLValuesSliceKeySuffixIndex": objx.URLValuesSliceKeySuffixIndex,

	"FromJSON": objx.FromJSON,

	"Nil": objx.Nil,

	"HashWithKey":                objx.HashWithKey,
	"SetURLValuesSliceKeySuffix": objx.SetURLValuesSliceKeySuffix,
}
