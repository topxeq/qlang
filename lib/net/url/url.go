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

	"pathEscape":    url.PathEscape,
	"pathUnescape":  url.PathUnescape,
	"queryEscape":   url.QueryEscape,
	"queryUnescape": url.QueryUnescape,

	"URL":             qlang.StructOf((*url.URL)(nil)),
	"Values":          qlang.StructOf((*url.Values)(nil)),
	"NewValues":       newValues,
	"parse":           url.Parse,
	"parseRequestURI": url.ParseRequestURI,
	"user":            url.User,
	"userPassword":    url.UserPassword,
}
