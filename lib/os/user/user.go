package user

import (
	"os/user"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "os/user",

	"Group":         qlang.StructOf((*user.Group)(nil)),
	"LookupGroup":   user.LookupGroup,
	"LookupGroupId": user.LookupGroupId,
	"User":          qlang.StructOf((*user.User)(nil)),
	"Current":       user.Current,
	"Lookup":        user.Lookup,
	"LookupId":      user.LookupId,
}
