package screenshot

import (
	"github.com/kbinani/screenshot"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/kbinani/screenshot",

	"Capture":           screenshot.Capture,
	"CaptureDisplay":    screenshot.CaptureDisplay,
	"CaptureRect":       screenshot.CaptureRect,
	"GetDisplayBounds":  screenshot.GetDisplayBounds,
	"NumActiveDisplays": screenshot.NumActiveDisplays,
}
