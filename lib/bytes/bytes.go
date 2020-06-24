package bytes

import (
	"bytes"
	"fmt"
	"reflect"

	qlang "github.com/topxeq/qlang/spec"
)

// -----------------------------------------------------------------------------

func newBuffer(a ...interface{}) *bytes.Buffer {

	if len(a) == 0 {
		return new(bytes.Buffer)
	}
	switch v := a[0].(type) {
	case []byte:
		return bytes.NewBuffer(v)
	case string:
		return bytes.NewBufferString(v)
	}
	panic("bytes.buffer() - unsupported argument type")
}

func from(v interface{}) []byte {

	switch args := v.(type) {
	case []int:
		b := make([]byte, len(args))
		for i, v := range args {
			b[i] = byte(v)
		}
		return b
	case string:
		return []byte(args)
	default:
		if v == nil {
			return nil
		}
		panic(fmt.Sprintf("can't convert from `%v` to []byte", reflect.TypeOf(v)))
	}
}

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name":           "bytes",
	"buffer":          newBuffer,
	"NewBuffer":       newBuffer,
	"NewBufferString": bytes.NewBufferString,
	"Buffer":          qlang.StructOf((*bytes.Buffer)(nil)),
	"From":            from,
	"Equal":           bytes.Equal,
	"NewReader":       bytes.NewReader,
	"Contains":        bytes.Contains,
	"Index":           bytes.Index,
	"IndexAny":        bytes.IndexAny,
	"Join":            bytes.Join,
	"Title":           bytes.Title,
	"ToLower":         bytes.ToLower,
	"ToTitle":         bytes.ToTitle,
	"ToUpper":         bytes.ToUpper,
	"Trim":            bytes.Trim,
}

// -----------------------------------------------------------------------------
