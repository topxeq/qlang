package tk

import (
	"github.com/topxeq/tk"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/tk",

	"TXDEF_BUFFER_LEN": tk.TXDEF_BUFFER_LEN,

	"AdjustFloat":                         tk.AdjustFloat,
	"AESDecrypt":                          tk.AESDecrypt,
	"AESEncrypt":                          tk.AESEncrypt,
	"AddDebug":                            tk.AddDebug,
	"AddDebugF":                           tk.AddDebugF,
	"AddLastSubString":                    tk.AddLastSubString,
	"AnalyzeCommandLineParamter":          tk.AnalyzeCommandLineParamter,
	"AnalyzeURLParams":                    tk.AnalyzeURLParams,
	"AppendDualLineList":                  tk.AppendDualLineList,
	"AppendSimpleMapFromFile":             tk.AppendSimpleMapFromFile,
	"AppendStringToFile":                  tk.AppendStringToFile,
	"BitXor":                              tk.BitXor,
	"BoolToInt":                           tk.BoolToInt,
	"ByteSliceToStringDec":                tk.ByteSliceToStringDec,
	"ByteToHex":                           tk.ByteToHex,
	"BytesToHex":                          tk.BytesToHex,
	"CalCosineSimilarityBetweenFloatsBig": tk.CalCosineSimilarityBetweenFloatsBig,
	"CheckErr":                            tk.CheckErr,
	"CheckErrCompact":                     tk.CheckErrCompact,
	"CheckErrStr":                         tk.CheckErrStr,
	"CheckErrStrf":                        tk.CheckErrStrf,
	"CheckErrf":                           tk.CheckErrf,
	"CheckError":                          tk.CheckError,
	"CheckErrorFunc":                      tk.CheckErrorFunc,
	"CheckErrorString":                    tk.CheckErrorString,
	"ClearDebug":                          tk.ClearDebug,
	"ColorToHex":                          tk.ColorToHex,
	"CompareBytes":                        tk.CompareBytes,
	"CompareTimeString":                   tk.CompareTimeString,
	"Contains":                            tk.Contains,
	"ContainsIgnoreCase":                  tk.ContainsIgnoreCase,
	"ContainsIn":                          tk.ContainsIn,
	"ContainsInStringList":                tk.ContainsInStringList,
	"ConvertStringToUTF8":                 tk.ConvertStringToUTF8,
	"ConvertToGB18030":                    tk.ConvertToGB18030,
	"ConvertToGB18030Bytes":               tk.ConvertToGB18030Bytes,
	"ConvertToUTF8":                       tk.ConvertToUTF8,
	"CopyFile":                            tk.CopyFile,
	"CreateSimpleEvent":                   tk.CreateSimpleEvent,
	"CreateString":                        tk.CreateString,
	"CreateStringEmpty":                   tk.CreateStringEmpty,
	"CreateStringError":                   tk.CreateStringError,
	"CreateStringErrorF":                  tk.CreateStringErrorF,
	"CreateStringErrorFromTXError":        tk.CreateStringErrorFromTXError,
	"CreateStringSimple":                  tk.CreateStringSimple,
	"CreateStringSuccess":                 tk.CreateStringSuccess,
	"CreateStringWithObject":              tk.CreateStringWithObject,
	"CreateTXCollection":                  tk.CreateTXCollection,
	"CreateTempFile":                      tk.CreateTempFile,
	"CreateFile":                          tk.CreateFile,
	"DebugModeG":                          tk.DebugModeG,
	"DecodeFromBase64":                    tk.DecodeFromBase64,
	"DecodeHTML":                          tk.DecodeHTML,
	"DecodeStringCustom":                  tk.DecodeStringCustom,
	"DecodeStringSimple":                  tk.DecodeStringSimple,
	"DecodeStringUnderline":               tk.DecodeStringUnderline,
	"DecryptDataByTXDEE":                  tk.DecryptDataByTXDEE,
	"DecryptDataByTXDEF":                  tk.DecryptDataByTXDEF,
	"DecryptFileByTXDEF":                  tk.DecryptFileByTXDEF,
	"DecryptFileByTXDEFS":                 tk.DecryptFileByTXDEFS,
	"DecryptFileByTXDEFStream":            tk.DecryptFileByTXDEFStream,
	"DecryptFileByTXDEFStreamS":           tk.DecryptFileByTXDEFStreamS,
	"DecryptStreamByTXDEF":                tk.DecryptStreamByTXDEF,
	"DecryptStringByTXDEE":                tk.DecryptStringByTXDEE,
	"DecryptStringByTXDEF":                tk.DecryptStringByTXDEF,
	"DecryptStringByTXTE":                 tk.DecryptStringByTXTE,
	"DeepClone":                           tk.DeepClone,
	"DeepCopyFromTo":                      tk.DeepCopyFromTo,
	"DeleteItemInArray":                   tk.DeleteItemInArray,
	"DeleteItemInFloat64Array":            tk.DeleteItemInFloat64Array,
	"DeleteItemInInt64Array":              tk.DeleteItemInInt64Array,
	"DeleteItemInIntArray":                tk.DeleteItemInIntArray,
	"DeleteItemInStringArray":             tk.DeleteItemInStringArray,
	"DownloadBytes":                       tk.DownloadBytes,
	"DownloadFile":                        tk.DownloadFile,
	"DownloadPage":                        tk.DownloadPage,
	"DownloadPageByMap":                   tk.DownloadPageByMap,
	"DownloadPageUTF8":                    tk.DownloadPageUTF8,
	"DownloadWebPage":                     tk.DownloadWebPage,
	"DownloadWebPageX":                    tk.DownloadWebPageX,
	"DownloadWebBytes":                    tk.DownloadWebBytes,
	"URIToPath":                           tk.URIToPath,
	"PathToURI":                           tk.PathToURI,
	"EncodeHTML":                          tk.EncodeHTML,
	"EncodeStringCustom":                  tk.EncodeStringCustom,
	"EncodeStringCustomEx":                tk.EncodeStringCustomEx,
	"EncodeStringSimple":                  tk.EncodeStringSimple,
	"EncodeStringUnderline":               tk.EncodeStringUnderline,
	"EncodeToBase64":                      tk.EncodeToBase64,
	"ToBase64":                            tk.ToBase64,
	"EncodeToXMLString":                   tk.EncodeToXMLString,
	"EncryptDataByTXDEE":                  tk.EncryptDataByTXDEE,
	"EncryptDataByTXDEF":                  tk.EncryptDataByTXDEF,
	"EncryptFileByTXDEF":                  tk.EncryptFileByTXDEF,
	"EncryptFileByTXDEFS":                 tk.EncryptFileByTXDEFS,
	"EncryptFileByTXDEFStream":            tk.EncryptFileByTXDEFStream,
	"EncryptFileByTXDEFStreamS":           tk.EncryptFileByTXDEFStreamS,
	"EncryptFileByTXDEFWithHeader":        tk.EncryptFileByTXDEFWithHeader,
	"EncryptStreamByTXDEF":                tk.EncryptStreamByTXDEF,
	"EncryptStringByTXDEE":                tk.EncryptStringByTXDEE,
	"EncryptStringByTXDEF":                tk.EncryptStringByTXDEF,
	"EncryptStringByTXTE":                 tk.EncryptStringByTXTE,
	"EndsWith":                            tk.EndsWith,
	"EndsWithIgnoreCase":                  tk.EndsWithIgnoreCase,
	"EnsureBasePath":                      tk.EnsureBasePath,
	"EnsureMakeDirs":                      tk.EnsureMakeDirs,
	"EnsureMakeDirsE":                     tk.EnsureMakeDirsE,
	"EnsureValidFileNameX":                tk.EnsureValidFileNameX,
	"ErrStr":                              tk.ErrStr,
	"ErrStrF":                             tk.ErrStrF,
	"ErrStrToErr":                         tk.ErrStrToErr,
	"ErrStrf":                             tk.ErrStrf,
	"ErrToStr":                            tk.ErrToStr,
	"ErrToStrF":                           tk.ErrToStrF,
	"Errf":                                tk.Errf,
	"ErrorStringToError":                  tk.ErrorStringToError,
	"ErrorToString":                       tk.ErrorToString,
	"ErrorToEmptyString":                  tk.ErrorToEmptyString,
	"ErrToErrStr":                         tk.ErrToErrStr,
	"Exit":                                tk.Exit,
	"FatalErr":                            tk.FatalErr,
	"FatalErrf":                           tk.FatalErrf,
	"Fatalf":                              tk.Fatalf,
	"FatalfByCondition":                   tk.FatalfByCondition,
	"FindFirstDiffIndex":                  tk.FindFirstDiffIndex,
	"FindSamePrefix":                      tk.FindSamePrefix,
	"FlattenXML":                          tk.FlattenXML,
	"Float32ArrayToFloat64Array":          tk.Float32ArrayToFloat64Array,
	"Float64ToStr":                        tk.Float64ToStr,
	"FormToMap":                           tk.FormToMap,
	"FormatStringSliceSlice":              tk.FormatStringSliceSlice,
	"FormatTime":                          tk.FormatTime,
	"Fpl":                                 tk.Fpl,
	"Fpr":                                 tk.Fpr,
	"FromJSON":                            tk.FromJSON,
	"FromJSONWithDefault":                 tk.FromJSONWithDefault,
	"FromXML":                             tk.FromXML,
	"FromXMLX":                            tk.FromXMLX,
	"FromXMLWithDefault":                  tk.FromXMLWithDefault,
	"GenerateErrorString":                 tk.GenerateErrorString,
	"GenerateErrorStringF":                tk.GenerateErrorStringF,
	"GenerateErrorStringFTX":              tk.GenerateErrorStringFTX,
	"GenerateErrorStringTX":               tk.GenerateErrorStringTX,
	"GenerateFileListInDir":               tk.GenerateFileListInDir,
	"GenerateFileListRecursively":         tk.GenerateFileListRecursively,
	"GenerateFileListRecursivelyWithExclusive": tk.GenerateFileListRecursivelyWithExclusive,
	"GetFileList":                           tk.GetFileList,
	"GenerateJSONPResponse":                 tk.GenerateJSONPResponse,
	"GenerateJSONPResponseMix":              tk.GenerateJSONPResponseMix,
	"GenerateJSONPResponseWith2Object":      tk.GenerateJSONPResponseWith2Object,
	"GenerateJSONPResponseWith3Object":      tk.GenerateJSONPResponseWith3Object,
	"GenerateJSONPResponseWithMore":         tk.GenerateJSONPResponseWithMore,
	"GenerateJSONPResponseWithObject":       tk.GenerateJSONPResponseWithObject,
	"JSONResponseToHTML":                    tk.JSONResponseToHTML,
	"GenerateRandomFloats":                  tk.GenerateRandomFloats,
	"GenerateRandomString":                  tk.GenerateRandomString,
	"GenerateRandomStringX":                 tk.GenerateRandomStringX,
	"GetAllOSParameters":                    tk.GetAllOSParameters,
	"GetAllParameters":                      tk.GetAllParameters,
	"GetAllSwitches":                        tk.GetAllSwitches,
	"GetApplicationPath":                    tk.GetApplicationPath,
	"GetAvailableFileName":                  tk.GetAvailableFileName,
	"GetClipText":                           tk.GetClipText,
	"GetClipboardTextDefaultEmpty":          tk.GetClipboardTextDefaultEmpty,
	"GetClipboardTextWithDefault":           tk.GetClipboardTextWithDefault,
	"GetCurrentDir":                         tk.GetCurrentDir,
	"GetCurrentThreadID":                    tk.GetCurrentThreadID,
	"GetDBConnection":                       tk.GetDBConnection,
	"GetDBResultArray":                      tk.GetDBResultArray,
	"GetDBResultString":                     tk.GetDBResultString,
	"GetDBRowCount":                         tk.GetDBRowCount,
	"GetDBRowCountCompact":                  tk.GetDBRowCountCompact,
	"GetDebug":                              tk.GetDebug,
	"GetDirOfFilePath":                      tk.GetDirOfFilePath,
	"GetEnv":                                tk.GetEnv,
	"SetEnv":                                tk.SetEnv,
	"GetErrStr":                             tk.GetErrStr,
	"GetErrorString":                        tk.GetErrorString,
	"GetErrorStringSafely":                  tk.GetErrorStringSafely,
	"GetFileExt":                            tk.GetFileExt,
	"GetFilePathSeperator":                  tk.GetFilePathSeperator,
	"GetFileVar":                            tk.GetFileVar,
	"GetFormValueWithDefaultValue":          tk.GetFormValueWithDefaultValue,
	"GetGlobalEnvList":                      tk.GetGlobalEnvList,
	"GetGlobalEnvString":                    tk.GetGlobalEnvString,
	"GetInputBufferedScan":                  tk.GetInputBufferedScan,
	"GetInputPasswordf":                     tk.GetInputPasswordf,
	"GetInputf":                             tk.GetInputf,
	"GetJSONNode":                           tk.GetJSONNode,
	"GetJSONNodeAny":                        tk.GetJSONNodeAny,
	"GetJSONSubNode":                        tk.GetJSONSubNode,
	"GetJSONSubNodeAny":                     tk.GetJSONSubNodeAny,
	"GetLastComponentOfFilePath":            tk.GetLastComponentOfFilePath,
	"GetLastComponentOfUrl":                 tk.GetLastComponentOfUrl,
	"GetLinesFromFile":                      tk.GetLinesFromFile,
	"GetLoginAuth":                          tk.GetLoginAuth,
	"GetMSIStringWithDefault":               tk.GetMSIStringWithDefault,
	"GetMSSArrayFromXML":                    tk.GetMSSArrayFromXML,
	"GetMSSFromXML":                         tk.GetMSSFromXML,
	"GetNodeStringFromXML":                  tk.GetNodeStringFromXML,
	"GetNowDateString":                      tk.GetNowDateString,
	"GetNowMinutesInDay":                    tk.GetNowMinutesInDay,
	"GetNowTimeOnlyStringBeijing":           tk.GetNowTimeOnlyStringBeijing,
	"GetNowTimeString":                      tk.GetNowTimeString,
	"GetNowTimeStringFormal":                tk.GetNowTimeStringFormal,
	"GetNowTimeStringFormat":                tk.GetNowTimeStringFormat,
	"GetNowTimeStringHourMinute":            tk.GetNowTimeStringHourMinute,
	"GetOSArgs":                             tk.GetOSArgs,
	"GetOSArgsShort":                        tk.GetOSArgsShort,
	"GetOSName":                             tk.GetOSName,
	"GetParam":                              tk.GetParam,
	"GetParamI":                             tk.GetParamI,
	"GetParameter":                          tk.GetParameter,
	"GetParameterByIndexWithDefaultValue":   tk.GetParameterByIndexWithDefaultValue,
	"GetPlainAuth":                          tk.GetPlainAuth,
	"GetRandomByte":                         tk.GetRandomByte,
	"GetRandomInt64InRange":                 tk.GetRandomInt64InRange,
	"GetRandomInt64LessThan":                tk.GetRandomInt64LessThan,
	"GetRandomIntInRange":                   tk.GetRandomIntInRange,
	"GetRandomIntLessThan":                  tk.GetRandomIntLessThan,
	"GetRandomFloat":                        tk.GetRandomFloat,
	"GetRandomItem":                         tk.GetRandomItem,
	"GetRandomStringItem":                   tk.GetRandomStringItem,
	"GetRandomSubDualList":                  tk.GetRandomSubDualList,
	"GetRandomizeInt64ArrayCopy":            tk.GetRandomizeInt64ArrayCopy,
	"GetRandomizeIntArrayCopy":              tk.GetRandomizeIntArrayCopy,
	"GetRandomizeStringArrayCopy":           tk.GetRandomizeStringArrayCopy,
	"GetRandomizeSubStringArrayCopy":        tk.GetRandomizeSubStringArrayCopy,
	"GetRuntimeStack":                       tk.GetRuntimeStack,
	"GetSliceMaxLen":                        tk.GetSliceMaxLen,
	"GetStringSliceFilled":                  tk.GetStringSliceFilled,
	"GetSuccessValue":                       tk.GetSuccessValue,
	"GetSwitch":                             tk.GetSwitch,
	"GetSwitchI":                            tk.GetSwitchI,
	"GetSwitchWithDefaultInt64Value":        tk.GetSwitchWithDefaultInt64Value,
	"GetSwitchWithDefaultIntValue":          tk.GetSwitchWithDefaultIntValue,
	"GetSwitchWithDefaultFloatValue":        tk.GetSwitchWithDefaultFloatValue,
	"GetSwitchWithDefaultValue":             tk.GetSwitchWithDefaultValue,
	"GetTextFromFileOrClipboard":            tk.GetTextFromFileOrClipboard,
	"GetTimeFromUnixTimeStamp":              tk.GetTimeFromUnixTimeStamp,
	"GetTimeFromUnixTimeStampMid":           tk.GetTimeFromUnixTimeStampMid,
	"GetTimeStamp":                          tk.GetTimeStamp,
	"GetTimeStampMid":                       tk.GetTimeStampMid,
	"GetTimeStampNano":                      tk.GetTimeStampNano,
	"GetTimeStringDiffMS":                   tk.GetTimeStringDiffMS,
	"GetUUID":                               tk.GetUUID,
	"GetUUID1":                              tk.GetUUID1,
	"GetUUID4":                              tk.GetUUID4,
	"GetUserInput":                          tk.GetUserInput,
	"GetValue":                              tk.GetValue,
	"GetPointer":                            tk.GetPointer,
	"GetAddr":                               tk.GetAddr,
	"GetValueOfMSS":                         tk.GetValueOfMSS,
	"GetVar":                                tk.GetVar,
	"GetVersion":                            tk.GetVersion,
	"GetXMLNode":                            tk.GetXMLNode,
	"GlobalEnvSetG":                         tk.GlobalEnvSetG,
	"HTMLToText":                            tk.HTMLToText,
	"HasGlobalEnv":                          tk.HasGlobalEnv,
	"HexToBytes":                            tk.HexToBytes,
	"HexToColor":                            tk.HexToColor,
	"HexToStr":                              tk.HexToStr,
	"HexToInt":                              tk.HexToInt,
	"HexToUInt":                             tk.HexToUInt,
	"HttpRequest":                           tk.HttpRequest,
	"IfFileExists":                          tk.IfFileExists,
	"GetFileSize":                           tk.GetFileSize,
	"GetFileSizeCompact":                    tk.GetFileSizeCompact,
	"IfSwitchExists":                        tk.IfSwitchExists,
	"IfSwitchExistsWhole":                   tk.IfSwitchExistsWhole,
	"IfSwitchExistsWholeI":                  tk.IfSwitchExistsWholeI,
	"InterfaceToStringArray":                tk.InterfaceToStringArray,
	"InStrings":                             tk.InStrings,
	"InStringsTrim":                         tk.InStringsTrim,
	"IndexInStringList":                     tk.IndexInStringList,
	"IndexInStringListFromEnd":              tk.IndexInStringListFromEnd,
	"Int64ArrayToFloat64Array":              tk.Int64ArrayToFloat64Array,
	"Int64ToStr":                            tk.Int64ToStr,
	"IntToKMGT":                             tk.IntToKMGT,
	"IntToStr":                              tk.IntToStr,
	"IntToStrX":                             tk.IntToStrX,
	"IntToWYZ":                              tk.IntToWYZ,
	"IsDirectory":                           tk.IsDirectory,
	"IsEmptyTrim":                           tk.IsEmptyTrim,
	"IsErrStr":                              tk.IsErrStr,
	"IsErrStrX":                             tk.IsErrStrX,
	"IsError":                               tk.IsError,
	"IsErrorString":                         tk.IsErrorString,
	"IsFile":                                tk.IsFile,
	"IsFloat64NearlyEqual":                  tk.IsFloat64NearlyEqual,
	"IsNil":                                 tk.IsNil,
	"IsNilOrEmpty":                          tk.IsNilOrEmpty,
	"IsValidEmail":                          tk.IsValidEmail,
	"IsYesterday":                           tk.IsYesterday,
	"JSONToMapStringFloat64Array":           tk.JSONToMapStringFloat64Array,
	"JSONToMapStringString":                 tk.JSONToMapStringString,
	"JSONToMapStringStringArray":            tk.JSONToMapStringStringArray,
	"JSONToObject":                          tk.JSONToObject,
	"JSONToObjectE":                         tk.JSONToObjectE,
	"JSONToStringArray":                     tk.JSONToStringArray,
	"JoinDualList":                          tk.JoinDualList,
	"JoinLines":                             tk.JoinLines,
	"JoinList":                              tk.JoinList,
	"JoinLinesBySeparator":                  tk.JoinLinesBySeparator,
	"JoinPath":                              tk.JoinPath,
	"GetFileInfo":                           tk.GetFileInfo,
	"JoinURL":                               tk.JoinURL,
	"KindOfValueReflect":                    tk.KindOfValueReflect,
	"Len64":                                 tk.Len64,
	"LessI":                                 tk.LessI,
	"LimitPrecision":                        tk.LimitPrecision,
	"LimitString":                           tk.LimitString,
	"LoadBytes":                             tk.LoadBytes,
	"LoadBytesFromFile":                     tk.LoadBytesFromFile,
	"LoadBytesFromFileE":                    tk.LoadBytesFromFileE,
	"LoadDualLineList":                      tk.LoadDualLineList,
	"LoadDualLineListFromString":            tk.LoadDualLineListFromString,
	"LoadJSONFromFile":                      tk.LoadJSONFromFile,
	"LoadJSONFromString":                    tk.LoadJSONFromString,
	"LoadJSONMapStringFloat64ArrayFromFile": tk.LoadJSONMapStringFloat64ArrayFromFile,
	"LoadMSSFromJSONFile":                   tk.LoadMSSFromJSONFile,
	"LoadSimpleMapFromDir":                  tk.LoadSimpleMapFromDir,
	"LoadSimpleMapFromFile":                 tk.LoadSimpleMapFromFile,
	"LoadSimpleMapFromFileE":                tk.LoadSimpleMapFromFileE,
	"LoadSimpleMapFromString":               tk.LoadSimpleMapFromString,
	"LoadSimpleMapFromStringE":              tk.LoadSimpleMapFromStringE,
	"LoadStringFromFile":                    tk.LoadStringFromFile,
	"LoadStringFromFileB":                   tk.LoadStringFromFileB,
	"LoadStringFromFileE":                   tk.LoadStringFromFileE,
	"LoadStringFromFileWithDefault":         tk.LoadStringFromFileWithDefault,
	"LoadStringList":                        tk.LoadStringList,
	"LoadStringListAsMap":                   tk.LoadStringListAsMap,
	"LoadStringListAsMapRemoveEmpty":        tk.LoadStringListAsMapRemoveEmpty,
	"LoadStringListBuffered":                tk.LoadStringListBuffered,
	"LoadStringListFromFile":                tk.LoadStringListFromFile,
	"LoadStringListRemoveEmpty":             tk.LoadStringListRemoveEmpty,
	"LoadStringTX":                          tk.LoadStringTX,
	"LogWithTime":                           tk.LogWithTime,
	"LogWithTimeCompact":                    tk.LogWithTimeCompact,
	"Ls":                                    tk.Ls,
	"Lsr":                                   tk.Lsr,
	"MD5Encrypt":                            tk.MD5Encrypt,
	"MD5EncryptToBytes":                     tk.MD5EncryptToBytes,
	"MSSFromJSON":                           tk.MSSFromJSON,
	"NewXMLFromMSI":                         tk.NewXMLFromMSI,
	"NewXMLFromAny":                         tk.NewXMLFromAny,
	"NewStringRing":                         tk.NewStringRing,
	"NewObject":                             tk.NewObject,
	"NewRandomGenerator":                    tk.NewRandomGenerator,
	"MSI2MSS":                               tk.MSI2MSS,
	"MSS2MSI":                               tk.MSS2MSI,
	"NewSSHClient":                          tk.NewSSHClient,
	"NewTK":                                 tk.NewTK,
	"NilToEmptyStr":                         tk.NilToEmptyStr,
	"NowToFileName":                         tk.NowToFileName,
	"ReplaceHtmlByMap":                      tk.ReplaceHtmlByMap,
	"CleanHtmlPlaceholders":                 tk.CleanHtmlPlaceholders,
	"NowToStrUTC":                           tk.NowToStrUTC,
	"ObjectsToS":                            tk.ObjectsToS,
	"ObjectToJSON":                          tk.ObjectToJSON,
	"ObjectToJSONIndent":                    tk.ObjectToJSONIndent,
	"PadString":                             tk.PadString,
	"ParseCommandLine":                      tk.ParseCommandLine,
	"ParseHexColor":                         tk.ParseHexColor,
	"Pass":                                  tk.Pass,
	"PickRandomItem":                        tk.PickRandomItem,
	"Pkcs7Padding":                          tk.Pkcs7Padding,
	"Pl":                                    tk.Pl,
	"PlAndExit":                             tk.PlAndExit,
	"PlErr":                                 tk.PlErr,
	"PlErrAndExit":                          tk.PlErrAndExit,
	"PlErrSimple":                           tk.PlErrSimple,
	"PlErrSimpleAndExit":                    tk.PlErrSimpleAndExit,
	"PlErrString":                           tk.PlErrString,
	"PlErrWithPrefix":                       tk.PlErrWithPrefix,
	"PlNow":                                 tk.PlNow,
	"PlSimpleErrorString":                   tk.PlSimpleErrorString,
	"PlTXErr":                               tk.PlTXErr,
	"PlVerbose":                             tk.PlVerbose,
	"Pln":                                   tk.Pln,
	"Plv":                                   tk.Plv,
	"PlvWithError":                          tk.PlvWithError,
	"Plvs":                                  tk.Plvs,
	"Plvsr":                                 tk.Plvsr,
	"Plvx":                                  tk.Plvx,
	"MapToPostData":                         tk.MapToPostData,
	"MapToPostDataI":                        tk.MapToPostDataI,
	"PostRequest":                           tk.PostRequest,
	"PostRequestBytesWithCookieX":           tk.PostRequestBytesWithCookieX,
	"PostRequestBytesWithMSSHeaderX":        tk.PostRequestBytesWithMSSHeaderX,
	"PostRequestBytesX":                     tk.PostRequestBytesX,
	"PostRequestX":                          tk.PostRequestX,
	"Pr":                                    tk.Pr,
	"Prf":                                   tk.Prf,
	"Printf":                                tk.Printf,
	"Printfln":                              tk.Printfln,
	"Prl":                                   tk.Prl,
	"PutRequestX":                           tk.PutRequestX,
	"FindSubStringAll":                      tk.FindSubStringAll,
	"Randomize":                             tk.Randomize,
	"ReadLineFromBufioReader":               tk.ReadLineFromBufioReader,
	"RegContainsIn":                         tk.RegContainsIn,
	"RegSplit":                              tk.RegSplit,
	"RegSplitX":                             tk.RegSplitX,
	"RegContains":                           tk.RegContains,
	"RegContainsX":                          tk.RegContainsX,
	"RegFindAll":                            tk.RegFindAll,
	"RegFindAllGroups":                      tk.RegFindAllGroups,
	"RegFindAllGroupsX":                     tk.RegFindAllGroupsX,
	"RegFindAllX":                           tk.RegFindAllX,
	"RegFindFirst":                          tk.RegFindFirst,
	"RegFindFirstIndex":                     tk.RegFindFirstIndex,
	"RegFindFirstIndexX":                    tk.RegFindFirstIndexX,
	"RegFindAllIndexX":                      tk.RegFindAllIndexX,
	"RegFindFirstTX":                        tk.RegFindFirstTX,
	"RegFindFirstX":                         tk.RegFindFirstX,
	"RegMatch":                              tk.RegMatch,
	"RegMatchX":                             tk.RegMatchX,
	"RegReplace":                            tk.RegReplace,
	"RegReplaceX":                           tk.RegReplaceX,
	"RegStartsWith":                         tk.RegStartsWith,
	"RegStartsWithX":                        tk.RegStartsWithX,
	"RemoveBOM":                             tk.RemoveBOM,
	"RemoveDuplicateInDualLineList":         tk.RemoveDuplicateInDualLineList,
	"RemoveFile":                            tk.RemoveFile,
	"RenameFile":                            tk.RenameFile,
	"RemoveFileExt":                         tk.RemoveFileExt,
	"RemoveGlobalEnv":                       tk.RemoveGlobalEnv,
	"RemoveHtmlTags":                        tk.RemoveHtmlTags,
	"RemoveHtmlTagsX":                       tk.RemoveHtmlTagsX,
	"RemoveItemsInArray":                    tk.RemoveItemsInArray,
	"RemoveLastSubString":                   tk.RemoveLastSubString,
	"Replace":                               tk.Replace,
	"ReplaceLineEnds":                       tk.ReplaceLineEnds,
	"RequestX":                              tk.RequestX,
	"ReshapeXML":                            tk.ReshapeXML,
	"RestoreLineEnds":                       tk.RestoreLineEnds,
	"RuneLen":                               tk.RuneLen,
	"RunWinFileWithSystemDefault":           tk.RunWinFileWithSystemDefault,
	"SafelyGetStringInArray":                tk.SafelyGetStringInArray,
	"SafelyGetFloat64ForKeyWithDefault":     tk.SafelyGetFloat64ForKeyWithDefault,
	"SafelyGetIntForKeyWithDefault":         tk.SafelyGetIntForKeyWithDefault,
	"SafelyGetStringForKeyWithDefault":      tk.SafelyGetStringForKeyWithDefault,
	"SaveBytesToFile":                       tk.SaveBytesToFile,
	"SaveBytesToFileE":                      tk.SaveBytesToFileE,
	"SaveDualLineList":                      tk.SaveDualLineList,
	"SaveJSONIndentToFile":                  tk.SaveJSONIndentToFile,
	"SaveJSONToFile":                        tk.SaveJSONToFile,
	"SaveMSSToJSONFile":                     tk.SaveMSSToJSONFile,
	"SaveSimpleMapToFile":                   tk.SaveSimpleMapToFile,
	"SaveStringList":                        tk.SaveStringList,
	"SaveStringListBuffered":                tk.SaveStringListBuffered,
	"SaveStringListBufferedByRange":         tk.SaveStringListBufferedByRange,
	"SaveStringListWin":                     tk.SaveStringListWin,
	"SaveStringToFile":                      tk.SaveStringToFile,
	"SaveStringToFileE":                     tk.SaveStringToFileE,
	"SetClipText":                           tk.SetClipText,
	"SetFileVar":                            tk.SetFileVar,
	"SetGlobalEnv":                          tk.SetGlobalEnv,
	"SetLogFile":                            tk.SetLogFile,
	"SetValue":                              tk.SetValue,
	"SetVar":                                tk.SetVar,
	"ShuffleStringArray":                    tk.ShuffleStringArray,
	"SimpleMapToString":                     tk.SimpleMapToString,
	"Sleep":                                 tk.Sleep,
	"SleepMilliSeconds":                     tk.SleepMilliSeconds,
	"SleepSeconds":                          tk.SleepSeconds,
	"SortX":                                 tk.SortX,
	"SortMSSArray":                          tk.SortMSSArray,
	"Split":                                 tk.Split,
	"SplitLines":                            tk.SplitLines,
	"SplitLinesRemoveEmpty":                 tk.SplitLinesRemoveEmpty,
	"SplitN":                                tk.SplitN,
	"SplitByLen":                            tk.SplitByLen,
	"SplitByLenRune":                        tk.SplitByLenRune,
	"Spr":                                   tk.Spr,
	"StartSocksClient":                      tk.StartSocksClient,
	"StartSocksServer":                      tk.StartSocksServer,
	"StartTransparentProxy":                 tk.StartTransparentProxy,
	"StartTransparentProxy2":                tk.StartTransparentProxy2,
	"StartsWith":                            tk.StartsWith,
	"StartsWithBOM":                         tk.StartsWithBOM,
	"StartsWithDigit":                       tk.StartsWithDigit,
	"StartsWithIgnoreCase":                  tk.StartsWithIgnoreCase,
	"StartsWithUpper":                       tk.StartsWithUpper,
	"StrToBool":                             tk.StrToBool,
	"ToBool":                                tk.ToBool,
	"StrToFloat64":                          tk.StrToFloat64,
	"StrToFloat64E":                         tk.StrToFloat64E,
	"StrToFloat64WithDefaultValue":          tk.StrToFloat64WithDefaultValue,
	"StrToInt":                              tk.StrToInt,
	"StrToInt64":                            tk.StrToInt64,
	"StrToInt64WithDefaultValue":            tk.StrToInt64WithDefaultValue,
	"StrToIntE":                             tk.StrToIntE,
	"StrToIntPositive":                      tk.StrToIntPositive,
	"StrToIntWithDefaultValue":              tk.StrToIntWithDefaultValue,
	"StrToIntQuick":                         tk.StrToIntQuick,
	"StrToHex":                              tk.StrToHex,
	"ToHex":                                 tk.ToHex,
	"ToByte":                                tk.ToByte,
	"GetChar":                               tk.GetChar,
	"GetChar2":                              tk.GetChar2,
	"GetSystemEndian":                       tk.GetSystemEndian,
	"GetSystemEndianInt":                    tk.GetSystemEndianInt,
	"IfThenElse":                            tk.IfThenElse,
	"StrToTime":                             tk.StrToTime,
	"StrToTimeByFormat":                     tk.StrToTimeByFormat,
	"StrToTimeCompact":                      tk.StrToTimeCompact,
	"StrToTimeCompactNoError":               tk.StrToTimeCompactNoError,
	"DataToBytes":                           tk.DataToBytes,
	"BytesToData":                           tk.BytesToData,
	"StringReplace":                         tk.StringReplace,
	"SumBytes":                              tk.SumBytes,
	"SystemCmd":                             tk.SystemCmd,
	"TKX":                                   tk.TKX,
	"TXResultFromString":                    tk.TXResultFromString,
	"TXResultFromStringE":                   tk.TXResultFromStringE,
	"TXResultFromStringSafely":              tk.TXResultFromStringSafely,
	"TableToMSSArray":                       tk.TableToMSSArray,
	"TableToMSSJSON":                        tk.TableToMSSJSON,
	"TableToMSSMap":                         tk.TableToMSSMap,
	"TableToMSSMapArray":                    tk.TableToMSSMapArray,
	"TimeFormat":                            tk.TimeFormat,
	"TimeFormatCompact":                     tk.TimeFormatCompact,
	"TimeFormatCompact2":                    tk.TimeFormatCompact2,
	"TimeFormatMS":                          tk.TimeFormatMS,
	"TimeFormatMSCompact":                   tk.TimeFormatMSCompact,
	"ToXML":                                 tk.ToXML,
	"ToTime":                                tk.ToTime,
	"ToVar":                                 tk.ToVar,
	"ToInterface":                           tk.ToInterface,
	"ToPointer":                             tk.ToPointer,
	"ToFloat":                               tk.ToFloat,
	"ToInt":                                 tk.ToInt,
	"ToIntI":                                tk.ToIntI,
	"ToRune":                                tk.ToRune,
	"ToJSON":                                tk.ToJSON,
	"ToJSONIndent":                          tk.ToJSONIndent,
	"ToJSONIndentWithDefault":               tk.ToJSONIndentWithDefault,
	"ToJSONWithDefault":                     tk.ToJSONWithDefault,
	"ToJSONX":                               tk.ToJSONX,
	"ToLower":                               tk.ToLower,
	"ToPointerFloat64Array":                 tk.ToPointerFloat64Array,
	"ToPointerStringArray":                  tk.ToPointerStringArray,
	"ToStr":                                 tk.ToStr,
	"ToUpper":                               tk.ToUpper,
	"Trim":                                  tk.Trim,
	"TrimSafely":                            tk.TrimSafely,
	"TrimCharSet":                           tk.TrimCharSet,
	"TypeOfValue":                           tk.TypeOfValue,
	"TypeOfValueReflect":                    tk.TypeOfValueReflect,
	"UrlDecode":                             tk.UrlDecode,
	"UrlEncode":                             tk.UrlEncode,
	"UrlEncode2":                            tk.UrlEncode2,
	"WrapError":                             tk.WrapError,
	"WriteResponse":                         tk.WriteResponse,
	"SetResponseHeader":                     tk.SetResponseHeader,
	"WriteResponseHeader":                   tk.WriteResponseHeader,
	"ConnectPop3":                           tk.ConnectPop3,
	"DialPop3":                              tk.DialPop3,
	"DialPop3TLS":                           tk.DialPop3TLS,
	"NewPop3Client":                         tk.NewPop3Client,

	"GenerateQR": tk.GenerateQR,

	"GetZip":          tk.GetZip,
	"GetZipArchiver":  tk.GetZipArchiver,
	"GetEncZipReader": tk.GetEncZipReader,
	"GetEncZipWriter": tk.GetEncZipWriter,

	"MAX_INT": tk.MAX_INT,

	"ServerInfo":   qlang.StructOf((*tk.ServerInfo)(nil)),
	"SimpleEvent":  qlang.StructOf((*tk.SimpleEvent)(nil)),
	"TK":           qlang.StructOf((*tk.TK)(nil)),
	"TXCollection": qlang.StructOf((*tk.TXCollection)(nil)),
	"TXResult":     qlang.StructOf((*tk.TXResult)(nil)),
	"TXString":     qlang.StructOf((*tk.TXString)(nil)),
}
