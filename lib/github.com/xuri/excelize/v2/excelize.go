package excelize

import (
	"github.com/xuri/excelize/v2"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
var Exports = map[string]interface{}{
	"_name": "github.com/xuri/excelize/v2",

	"Area":                        excelize.Area,
	"Area3D":                      excelize.Area3D,
	"Area3DPercentStacked":        excelize.Area3DPercentStacked,
	"Area3DStacked":               excelize.Area3DStacked,
	"AreaPercentStacked":          excelize.AreaPercentStacked,
	"AreaStacked":                 excelize.AreaStacked,
	"ArgEmpty":                    excelize.ArgEmpty,
	"ArgError":                    excelize.ArgError,
	"ArgList":                     excelize.ArgList,
	"ArgMatrix":                   excelize.ArgMatrix,
	"ArgNumber":                   excelize.ArgNumber,
	"ArgString":                   excelize.ArgString,
	"ArgUnknown":                  excelize.ArgUnknown,
	"Bar":                         excelize.Bar,
	"Bar3DClustered":              excelize.Bar3DClustered,
	"Bar3DConeClustered":          excelize.Bar3DConeClustered,
	"Bar3DConePercentStacked":     excelize.Bar3DConePercentStacked,
	"Bar3DConeStacked":            excelize.Bar3DConeStacked,
	"Bar3DCylinderClustered":      excelize.Bar3DCylinderClustered,
	"Bar3DCylinderPercentStacked": excelize.Bar3DCylinderPercentStacked,
	"Bar3DCylinderStacked":        excelize.Bar3DCylinderStacked,
	"Bar3DPercentStacked":         excelize.Bar3DPercentStacked,
	"Bar3DPyramidClustered":       excelize.Bar3DPyramidClustered,
	"Bar3DPyramidPercentStacked":  excelize.Bar3DPyramidPercentStacked,
	"Bar3DPyramidStacked":         excelize.Bar3DPyramidStacked,
	"Bar3DStacked":                excelize.Bar3DStacked,
	"BarOfPieChart":               excelize.BarOfPieChart,
	"BarPercentStacked":           excelize.BarPercentStacked,
	"BarStacked":                  excelize.BarStacked,
	"Bubble":                      excelize.Bubble,
	"Bubble3D":                    excelize.Bubble3D,
	"CellTypeBool":                excelize.CellTypeBool,
	"CellTypeDate":                excelize.CellTypeDate,
	"CellTypeError":               excelize.CellTypeError,
	"CellTypeNumber":              excelize.CellTypeNumber,
	// "CellTypeString":                     excelize.CellTypeString,
	"CellTypeUnset":                                excelize.CellTypeUnset,
	"Col":                                          excelize.Col,
	"Col3D":                                        excelize.Col3D,
	"Col3DClustered":                               excelize.Col3DClustered,
	"Col3DCone":                                    excelize.Col3DCone,
	"Col3DConeClustered":                           excelize.Col3DConeClustered,
	"Col3DConePercentStacked":                      excelize.Col3DConePercentStacked,
	"Col3DConeStacked":                             excelize.Col3DConeStacked,
	"Col3DCylinder":                                excelize.Col3DCylinder,
	"Col3DCylinderClustered":                       excelize.Col3DCylinderClustered,
	"Col3DCylinderPercentStacked":                  excelize.Col3DCylinderPercentStacked,
	"Col3DCylinderStacked":                         excelize.Col3DCylinderStacked,
	"Col3DPercentStacked":                          excelize.Col3DPercentStacked,
	"Col3DPyramid":                                 excelize.Col3DPyramid,
	"Col3DPyramidClustered":                        excelize.Col3DPyramidClustered,
	"Col3DPyramidPercentStacked":                   excelize.Col3DPyramidPercentStacked,
	"Col3DPyramidStacked":                          excelize.Col3DPyramidStacked,
	"Col3DStacked":                                 excelize.Col3DStacked,
	"ColPercentStacked":                            excelize.ColPercentStacked,
	"ColStacked":                                   excelize.ColStacked,
	"ColorMappingTypeAccent1":                      excelize.ColorMappingTypeAccent1,
	"ColorMappingTypeAccent2":                      excelize.ColorMappingTypeAccent2,
	"ColorMappingTypeAccent3":                      excelize.ColorMappingTypeAccent3,
	"ColorMappingTypeAccent4":                      excelize.ColorMappingTypeAccent4,
	"ColorMappingTypeAccent5":                      excelize.ColorMappingTypeAccent5,
	"ColorMappingTypeAccent6":                      excelize.ColorMappingTypeAccent6,
	"ColorMappingTypeDark1":                        excelize.ColorMappingTypeDark1,
	"ColorMappingTypeDark2":                        excelize.ColorMappingTypeDark2,
	"ColorMappingTypeFollowedHyperlink":            excelize.ColorMappingTypeFollowedHyperlink,
	"ColorMappingTypeHyperlink":                    excelize.ColorMappingTypeHyperlink,
	"ColorMappingTypeLight1":                       excelize.ColorMappingTypeLight1,
	"ColorMappingTypeLight2":                       excelize.ColorMappingTypeLight2,
	"ColorMappingTypeUnset":                        excelize.ColorMappingTypeUnset,
	"ContentTypeAddinMacro":                        excelize.ContentTypeAddinMacro,
	"ContentTypeDrawing":                           excelize.ContentTypeDrawing,
	"ContentTypeDrawingML":                         excelize.ContentTypeDrawingML,
	"ContentTypeMacro":                             excelize.ContentTypeMacro,
	"ContentTypeSheetML":                           excelize.ContentTypeSheetML,
	"ContentTypeSpreadSheetMLChartsheet":           excelize.ContentTypeSpreadSheetMLChartsheet,
	"ContentTypeSpreadSheetMLComments":             excelize.ContentTypeSpreadSheetMLComments,
	"ContentTypeSpreadSheetMLPivotCacheDefinition": excelize.ContentTypeSpreadSheetMLPivotCacheDefinition,
	"ContentTypeSpreadSheetMLPivotTable":           excelize.ContentTypeSpreadSheetMLPivotTable,
	"ContentTypeSpreadSheetMLSharedStrings":        excelize.ContentTypeSpreadSheetMLSharedStrings,
	"ContentTypeSpreadSheetMLTable":                excelize.ContentTypeSpreadSheetMLTable,
	"ContentTypeSpreadSheetMLWorksheet":            excelize.ContentTypeSpreadSheetMLWorksheet,
	"ContentTypeTemplate":                          excelize.ContentTypeTemplate,
	"ContentTypeTemplateMacro":                     excelize.ContentTypeTemplateMacro,
	"ContentTypeVBA":                               excelize.ContentTypeVBA,
	"ContentTypeVML":                               excelize.ContentTypeVML,
	"Contour":                                      excelize.Contour,
	"DataValidationErrorStyleInformation":          excelize.DataValidationErrorStyleInformation,
	"DataValidationErrorStyleStop":                 excelize.DataValidationErrorStyleStop,
	"DataValidationErrorStyleWarning":              excelize.DataValidationErrorStyleWarning,
	"DataValidationOperatorBetween":                excelize.DataValidationOperatorBetween,
	"DataValidationOperatorEqual":                  excelize.DataValidationOperatorEqual,
	"DataValidationOperatorGreaterThan":            excelize.DataValidationOperatorGreaterThan,
	"DataValidationOperatorGreaterThanOrEqual":     excelize.DataValidationOperatorGreaterThanOrEqual,
	"DataValidationOperatorLessThan":               excelize.DataValidationOperatorLessThan,
	"DataValidationOperatorLessThanOrEqual":        excelize.DataValidationOperatorLessThanOrEqual,
	"DataValidationOperatorNotBetween":             excelize.DataValidationOperatorNotBetween,
	"DataValidationOperatorNotEqual":               excelize.DataValidationOperatorNotEqual,
	"DataValidationTypeCustom":                     excelize.DataValidationTypeCustom,
	"DataValidationTypeDate":                       excelize.DataValidationTypeDate,
	"DataValidationTypeDecimal":                    excelize.DataValidationTypeDecimal,
	"DataValidationTypeTextLength":                 excelize.DataValidationTypeTextLength,
	"DataValidationTypeTime":                       excelize.DataValidationTypeTime,
	"DataValidationTypeWhole":                      excelize.DataValidationTypeWhole,
	"Doughnut":                                     excelize.Doughnut,
	"EMU":                                          excelize.EMU,
	"ExtURIConditionalFormattings":                 excelize.ExtURIConditionalFormattings,
	"ExtURIDataValidations":                        excelize.ExtURIDataValidations,
	"ExtURIDrawingBlip":                            excelize.ExtURIDrawingBlip,
	"ExtURIIgnoredErrors":                          excelize.ExtURIIgnoredErrors,
	"ExtURIMacExcelMX":                             excelize.ExtURIMacExcelMX,
	"ExtURIProtectedRanges":                        excelize.ExtURIProtectedRanges,
	"ExtURISlicerCachesListX14":                    excelize.ExtURISlicerCachesListX14,
	"ExtURISlicerListX14":                          excelize.ExtURISlicerListX14,
	"ExtURISlicerListX15":                          excelize.ExtURISlicerListX15,
	"ExtURISparklineGroups":                        excelize.ExtURISparklineGroups,
	"ExtURITimelineRefs":                           excelize.ExtURITimelineRefs,
	"ExtURIWebExtensions":                          excelize.ExtURIWebExtensions,
	"Line":                                         excelize.Line,
	"MaxCellStyles":                                excelize.MaxCellStyles,
	"MaxColumnWidth":                               excelize.MaxColumnWidth,
	"MaxColumns":                                   excelize.MaxColumns,
	"MaxFieldLength":                               excelize.MaxFieldLength,
	"MaxFilePathLength":                            excelize.MaxFilePathLength,
	"MaxFontFamilyLength":                          excelize.MaxFontFamilyLength,
	"MaxFontSize":                                  excelize.MaxFontSize,
	"MaxRowHeight":                                 excelize.MaxRowHeight,
	"MinColumns":                                   excelize.MinColumns,
	"MinFontSize":                                  excelize.MinFontSize,
	"NameSpaceDublinCore":                          excelize.NameSpaceDublinCore,
	"NameSpaceDublinCoreMetadataInitiative":        excelize.NameSpaceDublinCoreMetadataInitiative,
	"NameSpaceDublinCoreTerms":                     excelize.NameSpaceDublinCoreTerms,
	"NameSpaceXML":                                 excelize.NameSpaceXML,
	"NameSpaceXMLSchemaInstance":                   excelize.NameSpaceXMLSchemaInstance,
	// "OrientationLandscape":                         excelize.OrientationLandscape,
	// "OrientationPortrait":                          excelize.OrientationPortrait,
	"Pie":                                    excelize.Pie,
	"Pie3D":                                  excelize.Pie3D,
	"PieOfPieChart":                          excelize.PieOfPieChart,
	"Radar":                                  excelize.Radar,
	"STCellFormulaTypeArray":                 excelize.STCellFormulaTypeArray,
	"STCellFormulaTypeDataTable":             excelize.STCellFormulaTypeDataTable,
	"STCellFormulaTypeNormal":                excelize.STCellFormulaTypeNormal,
	"STCellFormulaTypeShared":                excelize.STCellFormulaTypeShared,
	"Scatter":                                excelize.Scatter,
	"SourceRelationshipChart":                excelize.SourceRelationshipChart,
	"SourceRelationshipChartsheet":           excelize.SourceRelationshipChartsheet,
	"SourceRelationshipComments":             excelize.SourceRelationshipComments,
	"SourceRelationshipDialogsheet":          excelize.SourceRelationshipDialogsheet,
	"SourceRelationshipDrawingML":            excelize.SourceRelationshipDrawingML,
	"SourceRelationshipDrawingVML":           excelize.SourceRelationshipDrawingVML,
	"SourceRelationshipHyperLink":            excelize.SourceRelationshipHyperLink,
	"SourceRelationshipImage":                excelize.SourceRelationshipImage,
	"SourceRelationshipOfficeDocument":       excelize.SourceRelationshipOfficeDocument,
	"SourceRelationshipPivotCache":           excelize.SourceRelationshipPivotCache,
	"SourceRelationshipPivotTable":           excelize.SourceRelationshipPivotTable,
	"SourceRelationshipSharedStrings":        excelize.SourceRelationshipSharedStrings,
	"SourceRelationshipTable":                excelize.SourceRelationshipTable,
	"SourceRelationshipVBAProject":           excelize.SourceRelationshipVBAProject,
	"SourceRelationshipWorkSheet":            excelize.SourceRelationshipWorkSheet,
	"StreamChunkSize":                        excelize.StreamChunkSize,
	"StrictNameSpaceSpreadSheet":             excelize.StrictNameSpaceSpreadSheet,
	"StrictSourceRelationship":               excelize.StrictSourceRelationship,
	"StrictSourceRelationshipChart":          excelize.StrictSourceRelationshipChart,
	"StrictSourceRelationshipComments":       excelize.StrictSourceRelationshipComments,
	"StrictSourceRelationshipImage":          excelize.StrictSourceRelationshipImage,
	"StrictSourceRelationshipOfficeDocument": excelize.StrictSourceRelationshipOfficeDocument,
	"Surface3D":                              excelize.Surface3D,
	"TotalCellChars":                         excelize.TotalCellChars,
	"TotalRows":                              excelize.TotalRows,
	"TotalSheetHyperlinks":                   excelize.TotalSheetHyperlinks,
	"UnzipSizeLimit":                         excelize.UnzipSizeLimit,
	"WireframeContour":                       excelize.WireframeContour,
	"WireframeSurface3D":                     excelize.WireframeSurface3D,

	"ErrAddVBAProject":               excelize.ErrAddVBAProject,
	"ErrAttrValBool":                 excelize.ErrAttrValBool,
	"ErrCellCharsLength":             excelize.ErrCellCharsLength,
	"ErrColumnNumber":                excelize.ErrColumnNumber,
	"ErrColumnWidth":                 excelize.ErrColumnWidth,
	"ErrCoordinates":                 excelize.ErrCoordinates,
	"ErrCustomNumFmt":                excelize.ErrCustomNumFmt,
	"ErrDataValidationFormulaLength": excelize.ErrDataValidationFormulaLength,
	"ErrDataValidationRange":         excelize.ErrDataValidationRange,
	"ErrDefinedNameDuplicate":        excelize.ErrDefinedNameDuplicate,
	"ErrDefinedNameScope":            excelize.ErrDefinedNameScope,
	// "ErrExistsWorksheet":                      excelize.ErrExistsWorksheet,
	"ErrFontLength":                           excelize.ErrFontLength,
	"ErrFontSize":                             excelize.ErrFontSize,
	"ErrGroupSheets":                          excelize.ErrGroupSheets,
	"ErrImgExt":                               excelize.ErrImgExt,
	"ErrInvalidFormula":                       excelize.ErrInvalidFormula,
	"ErrMaxFilePathLength":                    excelize.ErrMaxFilePathLength,
	"ErrMaxRowHeight":                         excelize.ErrMaxRowHeight,
	"ErrMaxRows":                              excelize.ErrMaxRows,
	"ErrOptionsUnzipSizeLimit":                excelize.ErrOptionsUnzipSizeLimit,
	"ErrOutlineLevel":                         excelize.ErrOutlineLevel,
	"ErrParameterInvalid":                     excelize.ErrParameterInvalid,
	"ErrParameterRequired":                    excelize.ErrParameterRequired,
	"ErrPasswordLengthInvalid":                excelize.ErrPasswordLengthInvalid,
	"ErrSave":                                 excelize.ErrSave,
	"ErrSheetIdx":                             excelize.ErrSheetIdx,
	"ErrSparkline":                            excelize.ErrSparkline,
	"ErrSparklineLocation":                    excelize.ErrSparklineLocation,
	"ErrSparklineRange":                       excelize.ErrSparklineRange,
	"ErrSparklineStyle":                       excelize.ErrSparklineStyle,
	"ErrSparklineType":                        excelize.ErrSparklineType,
	"ErrStreamSetColWidth":                    excelize.ErrStreamSetColWidth,
	"ErrTotalSheetHyperlinks":                 excelize.ErrTotalSheetHyperlinks,
	"ErrUnknownEncryptMechanism":              excelize.ErrUnknownEncryptMechanism,
	"ErrUnprotectSheet":                       excelize.ErrUnprotectSheet,
	"ErrUnprotectSheetPassword":               excelize.ErrUnprotectSheetPassword,
	"ErrUnsupportedEncryptMechanism":          excelize.ErrUnsupportedEncryptMechanism,
	"ErrUnsupportedHashAlgorithm":             excelize.ErrUnsupportedHashAlgorithm,
	"ErrUnsupportedNumberFormat":              excelize.ErrUnsupportedNumberFormat,
	"ErrWorkbookFileFormat":                   excelize.ErrWorkbookFileFormat,
	"ErrWorkbookPassword":                     excelize.ErrWorkbookPassword,
	"HSLModel":                                excelize.HSLModel,
	"NameSpaceDocumentPropertiesVariantTypes": excelize.NameSpaceDocumentPropertiesVariantTypes,
	"NameSpaceDrawingML":                      excelize.NameSpaceDrawingML,
	"NameSpaceDrawingMLChart":                 excelize.NameSpaceDrawingMLChart,
	"NameSpaceDrawingMLSpreadSheet":           excelize.NameSpaceDrawingMLSpreadSheet,
	"NameSpaceMacExcel2008Main":               excelize.NameSpaceMacExcel2008Main,
	"NameSpaceSpreadSheet":                    excelize.NameSpaceSpreadSheet,
	"NameSpaceSpreadSheetExcel2006Main":       excelize.NameSpaceSpreadSheetExcel2006Main,
	"NameSpaceSpreadSheetX14":                 excelize.NameSpaceSpreadSheetX14,
	"NameSpaceSpreadSheetX15":                 excelize.NameSpaceSpreadSheetX15,
	"SourceRelationship":                      excelize.SourceRelationship,
	"SourceRelationshipChart20070802":         excelize.SourceRelationshipChart20070802,
	"SourceRelationshipChart2014":             excelize.SourceRelationshipChart2014,
	"SourceRelationshipChart201506":           excelize.SourceRelationshipChart201506,
	"SourceRelationshipCompatibility":         excelize.SourceRelationshipCompatibility,

	"CellNameToCoordinates": excelize.CellNameToCoordinates,
	"ColumnNameToNumber":    excelize.ColumnNameToNumber,
	"ColumnNumberToName":    excelize.ColumnNumberToName,
	"CoordinatesToCellName": excelize.CoordinatesToCellName,
	"Decrypt":               excelize.Decrypt,
	"Encrypt":               excelize.Encrypt,
	"ExcelDateToTime":       excelize.ExcelDateToTime,
	"HSLToRGB":              excelize.HSLToRGB,
	"JoinCellName":          excelize.JoinCellName,
	"RGBToHSL":              excelize.RGBToHSL,
	"SplitCellName":         excelize.SplitCellName,
	"ThemeColor":            excelize.ThemeColor,

	"Alignment":         qlang.StructOf((*excelize.Alignment)(nil)),
	"AppProperties":     qlang.StructOf((*excelize.AppProperties)(nil)),
	"Border":            qlang.StructOf((*excelize.Border)(nil)),
	"Cell":              qlang.StructOf((*excelize.Cell)(nil)),
	"Comment":           qlang.StructOf((*excelize.Comment)(nil)),
	"DataIntegrity":     qlang.StructOf((*excelize.DataIntegrity)(nil)),
	"DataValidation":    qlang.StructOf((*excelize.DataValidation)(nil)),
	"datavalidation":    excelize.NewDataValidation,
	"NewDataValidation": excelize.NewDataValidation,
	"DefinedName":       qlang.StructOf((*excelize.DefinedName)(nil)),
	"DocProperties":     qlang.StructOf((*excelize.DocProperties)(nil)),
	"EncryptedKey":      qlang.StructOf((*excelize.EncryptedKey)(nil)),
	"Encryption":        qlang.StructOf((*excelize.Encryption)(nil)),
	"ErrSheetNotExist":  qlang.StructOf((*excelize.ErrSheetNotExist)(nil)),
	"file":              excelize.NewFile,
	"NewFile":           excelize.NewFile,
	"OpenFile":          excelize.OpenFile,
	"OpenReader":        excelize.OpenReader,
	"Fill":              qlang.StructOf((*excelize.Fill)(nil)),
	"Font":              qlang.StructOf((*excelize.Font)(nil)),
	// "FormatHeaderFooter":         qlang.StructOf((*excelize.FormatHeaderFooter)(nil)),
	// "FormatPageMargins":          qlang.StructOf((*excelize.FormatPageMargins)(nil)),
	// "FormatSheetProtection":      qlang.StructOf((*excelize.FormatSheetProtection)(nil)),
	"FormulaOpts":     qlang.StructOf((*excelize.FormulaOpts)(nil)),
	"HSL":             qlang.StructOf((*excelize.HSL)(nil)),
	"HyperlinkOpts":   qlang.StructOf((*excelize.HyperlinkOpts)(nil)),
	"KeyData":         qlang.StructOf((*excelize.KeyData)(nil)),
	"KeyEncryptor":    qlang.StructOf((*excelize.KeyEncryptor)(nil)),
	"KeyEncryptors":   qlang.StructOf((*excelize.KeyEncryptors)(nil)),
	"Options":         qlang.StructOf((*excelize.Options)(nil)),
	"PivotTableField": qlang.StructOf((*excelize.PivotTableField)(nil)),
	"Protection":      qlang.StructOf((*excelize.Protection)(nil)),
	"RichTextRun":     qlang.StructOf((*excelize.RichTextRun)(nil)),
	"RowOpts":         qlang.StructOf((*excelize.RowOpts)(nil)),
	// "SparklineOption":            qlang.StructOf((*excelize.SparklineOption)(nil)),
	"stack":                      excelize.NewStack,
	"NewStack":                   excelize.NewStack,
	"StandardEncryptionHeader":   qlang.StructOf((*excelize.StandardEncryptionHeader)(nil)),
	"StandardEncryptionVerifier": qlang.StructOf((*excelize.StandardEncryptionVerifier)(nil)),
	"Style":                      qlang.StructOf((*excelize.Style)(nil)),
}
