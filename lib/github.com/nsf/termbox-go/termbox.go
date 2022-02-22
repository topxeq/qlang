package termbox

import (
	"github.com/nsf/termbox-go"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/nsf/termbox-go",

	"AttrBlink":         termbox.AttrBlink,
	"AttrBold":          termbox.AttrBold,
	"AttrCursive":       termbox.AttrCursive,
	"AttrDim":           termbox.AttrDim,
	"AttrHidden":        termbox.AttrHidden,
	"AttrReverse":       termbox.AttrReverse,
	"AttrUnderline":     termbox.AttrUnderline,
	"ColorBlack":        termbox.ColorBlack,
	"ColorBlue":         termbox.ColorBlue,
	"ColorCyan":         termbox.ColorCyan,
	"ColorDarkGray":     termbox.ColorDarkGray,
	"ColorDefault":      termbox.ColorDefault,
	"ColorGreen":        termbox.ColorGreen,
	"ColorLightBlue":    termbox.ColorLightBlue,
	"ColorLightCyan":    termbox.ColorLightCyan,
	"ColorLightGray":    termbox.ColorLightGray,
	"ColorLightGreen":   termbox.ColorLightGreen,
	"ColorLightMagenta": termbox.ColorLightMagenta,
	"ColorLightRed":     termbox.ColorLightRed,
	"ColorLightYellow":  termbox.ColorLightYellow,
	"ColorMagenta":      termbox.ColorMagenta,
	"ColorRed":          termbox.ColorRed,
	"ColorWhite":        termbox.ColorWhite,
	"ColorYellow":       termbox.ColorYellow,
	"EventError":        termbox.EventError,
	"EventInterrupt":    termbox.EventInterrupt,
	"EventKey":          termbox.EventKey,
	"EventMouse":        termbox.EventMouse,
	"EventNone":         termbox.EventNone,
	"EventRaw":          termbox.EventRaw,
	"EventResize":       termbox.EventResize,
	"InputAlt":          termbox.InputAlt,
	"InputCurrent":      termbox.InputCurrent,
	"InputEsc":          termbox.InputEsc,
	"InputMouse":        termbox.InputMouse,
	"KeyArrowDown":      termbox.KeyArrowDown,
	"KeyArrowLeft":      termbox.KeyArrowLeft,
	"KeyArrowRight":     termbox.KeyArrowRight,
	"KeyArrowUp":        termbox.KeyArrowUp,
	"KeyBackspace":      termbox.KeyBackspace,
	"KeyBackspace2":     termbox.KeyBackspace2,
	"KeyCtrl2":          termbox.KeyCtrl2,
	"KeyCtrl3":          termbox.KeyCtrl3,
	"KeyCtrl4":          termbox.KeyCtrl4,
	"KeyCtrl5":          termbox.KeyCtrl5,
	"KeyCtrl6":          termbox.KeyCtrl6,
	"KeyCtrl7":          termbox.KeyCtrl7,
	"KeyCtrl8":          termbox.KeyCtrl8,
	"KeyCtrlA":          termbox.KeyCtrlA,
	"KeyCtrlB":          termbox.KeyCtrlB,
	"KeyCtrlBackslash":  termbox.KeyCtrlBackslash,
	"KeyCtrlC":          termbox.KeyCtrlC,
	"KeyCtrlD":          termbox.KeyCtrlD,
	"KeyCtrlE":          termbox.KeyCtrlE,
	"KeyCtrlF":          termbox.KeyCtrlF,
	"KeyCtrlG":          termbox.KeyCtrlG,
	"KeyCtrlH":          termbox.KeyCtrlH,
	"KeyCtrlI":          termbox.KeyCtrlI,
	"KeyCtrlJ":          termbox.KeyCtrlJ,
	"KeyCtrlK":          termbox.KeyCtrlK,
	"KeyCtrlL":          termbox.KeyCtrlL,
	"KeyCtrlLsqBracket": termbox.KeyCtrlLsqBracket,
	"KeyCtrlM":          termbox.KeyCtrlM,
	"KeyCtrlN":          termbox.KeyCtrlN,
	"KeyCtrlO":          termbox.KeyCtrlO,
	"KeyCtrlP":          termbox.KeyCtrlP,
	"KeyCtrlQ":          termbox.KeyCtrlQ,
	"KeyCtrlR":          termbox.KeyCtrlR,
	"KeyCtrlRsqBracket": termbox.KeyCtrlRsqBracket,
	"KeyCtrlS":          termbox.KeyCtrlS,
	"KeyCtrlSlash":      termbox.KeyCtrlSlash,
	"KeyCtrlSpace":      termbox.KeyCtrlSpace,
	"KeyCtrlT":          termbox.KeyCtrlT,
	"KeyCtrlTilde":      termbox.KeyCtrlTilde,
	"KeyCtrlU":          termbox.KeyCtrlU,
	"KeyCtrlUnderscore": termbox.KeyCtrlUnderscore,
	"KeyCtrlV":          termbox.KeyCtrlV,
	"KeyCtrlW":          termbox.KeyCtrlW,
	"KeyCtrlX":          termbox.KeyCtrlX,
	"KeyCtrlY":          termbox.KeyCtrlY,
	"KeyCtrlZ":          termbox.KeyCtrlZ,
	"KeyDelete":         termbox.KeyDelete,
	"KeyEnd":            termbox.KeyEnd,
	"KeyEnter":          termbox.KeyEnter,
	"KeyEsc":            termbox.KeyEsc,
	"KeyF1":             termbox.KeyF1,
	"KeyF10":            termbox.KeyF10,
	"KeyF11":            termbox.KeyF11,
	"KeyF12":            termbox.KeyF12,
	"KeyF2":             termbox.KeyF2,
	"KeyF3":             termbox.KeyF3,
	"KeyF4":             termbox.KeyF4,
	"KeyF5":             termbox.KeyF5,
	"KeyF6":             termbox.KeyF6,
	"KeyF7":             termbox.KeyF7,
	"KeyF8":             termbox.KeyF8,
	"KeyF9":             termbox.KeyF9,
	"KeyHome":           termbox.KeyHome,
	"KeyInsert":         termbox.KeyInsert,
	"KeyPgdn":           termbox.KeyPgdn,
	"KeyPgup":           termbox.KeyPgup,
	"KeySpace":          termbox.KeySpace,
	"KeyTab":            termbox.KeyTab,
	"ModAlt":            termbox.ModAlt,
	"ModMotion":         termbox.ModMotion,
	"MouseLeft":         termbox.MouseLeft,
	"MouseMiddle":       termbox.MouseMiddle,
	"MouseRelease":      termbox.MouseRelease,
	"MouseRight":        termbox.MouseRight,
	"MouseWheelDown":    termbox.MouseWheelDown,
	"MouseWheelUp":      termbox.MouseWheelUp,
	"Output216":         termbox.Output216,
	"Output256":         termbox.Output256,
	"OutputCurrent":     termbox.OutputCurrent,
	"OutputGrayscale":   termbox.OutputGrayscale,
	"OutputNormal":      termbox.OutputNormal,
	"OutputRGB":         termbox.OutputRGB,

	"IsInit": termbox.IsInit,

	"AttributeToRGB": termbox.AttributeToRGB,
	"Clear":          termbox.Clear,
	"Close":          termbox.Close,
	"Flush":          termbox.Flush,
	"HideCursor":     termbox.HideCursor,
	"Init":           termbox.Init,
	"Interrupt":      termbox.Interrupt,
	"SetBg":          termbox.SetBg,
	"SetCell":        termbox.SetCell,
	"SetChar":        termbox.SetChar,
	"SetCursor":      termbox.SetCursor,
	"SetFg":          termbox.SetFg,
	"Size":           termbox.Size,
	"Sync":           termbox.Sync,

	"Cell":       qlang.StructOf((*termbox.Cell)(nil)),
	"CellBuffer": termbox.CellBuffer,
	"GetCell":    termbox.GetCell,
	"Event":      qlang.StructOf((*termbox.Event)(nil)),
	"PollEvent":  termbox.PollEvent,
}
