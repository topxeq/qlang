package cookiejar

import (
	"net/http/cookiejar"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/http/cookiejar",

	"New":     cookiejar.New,
	"Options": qlang.StructOf((*cookiejar.Options)(nil)),
}
