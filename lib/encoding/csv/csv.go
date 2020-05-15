package csv

import (
	"encoding/csv"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "encoding/csv",

	"ErrBareQuote":     csv.ErrBareQuote,
	"ErrFieldCount":    csv.ErrFieldCount,
	"ErrQuote":         csv.ErrQuote,
	"ErrTrailingComma": csv.ErrTrailingComma,

	"reader": csv.NewReader,
	"writer": csv.NewWriter,

	"NewWriter": csv.NewWriter,
	"NewReader": csv.NewReader,
}
