package io

import (
	"io"
)

func init() {
	Exports["Discard"] = io.Discard
	Exports["NopCloser"] = io.NopCloser
	Exports["ReadAll"] = io.ReadAll
}
