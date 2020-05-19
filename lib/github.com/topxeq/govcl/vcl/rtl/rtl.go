package rtl

import (
	"github.com/topxeq/govcl/vcl/rtl"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/govcl/vcl/rtl",

	"PathSeparator": rtl.PathSeparator,

	"SysLocale": rtl.SysLocale,

	"Combine":               rtl.Combine,
	"CreateGUID":            rtl.CreateGUID,
	"CurrentThreadId":       rtl.CurrentThreadId,
	"Exclude":               rtl.Exclude,
	"ExtractFileExt":        rtl.ExtractFileExt,
	"ExtractFileName":       rtl.ExtractFileName,
	"ExtractFilePath":       rtl.ExtractFilePath,
	"FileExists":            rtl.FileExists,
	"GUIDToString":          rtl.GUIDToString,
	"GetDC":                 rtl.GetDC,
	"GetFileNameWithoutExt": rtl.GetFileNameWithoutExt,
	"GetLibResouceCount":    rtl.GetLibResouceCount,
	"GetLibResouceItem":     rtl.GetLibResouceItem,
	"GetLibResouceItems":    rtl.GetLibResouceItems,
	"GetStringArrOf":        rtl.GetStringArrOf,
	"HiByte":                rtl.HiByte,
	"HiWord":                rtl.HiWord,
	"InSets":                rtl.InSets,
	"Include":               rtl.Include,
	"IsIconic":              rtl.IsIconic,
	"IsNil":                 rtl.IsNil,
	"IsWindow":              rtl.IsWindow,
	"IsWindowVisible":       rtl.IsWindowVisible,
	"IsZoomed":              rtl.IsZoomed,
	"LcLLoaded":             rtl.LcLLoaded,
	"LibAbout":              rtl.LibAbout,
	"LibStringEncoding":     rtl.LibStringEncoding,
	"LibVersion":            rtl.LibVersion,
	"LocaleIDFromName":      rtl.LocaleIDFromName,
	"MainInstance":          rtl.MainInstance,
	"MainThreadId":          rtl.MainThreadId,
	"MakeLParam":            rtl.MakeLParam,
	"MakeLResult":           rtl.MakeLResult,
	"MakeLong":              rtl.MakeLong,
	"MakeWParam":            rtl.MakeWParam,
	"MakeWord":              rtl.MakeWord,
	"ModifyLibResouce":      rtl.ModifyLibResouce,
	"Move":                  rtl.Move,
	"PointToLParam":         rtl.PointToLParam,
	"PostMessage":           rtl.PostMessage,
	"ReleaseDC":             rtl.ReleaseDC,
	"SendMessage":           rtl.SendMessage,
	"SetForegroundWindow":   rtl.SetForegroundWindow,
	"SetPropertySecValue":   rtl.SetPropertySecValue,
	"SetPropertyValue":      rtl.SetPropertyValue,
	"ShiftStateToWord":      rtl.ShiftStateToWord,
	"ShortCutToText":        rtl.ShortCutToText,
	"StrLen":                rtl.StrLen,
	"StringToGUID":          rtl.StringToGUID,
	"SysOpen":               rtl.SysOpen,
	"TextToShortCut":        rtl.TextToShortCut,
	"WindowFromPoint":       rtl.WindowFromPoint,
}
