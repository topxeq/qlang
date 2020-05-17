package etree

import (
	"github.com/beevik/etree"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/beevik/etree",

	"NoIndent": etree.NoIndent,

	"ErrXML": etree.ErrXML,

	"charData":            etree.NewCharData,
	"newCData":            etree.NewCData,
	"newText":             etree.NewText,
	"comment":             etree.NewComment,
	"directive":           etree.NewDirective,
	"Document":            qlang.StructOf((*etree.Document)(nil)),
	"newDocument":         etree.NewDocument,
	"NewDocument":         etree.NewDocument,
	"newDocumentWithRoot": etree.NewDocumentWithRoot,
	"element":             etree.NewElement,
	"compilePath":         etree.CompilePath,
	"mustCompilePath":     etree.MustCompilePath,
	"procInst":            etree.NewProcInst,
	"ReadSettings":        qlang.StructOf((*etree.ReadSettings)(nil)),
	"WriteSettings":       qlang.StructOf((*etree.WriteSettings)(nil)),

	"NewCharData":         etree.NewCharData,
	"NewCData":            etree.NewCData,
	"NewText":             etree.NewText,
	"NewComment":          etree.NewComment,
	"NewDirective":        etree.NewDirective,
	"NewDocumentWithRoot": etree.NewDocumentWithRoot,
	"NewElement":          etree.NewElement,
	"CompilePath":         etree.CompilePath,
	"MustCompilePath":     etree.MustCompilePath,
	"NewProcInst":         etree.NewProcInst,
}
