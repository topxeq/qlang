package excelize

import (
	"github.com/xuri/excelize/v2"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/xuri/excelize",

	"Area":                               excelize.Area,
	"Area3D":                             excelize.Area3D,
	"Area3DPercentStacked":               excelize.Area3DPercentStacked,
	"Area3DStacked":                      excelize.Area3DStacked,
	"AreaPercentStacked":                 excelize.AreaPercentStacked,
	"AreaStacked":                        excelize.AreaStacked,
	"ArgEmpty":                           excelize.ArgEmpty,
	"ArgError":                           excelize.ArgError,
	"ArgList":                            excelize.ArgList,
	"ArgMatrix":                          excelize.ArgMatrix,
	"ArgNumber":                          excelize.ArgNumber,
	"ArgString":                          excelize.ArgString,
	"ArgUnknown":                         excelize.ArgUnknown,
	"Bar":                                excelize.Bar,
	"Bar3DClustered":                     excelize.Bar3DClustered,
	"Bar3DConeClustered":                 excelize.Bar3DConeClustered,
	"Bar3DConePercentStacked":            excelize.Bar3DConePercentStacked,
	"Bar3DConeStacked":                   excelize.Bar3DConeStacked,
	"Bar3DCylinderClustered":             excelize.Bar3DCylinderClustered,
	"Bar3DCylinderPercentStacked":        excelize.Bar3DCylinderPercentStacked,
	"Bar3DCylinderStacked":               excelize.Bar3DCylinderStacked,
	"Bar3DPercentStacked":                excelize.Bar3DPercentStacked,
	"Bar3DPyramidClustered":              excelize.Bar3DPyramidClustered,
	"Bar3DPyramidPercentStacked":         excelize.Bar3DPyramidPercentStacked,
	"Bar3DPyramidStacked":                excelize.Bar3DPyramidStacked,
	"Bar3DStacked":                       excelize.Bar3DStacked,
	"BarOfPieChart":                      excelize.BarOfPieChart,
	"BarPercentStacked":                  excelize.BarPercentStacked,
	"BarStacked":                         excelize.BarStacked,
	"Bubble":                             excelize.Bubble,
	"Bubble3D":                           excelize.Bubble3D,
	"Col":                                excelize.Col,
	"Col3D":                              excelize.Col3D,
	"Col3DClustered":                     excelize.Col3DClustered,
	"Col3DCone":                          excelize.Col3DCone,
	"Col3DConeClustered":                 excelize.Col3DConeClustered,
	"Col3DConePercentStacked":            excelize.Col3DConePercentStacked,
	"Col3DConeStacked":                   excelize.Col3DConeStacked,
	"Col3DCylinder":                      excelize.Col3DCylinder,
	"Col3DCylinderClustered":             excelize.Col3DCylinderClustered,
	"Col3DCylinderPercentStacked":        excelize.Col3DCylinderPercentStacked,
	"Col3DCylinderStacked":               excelize.Col3DCylinderStacked,
	"Col3DPercentStacked":                excelize.Col3DPercentStacked,
	"Col3DPyramid":                       excelize.Col3DPyramid,
	"Col3DPyramidClustered":              excelize.Col3DPyramidClustered,
	"Col3DPyramidPercentStacked":         excelize.Col3DPyramidPercentStacked,
	"Col3DPyramidStacked":                excelize.Col3DPyramidStacked,
	"Col3DStacked":                       excelize.Col3DStacked,
	"ColPercentStacked":                  excelize.ColPercentStacked,
	"ColStacked":                         excelize.ColStacked,
	"ContentTypeDrawing":                 excelize.ContentTypeDrawing,
	"ContentTypeDrawingML":               excelize.ContentTypeDrawingML,
	"ContentTypeMacro":                   excelize.ContentTypeMacro,
	"ContentTypeSpreadSheetMLChartsheet": excelize.ContentTypeSpreadSheetMLChartsheet,
	"ContentTypeSpreadSheetMLComments":   excelize.ContentTypeSpreadSheetMLComments,
	"ContentTypeSpreadSheetMLPivotCacheDefinition": excelize.ContentTypeSpreadSheetMLPivotCacheDefinition,
	"ContentTypeSpreadSheetMLPivotTable":           excelize.ContentTypeSpreadSheetMLPivotTable,
	"ContentTypeSpreadSheetMLSharedStrings":        excelize.ContentTypeSpreadSheetMLSharedStrings,
	"ContentTypeSpreadSheetMLTable":                excelize.ContentTypeSpreadSheetMLTable,
	"ContentTypeSpreadSheetMLWorksheet":            excelize.ContentTypeSpreadSheetMLWorksheet,
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
	"DataValidationTypeTextLeng":                   excelize.DataValidationTypeTextLeng,
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
	"MaxColumnWidth":                               excelize.MaxColumnWidth,
	"MaxFileNameLength":                            excelize.MaxFileNameLength,
	"MaxFontFamilyLength":                          excelize.MaxFontFamilyLength,
	"MaxFontSize":                                  excelize.MaxFontSize,
	"MaxRowHeight":                                 excelize.MaxRowHeight,
	"NameSpaceDublinCore":                          excelize.NameSpaceDublinCore,
	"NameSpaceDublinCoreMetadataIntiative":         excelize.NameSpaceDublinCoreMetadataIntiative,
	"NameSpaceDublinCoreTerms":                     excelize.NameSpaceDublinCoreTerms,
	"NameSpaceXML":                                 excelize.NameSpaceXML,
	"NameSpaceXMLSchemaInstance":                   excelize.NameSpaceXMLSchemaInstance,
	"OrientationLandscape":                         excelize.OrientationLandscape,
	"OrientationPortrait":                          excelize.OrientationPortrait,
	"Pie":                                          excelize.Pie,
	"Pie3D":                                        excelize.Pie3D,
	"PieOfPieChart":                                excelize.PieOfPieChart,
	"Radar":                                        excelize.Radar,
	"STCellFormulaTypeArray":                       excelize.STCellFormulaTypeArray,
	"STCellFormulaTypeDataTable":                   excelize.STCellFormulaTypeDataTable,
	"STCellFormulaTypeNormal":                      excelize.STCellFormulaTypeNormal,
	"STCellFormulaTypeShared":                      excelize.STCellFormulaTypeShared,
	"Scatter":                                      excelize.Scatter,
	"SourceRelationshipChart":                      excelize.SourceRelationshipChart,
	"SourceRelationshipChartsheet":                 excelize.SourceRelationshipChartsheet,
	"SourceRelationshipComments":                   excelize.SourceRelationshipComments,
	"SourceRelationshipDialogsheet":                excelize.SourceRelationshipDialogsheet,
	"SourceRelationshipDrawingML":                  excelize.SourceRelationshipDrawingML,
	"SourceRelationshipDrawingVML":                 excelize.SourceRelationshipDrawingVML,
	"SourceRelationshipHyperLink":                  excelize.SourceRelationshipHyperLink,
	"SourceRelationshipImage":                      excelize.SourceRelationshipImage,
	"SourceRelationshipOfficeDocument":             excelize.SourceRelationshipOfficeDocument,
	"SourceRelationshipPivotCache":                 excelize.SourceRelationshipPivotCache,
	"SourceRelationshipPivotTable":                 excelize.SourceRelationshipPivotTable,
	"SourceRelationshipSharedStrings":              excelize.SourceRelationshipSharedStrings,
	"SourceRelationshipTable":                      excelize.SourceRelationshipTable,
	"SourceRelationshipVBAProject":                 excelize.SourceRelationshipVBAProject,
	"SourceRelationshipWorkSheet":                  excelize.SourceRelationshipWorkSheet,
	"StreamChunkSize":                              excelize.StreamChunkSize,
	"StrictNameSpaceSpreadSheet":                   excelize.StrictNameSpaceSpreadSheet,
	"StrictSourceRelationship":                     excelize.StrictSourceRelationship,
	"StrictSourceRelationshipChart":                excelize.StrictSourceRelationshipChart,
	"StrictSourceRelationshipComments":             excelize.StrictSourceRelationshipComments,
	"StrictSourceRelationshipImage":                excelize.StrictSourceRelationshipImage,
	"StrictSourceRelationshipOfficeDocument":       excelize.StrictSourceRelationshipOfficeDocument,
	"Surface3D":                                    excelize.Surface3D,
	"TotalCellChars":                               excelize.TotalCellChars,
	"TotalColumns":                                 excelize.TotalColumns,
	"TotalRows":                                    excelize.TotalRows,
	"TotalSheetHyperlinks":                         excelize.TotalSheetHyperlinks,
	"WireframeContour":                             excelize.WireframeContour,
	"WireframeSurface3D":                           excelize.WireframeSurface3D,
	"XMLHeader":                                    excelize.XMLHeader,

	"ErrAddVBAProject":                  excelize.ErrAddVBAProject,
	"ErrCellCharsLength":                excelize.ErrCellCharsLength,
	"ErrColumnNumber":                   excelize.ErrColumnNumber,
	"ErrColumnWidth":                    excelize.ErrColumnWidth,
	"ErrCoordinates":                    excelize.ErrCoordinates,
	"ErrDataValidationFormulaLenth":     excelize.ErrDataValidationFormulaLenth,
	"ErrDataValidationRange":            excelize.ErrDataValidationRange,
	"ErrDefinedNameScope":               excelize.ErrDefinedNameScope,
	"ErrDefinedNameduplicate":           excelize.ErrDefinedNameduplicate,
	"ErrEncrypt":                        excelize.ErrEncrypt,
	"ErrExistsWorksheet":                excelize.ErrExistsWorksheet,
	"ErrFontLength":                     excelize.ErrFontLength,
	"ErrFontSize":                       excelize.ErrFontSize,
	"ErrGroupSheets":                    excelize.ErrGroupSheets,
	"ErrImgExt":                         excelize.ErrImgExt,
	"ErrInvalidFormula":                 excelize.ErrInvalidFormula,
	"ErrMaxFileNameLength":              excelize.ErrMaxFileNameLength,
	"ErrMaxRowHeight":                   excelize.ErrMaxRowHeight,
	"ErrOutlineLevel":                   excelize.ErrOutlineLevel,
	"ErrParameterInvalid":               excelize.ErrParameterInvalid,
	"ErrParameterRequired":              excelize.ErrParameterRequired,
	"ErrSheetIdx":                       excelize.ErrSheetIdx,
	"ErrStreamSetColWidth":              excelize.ErrStreamSetColWidth,
	"ErrToExcelTime":                    excelize.ErrToExcelTime,
	"ErrTotalSheetHyperlinks":           excelize.ErrTotalSheetHyperlinks,
	"ErrUnknownEncryptMechanism":        excelize.ErrUnknownEncryptMechanism,
	"ErrUnsupportEncryptMechanism":      excelize.ErrUnsupportEncryptMechanism,
	"HSLModel":                          excelize.HSLModel,
	"NameSpaceDrawingML":                excelize.NameSpaceDrawingML,
	"NameSpaceDrawingMLChart":           excelize.NameSpaceDrawingMLChart,
	"NameSpaceDrawingMLSpreadSheet":     excelize.NameSpaceDrawingMLSpreadSheet,
	"NameSpaceMacExcel2008Main":         excelize.NameSpaceMacExcel2008Main,
	"NameSpaceSpreadSheet":              excelize.NameSpaceSpreadSheet,
	"NameSpaceSpreadSheetExcel2006Main": excelize.NameSpaceSpreadSheetExcel2006Main,
	"NameSpaceSpreadSheetX14":           excelize.NameSpaceSpreadSheetX14,
	"NameSpaceSpreadSheetX15":           excelize.NameSpaceSpreadSheetX15,
	"SourceRelationship":                excelize.SourceRelationship,
	"SourceRelationshipChart20070802":   excelize.SourceRelationshipChart20070802,
	"SourceRelationshipChart2014":       excelize.SourceRelationshipChart2014,
	"SourceRelationshipChart201506":     excelize.SourceRelationshipChart201506,
	"SourceRelationshipCompatibility":   excelize.SourceRelationshipCompatibility,
	"XMLHeaderByte":                     excelize.XMLHeaderByte,

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
	"ReadZipReader":         excelize.ReadZipReader,
	"SplitCellName":         excelize.SplitCellName,
	"ThemeColor":            excelize.ThemeColor,

	"Alignment":                  qlang.StructOf((*excelize.Alignment)(nil)),
	"Border":                     qlang.StructOf((*excelize.Border)(nil)),
	"Cell":                       qlang.StructOf((*excelize.Cell)(nil)),
	"Comment":                    qlang.StructOf((*excelize.Comment)(nil)),
	"DataIntegrity":              qlang.StructOf((*excelize.DataIntegrity)(nil)),
	"DataValidation":             qlang.StructOf((*excelize.DataValidation)(nil)),
	"datavalidation":             excelize.NewDataValidation,
	"NewDataValidation":          excelize.NewDataValidation,
	"DefinedName":                qlang.StructOf((*excelize.DefinedName)(nil)),
	"DocProperties":              qlang.StructOf((*excelize.DocProperties)(nil)),
	"EncryptedKey":               qlang.StructOf((*excelize.EncryptedKey)(nil)),
	"Encryption":                 qlang.StructOf((*excelize.Encryption)(nil)),
	"ErrSheetNotExist":           qlang.StructOf((*excelize.ErrSheetNotExist)(nil)),
	"file":                       excelize.NewFile,
	"NewFile":                    excelize.NewFile,
	"OpenFile":                   excelize.OpenFile,
	"OpenReader":                 excelize.OpenReader,
	"Fill":                       qlang.StructOf((*excelize.Fill)(nil)),
	"Font":                       qlang.StructOf((*excelize.Font)(nil)),
	"FormatHeaderFooter":         qlang.StructOf((*excelize.FormatHeaderFooter)(nil)),
	"FormatPageMargins":          qlang.StructOf((*excelize.FormatPageMargins)(nil)),
	"FormatSheetProtection":      qlang.StructOf((*excelize.FormatSheetProtection)(nil)),
	"FormulaOpts":                qlang.StructOf((*excelize.FormulaOpts)(nil)),
	"HSL":                        qlang.StructOf((*excelize.HSL)(nil)),
	"HyperlinkOpts":              qlang.StructOf((*excelize.HyperlinkOpts)(nil)),
	"KeyData":                    qlang.StructOf((*excelize.KeyData)(nil)),
	"KeyEncryptor":               qlang.StructOf((*excelize.KeyEncryptor)(nil)),
	"KeyEncryptors":              qlang.StructOf((*excelize.KeyEncryptors)(nil)),
	"Options":                    qlang.StructOf((*excelize.Options)(nil)),
	"PivotTableField":            qlang.StructOf((*excelize.PivotTableField)(nil)),
	"Protection":                 qlang.StructOf((*excelize.Protection)(nil)),
	"RichTextRun":                qlang.StructOf((*excelize.RichTextRun)(nil)),
	"RowOpts":                    qlang.StructOf((*excelize.RowOpts)(nil)),
	"SparklineOption":            qlang.StructOf((*excelize.SparklineOption)(nil)),
	"stack":                      excelize.NewStack,
	"NewStack":                   excelize.NewStack,
	"StandardEncryptionHeader":   qlang.StructOf((*excelize.StandardEncryptionHeader)(nil)),
	"StandardEncryptionVerifier": qlang.StructOf((*excelize.StandardEncryptionVerifier)(nil)),
	"Style":                      qlang.StructOf((*excelize.Style)(nil)),
}