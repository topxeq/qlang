package container

import (
	"fyne.io/fyne/v2/container"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "fyne.io/fyne/v2/container",

	"ScrollBoth":           container.ScrollBoth,
	"ScrollHorizontalOnly": container.ScrollHorizontalOnly,
	"ScrollVerticalOnly":   container.ScrollVerticalOnly,
	"TabLocationBottom":    container.TabLocationBottom,
	"TabLocationLeading":   container.TabLocationLeading,
	"TabLocationTop":       container.TabLocationTop,
	"TabLocationTrailing":  container.TabLocationTrailing,

	"New":                container.New,
	"NewAdaptiveGrid":    container.NewAdaptiveGrid,
	"NewBorder":          container.NewBorder,
	"NewCenter":          container.NewCenter,
	"NewGridWithColumns": container.NewGridWithColumns,
	"NewGridWithRows":    container.NewGridWithRows,
	"NewGridWrap":        container.NewGridWrap,
	"NewHBox":            container.NewHBox,
	"NewMax":             container.NewMax,
	"NewPadded":          container.NewPadded,
	"NewVBox":            container.NewVBox,
	"NewWithoutLayout":   container.NewWithoutLayout,

	"apptabs":            container.NewAppTabs,
	"NewAppTabs":         container.NewAppTabs,
	"Split":              qlang.StructOf((*container.Split)(nil)),
	"NewHSplit":          container.NewHSplit,
	"NewVSplit":          container.NewVSplit,
	"TabItem":            qlang.StructOf((*container.TabItem)(nil)),
	"NewTabItem":         container.NewTabItem,
	"NewTabItemWithIcon": container.NewTabItemWithIcon,
}
