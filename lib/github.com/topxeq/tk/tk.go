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

	"DebugModeG":         tk.DebugModeG,
	"GlobalEnvSetG":      tk.GlobalEnvSetG,
	"TimeFormat":         tk.TimeFormat,
	"TimeFormatCompact":  tk.TimeFormatCompact,
	"TimeFormatCompact2": tk.TimeFormatCompact2,
	"TimeFormatMS":       tk.TimeFormatMS,

	"AESDecrypt":                               tk.AESDecrypt,
	"AESEncrypt":                               tk.AESEncrypt,
	"addDebug":                                 tk.AddDebug,
	"addDebugF":                                tk.AddDebugF,
	"addLastSubString":                         tk.AddLastSubString,
	"analyzeCommandLineParamter":               tk.AnalyzeCommandLineParamter,
	"analyzeURLParams":                         tk.AnalyzeURLParams,
	"appendDualLineList":                       tk.AppendDualLineList,
	"appendSimpleMapFromFile":                  tk.AppendSimpleMapFromFile,
	"appendStringToFile":                       tk.AppendStringToFile,
	"byteSliceToStringDec":                     tk.ByteSliceToStringDec,
	"byteToHex":                                tk.ByteToHex,
	"bytesToHex":                               tk.BytesToHex,
	"calCosineSimilarityBetweenFloatsBig":      tk.CalCosineSimilarityBetweenFloatsBig,
	"checkErr":                                 tk.CheckErr,
	"checkErrCompact":                          tk.CheckErrCompact,
	"checkErrf":                                tk.CheckErrf,
	"clearDebug":                               tk.ClearDebug,
	"contains":                                 tk.Contains,
	"containsIgnoreCase":                       tk.ContainsIgnoreCase,
	"containsIn":                               tk.ContainsIn,
	"containsInStringList":                     tk.ContainsInStringList,
	"convertStringToUTF8":                      tk.ConvertStringToUTF8,
	"convertToGB18030":                         tk.ConvertToGB18030,
	"convertToGB18030Bytes":                    tk.ConvertToGB18030Bytes,
	"convertToUTF8":                            tk.ConvertToUTF8,
	"decodeStringCustom":                       tk.DecodeStringCustom,
	"decodeStringSimple":                       tk.DecodeStringSimple,
	"decodeStringUnderline":                    tk.DecodeStringUnderline,
	"decryptDataByTXDEE":                       tk.DecryptDataByTXDEE,
	"decryptDataByTXDEF":                       tk.DecryptDataByTXDEF,
	"decryptFileByTXDEF":                       tk.DecryptFileByTXDEF,
	"decryptFileByTXDEFS":                      tk.DecryptFileByTXDEFS,
	"decryptFileByTXDEFStream":                 tk.DecryptFileByTXDEFStream,
	"decryptFileByTXDEFStreamS":                tk.DecryptFileByTXDEFStreamS,
	"decryptStreamByTXDEF":                     tk.DecryptStreamByTXDEF,
	"decryptStringByTXDEE":                     tk.DecryptStringByTXDEE,
	"decryptStringByTXDEF":                     tk.DecryptStringByTXDEF,
	"decryptStringByTXTE":                      tk.DecryptStringByTXTE,
	"deepClone":                                tk.DeepClone,
	"deepCopyFromTo":                           tk.DeepCopyFromTo,
	"deleteItemInInt64Array":                   tk.DeleteItemInInt64Array,
	"deleteItemInIntArray":                     tk.DeleteItemInIntArray,
	"deleteItemInStringArray":                  tk.DeleteItemInStringArray,
	"downloadBytes":                            tk.DownloadBytes,
	"downloadFile":                             tk.DownloadFile,
	"downloadPage":                             tk.DownloadPage,
	"downloadPageByMap":                        tk.DownloadPageByMap,
	"downloadPageUTF8":                         tk.DownloadPageUTF8,
	"encodeStringCustom":                       tk.EncodeStringCustom,
	"encodeStringSimple":                       tk.EncodeStringSimple,
	"encodeStringUnderline":                    tk.EncodeStringUnderline,
	"encodeToXMLString":                        tk.EncodeToXMLString,
	"encryptDataByTXDEE":                       tk.EncryptDataByTXDEE,
	"encryptDataByTXDEF":                       tk.EncryptDataByTXDEF,
	"encryptFileByTXDEF":                       tk.EncryptFileByTXDEF,
	"encryptFileByTXDEFS":                      tk.EncryptFileByTXDEFS,
	"encryptFileByTXDEFStream":                 tk.EncryptFileByTXDEFStream,
	"encryptFileByTXDEFStreamS":                tk.EncryptFileByTXDEFStreamS,
	"encryptFileByTXDEFWithHeader":             tk.EncryptFileByTXDEFWithHeader,
	"encryptStreamByTXDEF":                     tk.EncryptStreamByTXDEF,
	"encryptStringByTXDEE":                     tk.EncryptStringByTXDEE,
	"encryptStringByTXDEF":                     tk.EncryptStringByTXDEF,
	"encryptStringByTXTE":                      tk.EncryptStringByTXTE,
	"endsWith":                                 tk.EndsWith,
	"endsWithIgnoreCase":                       tk.EndsWithIgnoreCase,
	"ensureMakeDirs":                           tk.EnsureMakeDirs,
	"ensureMakeDirsE":                          tk.EnsureMakeDirsE,
	"ensureValidFileNameX":                     tk.EnsureValidFileNameX,
	"errf":                                     tk.Errf,
	"errorStringToError":                       tk.ErrorStringToError,
	"errorToString":                            tk.ErrorToString,
	"exit":                                     tk.Exit,
	"fatalErr":                                 tk.FatalErr,
	"fatalErrf":                                tk.FatalErrf,
	"fatalf":                                   tk.Fatalf,
	"findFirstDiffIndex":                       tk.FindFirstDiffIndex,
	"findSamePrefix":                           tk.FindSamePrefix,
	"float32ArrayToFloat64Array":               tk.Float32ArrayToFloat64Array,
	"float64ToStr":                             tk.Float64ToStr,
	"formatStringSliceSlice":                   tk.FormatStringSliceSlice,
	"formatTime":                               tk.FormatTime,
	"fpl":                                      tk.Fpl,
	"fpr":                                      tk.Fpr,
	"fromJSON":                                 tk.FromJSON,
	"generateErrorString":                      tk.GenerateErrorString,
	"generateErrorStringF":                     tk.GenerateErrorStringF,
	"generateFileListInDir":                    tk.GenerateFileListInDir,
	"generateFileListRecursively":              tk.GenerateFileListRecursively,
	"generateFileListRecursivelyWithExclusive": tk.GenerateFileListRecursivelyWithExclusive,
	"generateJSONPResponse":                    tk.GenerateJSONPResponse,
	"generateJSONPResponseWith2Object":         tk.GenerateJSONPResponseWith2Object,
	"generateJSONPResponseWith3Object":         tk.GenerateJSONPResponseWith3Object,
	"generateJSONPResponseWithObject":          tk.GenerateJSONPResponseWithObject,
	"generateRandomString":                     tk.GenerateRandomString,
	"getAllParameters":                         tk.GetAllParameters,
	"getAllSwitches":                           tk.GetAllSwitches,
	"getApplicationPath":                       tk.GetApplicationPath,
	"getAvailableFileName":                     tk.GetAvailableFileName,
	"getCurrentDir":                            tk.GetCurrentDir,
	"getCurrentThreadID":                       tk.GetCurrentThreadID,
	"getDBConnection":                          tk.GetDBConnection,
	"getDBResultArray":                         tk.GetDBResultArray,
	"getDBResultString":                        tk.GetDBResultString,
	"getDBRowCount":                            tk.GetDBRowCount,
	"getDBRowCountCompact":                     tk.GetDBRowCountCompact,
	"getDebug":                                 tk.GetDebug,
	"getDirOfFilePath":                         tk.GetDirOfFilePath,
	"getEnv":                                   tk.GetEnv,
	"getErrorString":                           tk.GetErrorString,
	"getErrorStringSafely":                     tk.GetErrorStringSafely,
	"getFileExt":                               tk.GetFileExt,
	"getFilePathSeperator":                     tk.GetFilePathSeperator,
	"getFormValueWithDefaultValue":             tk.GetFormValueWithDefaultValue,
	"getGlobalEnvList":                         tk.GetGlobalEnvList,
	"getGlobalEnvString":                       tk.GetGlobalEnvString,
	"getInputBufferedScan":                     tk.GetInputBufferedScan,
	"getInputf":                                tk.GetInputf,
	"getJSONNode":                              tk.GetJSONNode,
	"getJSONNodeAny":                           tk.GetJSONNodeAny,
	"getJSONSubNode":                           tk.GetJSONSubNode,
	"getJSONSubNodeAny":                        tk.GetJSONSubNodeAny,
	"getLastComponentOfFilePath":               tk.GetLastComponentOfFilePath,
	"getLastComponentOfUrl":                    tk.GetLastComponentOfUrl,
	"getNowDateString":                         tk.GetNowDateString,
	"getNowMinutesInDay":                       tk.GetNowMinutesInDay,
	"getNowTimeOnlyStringBeijing":              tk.GetNowTimeOnlyStringBeijing,
	"getNowTimeString":                         tk.GetNowTimeString,
	"getNowTimeStringFormal":                   tk.GetNowTimeStringFormal,
	"getNowTimeStringFormat":                   tk.GetNowTimeStringFormat,
	"getNowTimeStringHourMinute":               tk.GetNowTimeStringHourMinute,
	"getOSName":                                tk.GetOSName,
	"getParameterByIndexWithDefaultValue":      tk.GetParameterByIndexWithDefaultValue,
	"getRandomByte":                            tk.GetRandomByte,
	"getRandomInt64InRange":                    tk.GetRandomInt64InRange,
	"getRandomInt64LessThan":                   tk.GetRandomInt64LessThan,
	"getRandomIntInRange":                      tk.GetRandomIntInRange,
	"getRandomIntLessThan":                     tk.GetRandomIntLessThan,
	"getRandomSubDualList":                     tk.GetRandomSubDualList,
	"getRandomizeInt64ArrayCopy":               tk.GetRandomizeInt64ArrayCopy,
	"getRandomizeIntArrayCopy":                 tk.GetRandomizeIntArrayCopy,
	"getRandomizeStringArrayCopy":              tk.GetRandomizeStringArrayCopy,
	"getRuntimeStack":                          tk.GetRuntimeStack,
	"getSliceMaxLen":                           tk.GetSliceMaxLen,
	"getStringSliceFilled":                     tk.GetStringSliceFilled,
	"getSuccessValue":                          tk.GetSuccessValue,
	"getSwitchWithDefaultInt64Value":           tk.GetSwitchWithDefaultInt64Value,
	"getSwitchWithDefaultIntValue":             tk.GetSwitchWithDefaultIntValue,
	"getSwitchWithDefaultValue":                tk.GetSwitchWithDefaultValue,
	"getTimeFromUnixTimeStamp":                 tk.GetTimeFromUnixTimeStamp,
	"getTimeFromUnixTimeStampMid":              tk.GetTimeFromUnixTimeStampMid,
	"getTimeStamp":                             tk.GetTimeStamp,
	"getTimeStampMid":                          tk.GetTimeStampMid,
	"getTimeStampNano":                         tk.GetTimeStampNano,
	"getTimeStringDiffMS":                      tk.GetTimeStringDiffMS,
	"getUserInput":                             tk.GetUserInput,
	"getValueOfMSS":                            tk.GetValueOfMSS,
	"hasGlobalEnv":                             tk.HasGlobalEnv,
	"hexToBytes":                               tk.HexToBytes,
	"hexToInt":                                 tk.HexToInt,
	"ifFileExists":                             tk.IfFileExists,
	"ifSwitchExists":                           tk.IfSwitchExists,
	"ifSwitchExistsWhole":                      tk.IfSwitchExistsWhole,
	"inStrings":                                tk.InStrings,
	"indexInStringList":                        tk.IndexInStringList,
	"indexInStringListFromEnd":                 tk.IndexInStringListFromEnd,
	"int64ArrayToFloat64Array":                 tk.Int64ArrayToFloat64Array,
	"int64ToStr":                               tk.Int64ToStr,
	"intToKMGT":                                tk.IntToKMGT,
	"intToStr":                                 tk.IntToStr,
	"intToWYZ":                                 tk.IntToWYZ,
	"isDirectory":                              tk.IsDirectory,
	"isEmptyTrim":                              tk.IsEmptyTrim,
	"isErrorString":                            tk.IsErrorString,
	"isFile":                                   tk.IsFile,
	"isYesterday":                              tk.IsYesterday,
	"JSONToMapStringString":                    tk.JSONToMapStringString,
	"JSONToObject":                             tk.JSONToObject,
	"JSONToObjectE":                            tk.JSONToObjectE,
	"JSONToStringArray":                        tk.JSONToStringArray,
	"joinDualList":                             tk.JoinDualList,
	"joinLines":                                tk.JoinLines,
	"joinLinesBySeparator":                     tk.JoinLinesBySeparator,
	"joinPath":                                 tk.JoinPath,
	"joinURL":                                  tk.JoinURL,
	"len64":                                    tk.Len64,
	"loadBytes":                                tk.LoadBytes,
	"loadBytesFromFileE":                       tk.LoadBytesFromFileE,
	"loadDualLineList":                         tk.LoadDualLineList,
	"loadDualLineListFromString":               tk.LoadDualLineListFromString,
	"loadSimpleMapFromDir":                     tk.LoadSimpleMapFromDir,
	"loadSimpleMapFromFile":                    tk.LoadSimpleMapFromFile,
	"loadSimpleMapFromFileE":                   tk.LoadSimpleMapFromFileE,
	"loadSimpleMapFromString":                  tk.LoadSimpleMapFromString,
	"loadSimpleMapFromStringE":                 tk.LoadSimpleMapFromStringE,
	"loadStringFromFile":                       tk.LoadStringFromFile,
	"loadStringFromFileB":                      tk.LoadStringFromFileB,
	"loadStringFromFileE":                      tk.LoadStringFromFileE,
	"loadStringFromFileWithDefault":            tk.LoadStringFromFileWithDefault,
	"loadStringList":                           tk.LoadStringList,
	"loadStringListBuffered":                   tk.LoadStringListBuffered,
	"loadStringListFromFile":                   tk.LoadStringListFromFile,
	"logWithTime":                              tk.LogWithTime,
	"logWithTimeCompact":                       tk.LogWithTimeCompact,
	"MD5Encrypt":                               tk.MD5Encrypt,
	"nowToFileName":                            tk.NowToFileName,
	"nowToStrUTC":                              tk.NowToStrUTC,
	"objectToJSON":                             tk.ObjectToJSON,
	"objectToJSONIndent":                       tk.ObjectToJSONIndent,
	"parseCommandLine":                         tk.ParseCommandLine,
	"parseHexColor":                            tk.ParseHexColor,
	"pkcs7Padding":                             tk.Pkcs7Padding,
	"pl":                                       tk.Pl,
	"plAndExit":                                tk.PlAndExit,
	"plErr":                                    tk.PlErr,
	"plErrAndExit":                             tk.PlErrAndExit,
	"plErrSimple":                              tk.PlErrSimple,
	"plErrSimpleAndExit":                       tk.PlErrSimpleAndExit,
	"plErrString":                              tk.PlErrString,
	"plErrWithPrefix":                          tk.PlErrWithPrefix,
	"plSimpleErrorString":                      tk.PlSimpleErrorString,
	"plTXErr":                                  tk.PlTXErr,
	"plVerbose":                                tk.PlVerbose,
	"pln":                                      tk.Pln,
	"plv":                                      tk.Plv,
	"plvWithError":                             tk.PlvWithError,
	"plvs":                                     tk.Plvs,
	"plvsr":                                    tk.Plvsr,
	"postRequest":                              tk.PostRequest,
	"postRequestBytesWithCookieX":              tk.PostRequestBytesWithCookieX,
	"postRequestBytesWithMSSHeaderX":           tk.PostRequestBytesWithMSSHeaderX,
	"postRequestBytesX":                        tk.PostRequestBytesX,
	"postRequestX":                             tk.PostRequestX,
	"pr":                                       tk.Pr,
	"prf":                                      tk.Prf,
	"printf":                                   tk.Printf,
	"printfln":                                 tk.Printfln,
	"prl":                                      tk.Prl,
	"randomize":                                tk.Randomize,
	"readLineFromBufioReader":                  tk.ReadLineFromBufioReader,
	"regContains":                              tk.RegContains,
	"regFindAll":                               tk.RegFindAll,
	"regFindFirst":                             tk.RegFindFirst,
	"regFindFirstIndex":                        tk.RegFindFirstIndex,
	"regMatch":                                 tk.RegMatch,
	"regReplace":                               tk.RegReplace,
	"regStartsWith":                            tk.RegStartsWith,
	"removeBOM":                                tk.RemoveBOM,
	"removeDuplicateInDualLineList":            tk.RemoveDuplicateInDualLineList,
	"removeFileExt":                            tk.RemoveFileExt,
	"removeGlobalEnv":                          tk.RemoveGlobalEnv,
	"removeLastSubString":                      tk.RemoveLastSubString,
	"replace":                                  tk.Replace,
	"replaceLineEnds":                          tk.ReplaceLineEnds,
	"restoreLineEnds":                          tk.RestoreLineEnds,
	"runWinFileWithSystemDefault":              tk.RunWinFileWithSystemDefault,
	"safelyGetFloat64ForKeyWithDefault":        tk.SafelyGetFloat64ForKeyWithDefault,
	"safelyGetIntForKeyWithDefault":            tk.SafelyGetIntForKeyWithDefault,
	"safelyGetStringForKeyWithDefault":         tk.SafelyGetStringForKeyWithDefault,
	"saveDualLineList":                         tk.SaveDualLineList,
	"saveSimpleMapToFile":                      tk.SaveSimpleMapToFile,
	"saveStringList":                           tk.SaveStringList,
	"saveStringListBuffered":                   tk.SaveStringListBuffered,
	"saveStringListBufferedByRange":            tk.SaveStringListBufferedByRange,
	"saveStringListWin":                        tk.SaveStringListWin,
	"saveStringToFile":                         tk.SaveStringToFile,
	"saveStringToFileE":                        tk.SaveStringToFileE,
	"setGlobalEnv":                             tk.SetGlobalEnv,
	"setLogFile":                               tk.SetLogFile,
	"setValue":                                 tk.SetValue,
	"shuffleStringArray":                       tk.ShuffleStringArray,
	"simpleMapToString":                        tk.SimpleMapToString,
	"sleepMilliSeconds":                        tk.SleepMilliSeconds,
	"sleepSeconds":                             tk.SleepSeconds,
	"split":                                    tk.Split,
	"splitLines":                               tk.SplitLines,
	"splitLinesRemoveEmpty":                    tk.SplitLinesRemoveEmpty,
	"splitN":                                   tk.SplitN,
	"spr":                                      tk.Spr,
	"startsWith":                               tk.StartsWith,
	"startsWithBOM":                            tk.StartsWithBOM,
	"startsWithDigit":                          tk.StartsWithDigit,
	"startsWithIgnoreCase":                     tk.StartsWithIgnoreCase,
	"startsWithUpper":                          tk.StartsWithUpper,
	"strToBool":                                tk.StrToBool,
	"strToFloat64":                             tk.StrToFloat64,
	"strToFloat64WithDefaultValue":             tk.StrToFloat64WithDefaultValue,
	"strToInt":                                 tk.StrToInt,
	"strToInt64WithDefaultValue":               tk.StrToInt64WithDefaultValue,
	"strToIntPositive":                         tk.StrToIntPositive,
	"strToIntWithDefaultValue":                 tk.StrToIntWithDefaultValue,
	"strToTime":                                tk.StrToTime,
	"strToTimeByFormat":                        tk.StrToTimeByFormat,
	"strToTimeCompact":                         tk.StrToTimeCompact,
	"strToTimeCompactNoError":                  tk.StrToTimeCompactNoError,
	"stringReplace":                            tk.StringReplace,
	"sumBytes":                                 tk.SumBytes,
	"toJSON":                                   tk.ToJSON,
	"toJSONIndent":                             tk.ToJSONIndent,
	"toLower":                                  tk.ToLower,
	"toUpper":                                  tk.ToUpper,
	"trim":                                     tk.Trim,
	"trimCharSet":                              tk.TrimCharSet,
	"urlDecode":                                tk.UrlDecode,
	"urlEncode":                                tk.UrlEncode,
	"urlEncode2":                               tk.UrlEncode2,

	"newRandomGenerator":           tk.NewRandomGenerator,
	"SimpleEvent":                  qlang.StructOf((*tk.SimpleEvent)(nil)),
	"createSimpleEvent":            tk.CreateSimpleEvent,
	"TXCollection":                 qlang.StructOf((*tk.TXCollection)(nil)),
	"createTXCollection":           tk.CreateTXCollection,
	"TXResult":                     qlang.StructOf((*tk.TXResult)(nil)),
	"TXResultFromString":           tk.TXResultFromString,
	"TXString":                     qlang.StructOf((*tk.TXString)(nil)),
	"createString":                 tk.CreateString,
	"createStringEmpty":            tk.CreateStringEmpty,
	"createStringError":            tk.CreateStringError,
	"createStringErrorF":           tk.CreateStringErrorF,
	"createStringErrorFromTXError": tk.CreateStringErrorFromTXError,
	"createStringSimple":           tk.CreateStringSimple,
	"createStringSuccess":          tk.CreateStringSuccess,
	"createStringWithObject":       tk.CreateStringWithObject,
	"generateErrorStringFTX":       tk.GenerateErrorStringFTX,
	"generateErrorStringTX":        tk.GenerateErrorStringTX,
	"loadStringTX":                 tk.LoadStringTX,
	"regFindFirstTX":               tk.RegFindFirstTX,
}