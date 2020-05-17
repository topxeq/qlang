package url

import (
	"net/url"

	qlang "github.com/topxeq/qlang/spec"
)

func newValues() *url.Values {
	return &url.Values{}
}

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/url",

	"PathEscape":    url.PathEscape,
	"PathUnescape":  url.PathUnescape,
	"QueryEscape":   url.QueryEscape,
	"QueryUnescape": url.QueryUnescape,

	"URL":             qlang.StructOf((*url.URL)(nil)),
	"Values":          qlang.StructOf((*url.Values)(nil)),
	"NewValues":       newValues,
	"Parse":           url.Parse,
	"ParseRequestURI": url.ParseRequestURI,
	"User":            url.User,
	"UserPassword":    url.UserPassword,
}
