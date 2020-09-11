package xmlx

import (
	"github.com/topxeq/xmlx"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/xmlx",

	"Chunk":    xmlx.Chunk,
	"ChunkAll": xmlx.ChunkAll,

	"GetXMLNode": xmlx.GetXMLNode,
}
