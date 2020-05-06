package base64

import (
	"encoding/base64"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "encoding/base64",

	"NoPadding":  base64.NoPadding,
	"StdPadding": base64.StdPadding,

	"RawStdEncoding": base64.RawStdEncoding,
	"RawURLEncoding": base64.RawURLEncoding,
	"StdEncoding":    base64.StdEncoding,
	"URLEncoding":    base64.URLEncoding,

	"newDecoder": base64.NewDecoder,
	"newEncoder": base64.NewEncoder,

	"encoding": base64.NewEncoding,
}
