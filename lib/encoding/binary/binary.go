package binary

import (
	"encoding/binary"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "encoding/binary",

	"MaxVarintLen16": binary.MaxVarintLen16,
	"MaxVarintLen32": binary.MaxVarintLen32,
	"MaxVarintLen64": binary.MaxVarintLen64,

	"BigEndian":    &binary.BigEndian,
	"LittleEndian": &binary.LittleEndian,

	"PutUvarint":  binary.PutUvarint,
	"PutVarint":   binary.PutVarint,
	"Read":        binary.Read,
	"ReadUvarint": binary.ReadUvarint,
	"ReadVarint":  binary.ReadVarint,
	"Size":        binary.Size,
	"Uvarint":     binary.Uvarint,
	"Varint":      binary.Varint,
	"Write":       binary.Write,
}
