package docxrepl

import (
	"github.com/topxeq/docxrepl"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/docxrepl",

	"CloseDelimiter":  docxrepl.CloseDelimiter,
	"DocumentXml":     docxrepl.DocumentXml,
	"OpenDelimiter":   docxrepl.OpenDelimiter,
	"RunElementName":  docxrepl.RunElementName,
	"TextElementName": docxrepl.TextElementName,

	"CloseDelimiterRegex":    docxrepl.CloseDelimiterRegex,
	"ErrPlaceholderNotFound": docxrepl.ErrPlaceholderNotFound,
	"ErrTagsInvalid":         docxrepl.ErrTagsInvalid,
	"FooterPathRegex":        docxrepl.FooterPathRegex,
	"HeaderPathRegex":        docxrepl.HeaderPathRegex,
	"OpenDelimiterRegex":     docxrepl.OpenDelimiterRegex,
	"RunCloseTagRegex":       docxrepl.RunCloseTagRegex,
	"RunOpenTagRegex":        docxrepl.RunOpenTagRegex,
	"RunSingletonTagRegex":   docxrepl.RunSingletonTagRegex,
	"TextCloseTagRegex":      docxrepl.TextCloseTagRegex,
	"TextOpenTagRegex":       docxrepl.TextOpenTagRegex,

	"AddPlaceholderDelimiter":    docxrepl.AddPlaceholderDelimiter,
	"IsDelimitedPlaceholder":     docxrepl.IsDelimitedPlaceholder,
	"NewFragmentID":              docxrepl.NewFragmentID,
	"NewRunID":                   docxrepl.NewRunID,
	"RemovePlaceholderDelimiter": docxrepl.RemovePlaceholderDelimiter,
	"ResetFragmentIdCounter":     docxrepl.ResetFragmentIdCounter,
	"ResetRunIdCounter":          docxrepl.ResetRunIdCounter,
	"ValidatePositions":          docxrepl.ValidatePositions,

	"Open":                   docxrepl.Open,
	"OpenBytes":              docxrepl.OpenBytes,
	"Placeholder":            qlang.StructOf((*docxrepl.Placeholder)(nil)),
	"ParsePlaceholders":      docxrepl.ParsePlaceholders,
	"PlaceholderFragment":    qlang.StructOf((*docxrepl.PlaceholderFragment)(nil)),
	"placeholderfragment":    docxrepl.NewPlaceholderFragment,
	"NewPlaceholderFragment": docxrepl.NewPlaceholderFragment,
	"Position":               qlang.StructOf((*docxrepl.Position)(nil)),
	"reader":                 docxrepl.NewReader,
	"NewReader":              docxrepl.NewReader,
	"replacer":               docxrepl.NewReplacer,
	"NewReplacer":            docxrepl.NewReplacer,
	"Run":                    qlang.StructOf((*docxrepl.Run)(nil)),
	"NewEmptyRun":            docxrepl.NewEmptyRun,
	"runparser":              docxrepl.NewRunParser,
	"NewRunParser":           docxrepl.NewRunParser,
	"TagPair":                qlang.StructOf((*docxrepl.TagPair)(nil)),

	"ConvertToReplaceMap": docxrepl.ConvertToReplaceMap,
	"ReplaceInWordBytes":  docxrepl.ReplaceInWordBytes,
}
