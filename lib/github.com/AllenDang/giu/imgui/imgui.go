package imgui

import (
	"github.com/AllenDang/giu/imgui"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/AllenDang/giu/imgui",

	"BackendFlagHasGamepad":                    imgui.BackendFlagHasGamepad,
	"BackendFlagHasMouseCursors":               imgui.BackendFlagHasMouseCursors,
	"BackendFlagHasSetMousePos":                imgui.BackendFlagHasSetMousePos,
	"BackendFlagNone":                          imgui.BackendFlagNone,
	"BackendFlagsRendererHasVtxOffset":         imgui.BackendFlagsRendererHasVtxOffset,
	"ColorEditFlagsAlphaBar":                   imgui.ColorEditFlagsAlphaBar,
	"ColorEditFlagsAlphaPreview":               imgui.ColorEditFlagsAlphaPreview,
	"ColorEditFlagsAlphaPreviewHalf":           imgui.ColorEditFlagsAlphaPreviewHalf,
	"ColorEditFlagsFloat":                      imgui.ColorEditFlagsFloat,
	"ColorEditFlagsHDR":                        imgui.ColorEditFlagsHDR,
	"ColorEditFlagsHEX":                        imgui.ColorEditFlagsHEX,
	"ColorEditFlagsHSV":                        imgui.ColorEditFlagsHSV,
	"ColorEditFlagsNoAlpha":                    imgui.ColorEditFlagsNoAlpha,
	"ColorEditFlagsNoDragDrop":                 imgui.ColorEditFlagsNoDragDrop,
	"ColorEditFlagsNoInputs":                   imgui.ColorEditFlagsNoInputs,
	"ColorEditFlagsNoLabel":                    imgui.ColorEditFlagsNoLabel,
	"ColorEditFlagsNoOptions":                  imgui.ColorEditFlagsNoOptions,
	"ColorEditFlagsNoPicker":                   imgui.ColorEditFlagsNoPicker,
	"ColorEditFlagsNoSmallPreview":             imgui.ColorEditFlagsNoSmallPreview,
	"ColorEditFlagsNoTooltip":                  imgui.ColorEditFlagsNoTooltip,
	"ColorEditFlagsNone":                       imgui.ColorEditFlagsNone,
	"ColorEditFlagsRGB":                        imgui.ColorEditFlagsRGB,
	"ColorEditFlagsUint8":                      imgui.ColorEditFlagsUint8,
	"ColorPickerFlagsAlphaBar":                 imgui.ColorPickerFlagsAlphaBar,
	"ColorPickerFlagsAlphaPreview":             imgui.ColorPickerFlagsAlphaPreview,
	"ColorPickerFlagsAlphaPreviewHalf":         imgui.ColorPickerFlagsAlphaPreviewHalf,
	"ColorPickerFlagsFloat":                    imgui.ColorPickerFlagsFloat,
	"ColorPickerFlagsHEX":                      imgui.ColorPickerFlagsHEX,
	"ColorPickerFlagsHSV":                      imgui.ColorPickerFlagsHSV,
	"ColorPickerFlagsNoAlpha":                  imgui.ColorPickerFlagsNoAlpha,
	"ColorPickerFlagsNoInputs":                 imgui.ColorPickerFlagsNoInputs,
	"ColorPickerFlagsNoLabel":                  imgui.ColorPickerFlagsNoLabel,
	"ColorPickerFlagsNoSidePreview":            imgui.ColorPickerFlagsNoSidePreview,
	"ColorPickerFlagsNoSmallPreview":           imgui.ColorPickerFlagsNoSmallPreview,
	"ColorPickerFlagsNoTooltip":                imgui.ColorPickerFlagsNoTooltip,
	"ColorPickerFlagsNone":                     imgui.ColorPickerFlagsNone,
	"ColorPickerFlagsPickerHueBar":             imgui.ColorPickerFlagsPickerHueBar,
	"ColorPickerFlagsPickerHueWheel":           imgui.ColorPickerFlagsPickerHueWheel,
	"ColorPickerFlagsRGB":                      imgui.ColorPickerFlagsRGB,
	"ColorPickerFlagsUint8":                    imgui.ColorPickerFlagsUint8,
	"ComboFlagHeightLarge":                     imgui.ComboFlagHeightLarge,
	"ComboFlagHeightLargest":                   imgui.ComboFlagHeightLargest,
	"ComboFlagHeightRegular":                   imgui.ComboFlagHeightRegular,
	"ComboFlagHeightSmall":                     imgui.ComboFlagHeightSmall,
	"ComboFlagNoArrowButton":                   imgui.ComboFlagNoArrowButton,
	"ComboFlagNoPreview":                       imgui.ComboFlagNoPreview,
	"ComboFlagNone":                            imgui.ComboFlagNone,
	"ComboFlagPopupAlignLeft":                  imgui.ComboFlagPopupAlignLeft,
	"ConditionAlways":                          imgui.ConditionAlways,
	"ConditionAppearing":                       imgui.ConditionAppearing,
	"ConditionFirstUseEver":                    imgui.ConditionFirstUseEver,
	"ConditionOnce":                            imgui.ConditionOnce,
	"ConfigFlagEnablePowerSavingMode":          imgui.ConfigFlagEnablePowerSavingMode,
	"ConfigFlagIsSRGB":                         imgui.ConfigFlagIsSRGB,
	"ConfigFlagIsTouchScreen":                  imgui.ConfigFlagIsTouchScreen,
	"ConfigFlagNavEnableGamepad":               imgui.ConfigFlagNavEnableGamepad,
	"ConfigFlagNavEnableKeyboard":              imgui.ConfigFlagNavEnableKeyboard,
	"ConfigFlagNavEnableSetMousePos":           imgui.ConfigFlagNavEnableSetMousePos,
	"ConfigFlagNavNoCaptureKeyboard":           imgui.ConfigFlagNavNoCaptureKeyboard,
	"ConfigFlagNoMouse":                        imgui.ConfigFlagNoMouse,
	"ConfigFlagNoMouseCursorChange":            imgui.ConfigFlagNoMouseCursorChange,
	"ConfigFlagNone":                           imgui.ConfigFlagNone,
	"DefaultFont":                              imgui.DefaultFont,
	"DefaultFontConfig":                        imgui.DefaultFontConfig,
	"EmptyGlyphRanges":                         imgui.EmptyGlyphRanges,
	"ErrFreeTypeFailed":                        imgui.ErrFreeTypeFailed,
	"ErrFreeTypeNotAvailable":                  imgui.ErrFreeTypeNotAvailable,
	"FreeTypeRasterizerFlagsBold":              imgui.FreeTypeRasterizerFlagsBold,
	"FreeTypeRasterizerFlagsForceAutoHint":     imgui.FreeTypeRasterizerFlagsForceAutoHint,
	"FreeTypeRasterizerFlagsLightHinting":      imgui.FreeTypeRasterizerFlagsLightHinting,
	"FreeTypeRasterizerFlagsMonoHinting":       imgui.FreeTypeRasterizerFlagsMonoHinting,
	"FreeTypeRasterizerFlagsMonochrome":        imgui.FreeTypeRasterizerFlagsMonochrome,
	"FreeTypeRasterizerFlagsNoAutoHint":        imgui.FreeTypeRasterizerFlagsNoAutoHint,
	"FreeTypeRasterizerFlagsNoHinting":         imgui.FreeTypeRasterizerFlagsNoHinting,
	"FreeTypeRasterizerFlagsOblique":           imgui.FreeTypeRasterizerFlagsOblique,
	"GLFWWindowFlagsFloating":                  imgui.GLFWWindowFlagsFloating,
	"GLFWWindowFlagsFrameless":                 imgui.GLFWWindowFlagsFrameless,
	"GLFWWindowFlagsMaximized":                 imgui.GLFWWindowFlagsMaximized,
	"GLFWWindowFlagsNotResizable":              imgui.GLFWWindowFlagsNotResizable,
	"GLFWWindowFlagsTransparent":               imgui.GLFWWindowFlagsTransparent,
	"HoveredFlagsAllowWhenBlockedByActiveItem": imgui.HoveredFlagsAllowWhenBlockedByActiveItem,
	"HoveredFlagsAllowWhenBlockedByPopup":      imgui.HoveredFlagsAllowWhenBlockedByPopup,
	"HoveredFlagsAllowWhenDisabled":            imgui.HoveredFlagsAllowWhenDisabled,
	"HoveredFlagsAllowWhenOverlapped":          imgui.HoveredFlagsAllowWhenOverlapped,
	"HoveredFlagsAnyWindow":                    imgui.HoveredFlagsAnyWindow,
	"HoveredFlagsChildWindows":                 imgui.HoveredFlagsChildWindows,
	"HoveredFlagsNone":                         imgui.HoveredFlagsNone,
	"HoveredFlagsRectOnly":                     imgui.HoveredFlagsRectOnly,
	"HoveredFlagsRootAndChildWindows":          imgui.HoveredFlagsRootAndChildWindows,
	"HoveredFlagsRootWindow":                   imgui.HoveredFlagsRootWindow,
	"InputTextFlagsAllowTabInput":              imgui.InputTextFlagsAllowTabInput,
	"InputTextFlagsAlwaysInsertMode":           imgui.InputTextFlagsAlwaysInsertMode,
	"InputTextFlagsAutoSelectAll":              imgui.InputTextFlagsAutoSelectAll,
	"InputTextFlagsCallbackAlways":             imgui.InputTextFlagsCallbackAlways,
	"InputTextFlagsCallbackCharFilter":         imgui.InputTextFlagsCallbackCharFilter,
	"InputTextFlagsCallbackCompletion":         imgui.InputTextFlagsCallbackCompletion,
	"InputTextFlagsCallbackHistory":            imgui.InputTextFlagsCallbackHistory,
	"InputTextFlagsCharsDecimal":               imgui.InputTextFlagsCharsDecimal,
	"InputTextFlagsCharsHexadecimal":           imgui.InputTextFlagsCharsHexadecimal,
	"InputTextFlagsCharsNoBlank":               imgui.InputTextFlagsCharsNoBlank,
	"InputTextFlagsCharsScientific":            imgui.InputTextFlagsCharsScientific,
	"InputTextFlagsCharsUppercase":             imgui.InputTextFlagsCharsUppercase,
	"InputTextFlagsCtrlEnterForNewLine":        imgui.InputTextFlagsCtrlEnterForNewLine,
	"InputTextFlagsEnterReturnsTrue":           imgui.InputTextFlagsEnterReturnsTrue,
	"InputTextFlagsNoHorizontalScroll":         imgui.InputTextFlagsNoHorizontalScroll,
	"InputTextFlagsNoUndoRedo":                 imgui.InputTextFlagsNoUndoRedo,
	"InputTextFlagsNone":                       imgui.InputTextFlagsNone,
	"InputTextFlagsPassword":                   imgui.InputTextFlagsPassword,
	"InputTextFlagsReadOnly":                   imgui.InputTextFlagsReadOnly,
	"KeyA":                                     imgui.KeyA,
	"KeyBackspace":                             imgui.KeyBackspace,
	"KeyC":                                     imgui.KeyC,
	"KeyDelete":                                imgui.KeyDelete,
	"KeyDownArrow":                             imgui.KeyDownArrow,
	"KeyEnd":                                   imgui.KeyEnd,
	"KeyEnter":                                 imgui.KeyEnter,
	"KeyEscape":                                imgui.KeyEscape,
	"KeyHome":                                  imgui.KeyHome,
	"KeyInsert":                                imgui.KeyInsert,
	"KeyKeyPadEnter":                           imgui.KeyKeyPadEnter,
	"KeyLeftArrow":                             imgui.KeyLeftArrow,
	"KeyPageDown":                              imgui.KeyPageDown,
	"KeyPageUp":                                imgui.KeyPageUp,
	"KeyRightArrow":                            imgui.KeyRightArrow,
	"KeySpace":                                 imgui.KeySpace,
	"KeyTab":                                   imgui.KeyTab,
	"KeyUpArrow":                               imgui.KeyUpArrow,
	"KeyV":                                     imgui.KeyV,
	"KeyX":                                     imgui.KeyX,
	"KeyY":                                     imgui.KeyY,
	"KeyZ":                                     imgui.KeyZ,
	"MouseCursorArrow":                         imgui.MouseCursorArrow,
	"MouseCursorCount":                         imgui.MouseCursorCount,
	"MouseCursorHand":                          imgui.MouseCursorHand,
	"MouseCursorNone":                          imgui.MouseCursorNone,
	"MouseCursorResizeAll":                     imgui.MouseCursorResizeAll,
	"MouseCursorResizeEW":                      imgui.MouseCursorResizeEW,
	"MouseCursorResizeNESW":                    imgui.MouseCursorResizeNESW,
	"MouseCursorResizeNS":                      imgui.MouseCursorResizeNS,
	"MouseCursorResizeNWSE":                    imgui.MouseCursorResizeNWSE,
	"MouseCursorTextInput":                     imgui.MouseCursorTextInput,
	"SelectableFlagsAllowDoubleClick":          imgui.SelectableFlagsAllowDoubleClick,
	"SelectableFlagsDisabled":                  imgui.SelectableFlagsDisabled,
	"SelectableFlagsDontClosePopups":           imgui.SelectableFlagsDontClosePopups,
	"SelectableFlagsNone":                      imgui.SelectableFlagsNone,
	"SelectableFlagsSpanAllColumns":            imgui.SelectableFlagsSpanAllColumns,
	"StyleColorBorder":                         imgui.StyleColorBorder,
	"StyleColorBorderShadow":                   imgui.StyleColorBorderShadow,
	"StyleColorButton":                         imgui.StyleColorButton,
	"StyleColorButtonActive":                   imgui.StyleColorButtonActive,
	"StyleColorButtonHovered":                  imgui.StyleColorButtonHovered,
	"StyleColorCheckMark":                      imgui.StyleColorCheckMark,
	"StyleColorChildBg":                        imgui.StyleColorChildBg,
	"StyleColorDragDropTarget":                 imgui.StyleColorDragDropTarget,
	"StyleColorFrameBg":                        imgui.StyleColorFrameBg,
	"StyleColorFrameBgActive":                  imgui.StyleColorFrameBgActive,
	"StyleColorFrameBgHovered":                 imgui.StyleColorFrameBgHovered,
	"StyleColorHeader":                         imgui.StyleColorHeader,
	"StyleColorHeaderActive":                   imgui.StyleColorHeaderActive,
	"StyleColorHeaderHovered":                  imgui.StyleColorHeaderHovered,
	"StyleColorMenuBarBg":                      imgui.StyleColorMenuBarBg,
	"StyleColorModalWindowDarkening":           imgui.StyleColorModalWindowDarkening,
	"StyleColorNavHighlight":                   imgui.StyleColorNavHighlight,
	"StyleColorNavWindowingDarkening":          imgui.StyleColorNavWindowingDarkening,
	"StyleColorNavWindowingHighlight":          imgui.StyleColorNavWindowingHighlight,
	"StyleColorPlotHistogram":                  imgui.StyleColorPlotHistogram,
	"StyleColorPlotHistogramHovered":           imgui.StyleColorPlotHistogramHovered,
	"StyleColorPlotLines":                      imgui.StyleColorPlotLines,
	"StyleColorPlotLinesHovered":               imgui.StyleColorPlotLinesHovered,
	"StyleColorPopupBg":                        imgui.StyleColorPopupBg,
	"StyleColorResizeGrip":                     imgui.StyleColorResizeGrip,
	"StyleColorResizeGripActive":               imgui.StyleColorResizeGripActive,
	"StyleColorResizeGripHovered":              imgui.StyleColorResizeGripHovered,
	"StyleColorScrollbarBg":                    imgui.StyleColorScrollbarBg,
	"StyleColorScrollbarGrab":                  imgui.StyleColorScrollbarGrab,
	"StyleColorScrollbarGrabActive":            imgui.StyleColorScrollbarGrabActive,
	"StyleColorScrollbarGrabHovered":           imgui.StyleColorScrollbarGrabHovered,
	"StyleColorSeparator":                      imgui.StyleColorSeparator,
	"StyleColorSeparatorActive":                imgui.StyleColorSeparatorActive,
	"StyleColorSeparatorHovered":               imgui.StyleColorSeparatorHovered,
	"StyleColorSliderGrab":                     imgui.StyleColorSliderGrab,
	"StyleColorSliderGrabActive":               imgui.StyleColorSliderGrabActive,
	"StyleColorTab":                            imgui.StyleColorTab,
	"StyleColorTabActive":                      imgui.StyleColorTabActive,
	"StyleColorTabHovered":                     imgui.StyleColorTabHovered,
	"StyleColorTabUnfocused":                   imgui.StyleColorTabUnfocused,
	"StyleColorTabUnfocusedActive":             imgui.StyleColorTabUnfocusedActive,
	"StyleColorText":                           imgui.StyleColorText,
	"StyleColorTextDisabled":                   imgui.StyleColorTextDisabled,
	"StyleColorTextSelectedBg":                 imgui.StyleColorTextSelectedBg,
	"StyleColorTitleBg":                        imgui.StyleColorTitleBg,
	"StyleColorTitleBgActive":                  imgui.StyleColorTitleBgActive,
	"StyleColorTitleBgCollapsed":               imgui.StyleColorTitleBgCollapsed,
	"StyleColorWindowBg":                       imgui.StyleColorWindowBg,
	"StyleVarAlpha":                            imgui.StyleVarAlpha,
	"StyleVarButtonTextAlign":                  imgui.StyleVarButtonTextAlign,
	"StyleVarChildBorderSize":                  imgui.StyleVarChildBorderSize,
	"StyleVarChildRounding":                    imgui.StyleVarChildRounding,
	"StyleVarFrameBorderSize":                  imgui.StyleVarFrameBorderSize,
	"StyleVarFramePadding":                     imgui.StyleVarFramePadding,
	"StyleVarFrameRounding":                    imgui.StyleVarFrameRounding,
	"StyleVarGrabMinSize":                      imgui.StyleVarGrabMinSize,
	"StyleVarGrabRounding":                     imgui.StyleVarGrabRounding,
	"StyleVarIndentSpacing":                    imgui.StyleVarIndentSpacing,
	"StyleVarItemInnerSpacing":                 imgui.StyleVarItemInnerSpacing,
	"StyleVarItemSpacing":                      imgui.StyleVarItemSpacing,
	"StyleVarPopupBorderSize":                  imgui.StyleVarPopupBorderSize,
	"StyleVarPopupRounding":                    imgui.StyleVarPopupRounding,
	"StyleVarScrollbarRounding":                imgui.StyleVarScrollbarRounding,
	"StyleVarScrollbarSize":                    imgui.StyleVarScrollbarSize,
	"StyleVarSelectableTextAlign":              imgui.StyleVarSelectableTextAlign,
	"StyleVarTabRounding":                      imgui.StyleVarTabRounding,
	"StyleVarWindowBorderSize":                 imgui.StyleVarWindowBorderSize,
	"StyleVarWindowMinSize":                    imgui.StyleVarWindowMinSize,
	"StyleVarWindowPadding":                    imgui.StyleVarWindowPadding,
	"StyleVarWindowRounding":                   imgui.StyleVarWindowRounding,
	"StyleVarWindowTitleAlign":                 imgui.StyleVarWindowTitleAlign,
	"TabBarFlagsAutoSelectNewTabs":             imgui.TabBarFlagsAutoSelectNewTabs,
	"TabBarFlagsFittingPolicyDefault":          imgui.TabBarFlagsFittingPolicyDefault,
	"TabBarFlagsFittingPolicyMask":             imgui.TabBarFlagsFittingPolicyMask,
	"TabBarFlagsFittingPolicyResizeDown":       imgui.TabBarFlagsFittingPolicyResizeDown,
	"TabBarFlagsFittingPolicyScroll":           imgui.TabBarFlagsFittingPolicyScroll,
	"TabBarFlagsNoCloseWithMiddleMouseButton":  imgui.TabBarFlagsNoCloseWithMiddleMouseButton,
	"TabBarFlagsNoTabListScrollingButtons":     imgui.TabBarFlagsNoTabListScrollingButtons,
	"TabBarFlagsNoTooltip":                     imgui.TabBarFlagsNoTooltip,
	"TabBarFlagsNone":                          imgui.TabBarFlagsNone,
	"TabBarFlagsReorderable":                   imgui.TabBarFlagsReorderable,
	"TabBarFlagsTabListPopupButton":            imgui.TabBarFlagsTabListPopupButton,
	"TabItemFlagsNoCloseWithMiddleMouseButton": imgui.TabItemFlagsNoCloseWithMiddleMouseButton,
	"TabItemFlagsNoPushID":                     imgui.TabItemFlagsNoPushID,
	"TabItemFlagsNone":                         imgui.TabItemFlagsNone,
	"TabItemFlagsSetSelected":                  imgui.TabItemFlagsSetSelected,
	"TabItemFlagsUnsavedDocument":              imgui.TabItemFlagsUnsavedDocument,
	"TreeNodeFlagsAllowItemOverlap":            imgui.TreeNodeFlagsAllowItemOverlap,
	"TreeNodeFlagsBullet":                      imgui.TreeNodeFlagsBullet,
	"TreeNodeFlagsCollapsingHeader":            imgui.TreeNodeFlagsCollapsingHeader,
	"TreeNodeFlagsDefaultOpen":                 imgui.TreeNodeFlagsDefaultOpen,
	"TreeNodeFlagsFramePadding":                imgui.TreeNodeFlagsFramePadding,
	"TreeNodeFlagsFramed":                      imgui.TreeNodeFlagsFramed,
	"TreeNodeFlagsLeaf":                        imgui.TreeNodeFlagsLeaf,
	"TreeNodeFlagsNavLeftJumpsBackHere":        imgui.TreeNodeFlagsNavLeftJumpsBackHere,
	"TreeNodeFlagsNoAutoOpenOnLog":             imgui.TreeNodeFlagsNoAutoOpenOnLog,
	"TreeNodeFlagsNoTreePushOnOpen":            imgui.TreeNodeFlagsNoTreePushOnOpen,
	"TreeNodeFlagsNone":                        imgui.TreeNodeFlagsNone,
	"TreeNodeFlagsOpenOnArrow":                 imgui.TreeNodeFlagsOpenOnArrow,
	"TreeNodeFlagsOpenOnDoubleClick":           imgui.TreeNodeFlagsOpenOnDoubleClick,
	"TreeNodeFlagsSelected":                    imgui.TreeNodeFlagsSelected,
	"TreeNodeFlagsSpanAvailWidth":              imgui.TreeNodeFlagsSpanAvailWidth,
	"TreeNodeFlagsSpanFullWidth":               imgui.TreeNodeFlagsSpanFullWidth,
	"WindowFlagsAlwaysAutoResize":              imgui.WindowFlagsAlwaysAutoResize,
	"WindowFlagsAlwaysHorizontalScrollbar":     imgui.WindowFlagsAlwaysHorizontalScrollbar,
	"WindowFlagsAlwaysUseWindowPadding":        imgui.WindowFlagsAlwaysUseWindowPadding,
	"WindowFlagsAlwaysVerticalScrollbar":       imgui.WindowFlagsAlwaysVerticalScrollbar,
	"WindowFlagsHorizontalScrollbar":           imgui.WindowFlagsHorizontalScrollbar,
	"WindowFlagsMenuBar":                       imgui.WindowFlagsMenuBar,
	"WindowFlagsNoBackground":                  imgui.WindowFlagsNoBackground,
	"WindowFlagsNoBringToFrontOnFocus":         imgui.WindowFlagsNoBringToFrontOnFocus,
	"WindowFlagsNoCollapse":                    imgui.WindowFlagsNoCollapse,
	"WindowFlagsNoDecoration":                  imgui.WindowFlagsNoDecoration,
	"WindowFlagsNoFocusOnAppearing":            imgui.WindowFlagsNoFocusOnAppearing,
	"WindowFlagsNoInputs":                      imgui.WindowFlagsNoInputs,
	"WindowFlagsNoMouseInputs":                 imgui.WindowFlagsNoMouseInputs,
	"WindowFlagsNoMove":                        imgui.WindowFlagsNoMove,
	"WindowFlagsNoNav":                         imgui.WindowFlagsNoNav,
	"WindowFlagsNoNavFocus":                    imgui.WindowFlagsNoNavFocus,
	"WindowFlagsNoNavInputs":                   imgui.WindowFlagsNoNavInputs,
	"WindowFlagsNoResize":                      imgui.WindowFlagsNoResize,
	"WindowFlagsNoSavedSettings":               imgui.WindowFlagsNoSavedSettings,
	"WindowFlagsNoScrollWithMouse":             imgui.WindowFlagsNoScrollWithMouse,
	"WindowFlagsNoScrollbar":                   imgui.WindowFlagsNoScrollbar,
	"WindowFlagsNoTitleBar":                    imgui.WindowFlagsNoTitleBar,
	"WindowFlagsNone":                          imgui.WindowFlagsNone,
	"WindowFlagsUnsavedDocument":               imgui.WindowFlagsUnsavedDocument,

	"DPIScale":            imgui.DPIScale,
	"EnableFreeType":      imgui.EnableFreeType,
	"ErrContextDestroyed": imgui.ErrContextDestroyed,
	"ErrNoContext":        imgui.ErrNoContext,

	"AlignTextToFramePadding":   imgui.AlignTextToFramePadding,
	"Begin":                     imgui.Begin,
	"BeginChild":                imgui.BeginChild,
	"BeginChildV":               imgui.BeginChildV,
	"BeginCombo":                imgui.BeginCombo,
	"BeginComboV":               imgui.BeginComboV,
	"BeginDragDropSource":       imgui.BeginDragDropSource,
	"BeginDragDropSourceV":      imgui.BeginDragDropSourceV,
	"BeginDragDropTarget":       imgui.BeginDragDropTarget,
	"BeginGroup":                imgui.BeginGroup,
	"BeginMainMenuBar":          imgui.BeginMainMenuBar,
	"BeginMenu":                 imgui.BeginMenu,
	"BeginMenuBar":              imgui.BeginMenuBar,
	"BeginMenuV":                imgui.BeginMenuV,
	"BeginPopupContextItem":     imgui.BeginPopupContextItem,
	"BeginPopupContextItemV":    imgui.BeginPopupContextItemV,
	"BeginPopupModal":           imgui.BeginPopupModal,
	"BeginPopupModalV":          imgui.BeginPopupModalV,
	"BeginTabBar":               imgui.BeginTabBar,
	"BeginTabBarV":              imgui.BeginTabBarV,
	"BeginTabItem":              imgui.BeginTabItem,
	"BeginTabItemV":             imgui.BeginTabItemV,
	"BeginTooltip":              imgui.BeginTooltip,
	"BeginV":                    imgui.BeginV,
	"Button":                    imgui.Button,
	"ButtonV":                   imgui.ButtonV,
	"CalcItemWidth":             imgui.CalcItemWidth,
	"Checkbox":                  imgui.Checkbox,
	"CloseCurrentPopup":         imgui.CloseCurrentPopup,
	"ColorEdit3":                imgui.ColorEdit3,
	"ColorEdit3V":               imgui.ColorEdit3V,
	"ColorEdit4":                imgui.ColorEdit4,
	"ColorEdit4V":               imgui.ColorEdit4V,
	"ColorPicker3":              imgui.ColorPicker3,
	"ColorPicker3V":             imgui.ColorPicker3V,
	"ColorPicker4":              imgui.ColorPicker4,
	"ColorPicker4V":             imgui.ColorPicker4V,
	"ColumnIndex":               imgui.ColumnIndex,
	"ColumnOffset":              imgui.ColumnOffset,
	"ColumnOffsetV":             imgui.ColumnOffsetV,
	"ColumnWidth":               imgui.ColumnWidth,
	"ColumnWidthV":              imgui.ColumnWidthV,
	"Columns":                   imgui.Columns,
	"ColumnsCount":              imgui.ColumnsCount,
	"ColumnsV":                  imgui.ColumnsV,
	"CursorPosX":                imgui.CursorPosX,
	"CursorPosY":                imgui.CursorPosY,
	"DragFloat":                 imgui.DragFloat,
	"DragFloatV":                imgui.DragFloatV,
	"DragInt":                   imgui.DragInt,
	"DragIntV":                  imgui.DragIntV,
	"Dummy":                     imgui.Dummy,
	"End":                       imgui.End,
	"EndChild":                  imgui.EndChild,
	"EndCombo":                  imgui.EndCombo,
	"EndDragDropSource":         imgui.EndDragDropSource,
	"EndDragDropTarget":         imgui.EndDragDropTarget,
	"EndFrame":                  imgui.EndFrame,
	"EndGroup":                  imgui.EndGroup,
	"EndMainMenuBar":            imgui.EndMainMenuBar,
	"EndMenu":                   imgui.EndMenu,
	"EndMenuBar":                imgui.EndMenuBar,
	"EndPopup":                  imgui.EndPopup,
	"EndTabBar":                 imgui.EndTabBar,
	"EndTabItem":                imgui.EndTabItem,
	"EndTooltip":                imgui.EndTooltip,
	"FontSize":                  imgui.FontSize,
	"GetColorU32":               imgui.GetColorU32,
	"GetEventWaitingTime":       imgui.GetEventWaitingTime,
	"Image":                     imgui.Image,
	"ImageButton":               imgui.ImageButton,
	"ImageButtonV":              imgui.ImageButtonV,
	"ImageV":                    imgui.ImageV,
	"IndexBufferLayout":         imgui.IndexBufferLayout,
	"InputFloat":                imgui.InputFloat,
	"InputFloatV":               imgui.InputFloatV,
	"InputInt":                  imgui.InputInt,
	"InputIntV":                 imgui.InputIntV,
	"InputText":                 imgui.InputText,
	"InputTextMultiline":        imgui.InputTextMultiline,
	"InputTextMultilineV":       imgui.InputTextMultilineV,
	"InputTextV":                imgui.InputTextV,
	"InvisibleButton":           imgui.InvisibleButton,
	"IsAnyItemActive":           imgui.IsAnyItemActive,
	"IsAnyItemFocused":          imgui.IsAnyItemFocused,
	"IsAnyMouseDown":            imgui.IsAnyMouseDown,
	"IsItemActive":              imgui.IsItemActive,
	"IsItemFocused":             imgui.IsItemFocused,
	"IsItemHovered":             imgui.IsItemHovered,
	"IsItemHoveredV":            imgui.IsItemHoveredV,
	"IsKeyDown":                 imgui.IsKeyDown,
	"IsKeyPressed":              imgui.IsKeyPressed,
	"IsKeyPressedV":             imgui.IsKeyPressedV,
	"IsKeyReleased":             imgui.IsKeyReleased,
	"IsMouseClicked":            imgui.IsMouseClicked,
	"IsMouseClickedV":           imgui.IsMouseClickedV,
	"IsMouseDoubleClicked":      imgui.IsMouseDoubleClicked,
	"IsMouseDown":               imgui.IsMouseDown,
	"IsMouseReleased":           imgui.IsMouseReleased,
	"IsWindowAppearing":         imgui.IsWindowAppearing,
	"LabelText":                 imgui.LabelText,
	"ListBox":                   imgui.ListBox,
	"ListBoxV":                  imgui.ListBoxV,
	"MenuItem":                  imgui.MenuItem,
	"MenuItemV":                 imgui.MenuItemV,
	"MouseCursor":               imgui.MouseCursor,
	"NewFrame":                  imgui.NewFrame,
	"NextColumn":                imgui.NextColumn,
	"OpenPopup":                 imgui.OpenPopup,
	"PlotHistogram":             imgui.PlotHistogram,
	"PlotHistogramV":            imgui.PlotHistogramV,
	"PlotLines":                 imgui.PlotLines,
	"PlotLinesV":                imgui.PlotLinesV,
	"PopFont":                   imgui.PopFont,
	"PopID":                     imgui.PopID,
	"PopItemWidth":              imgui.PopItemWidth,
	"PopStyleColor":             imgui.PopStyleColor,
	"PopStyleColorV":            imgui.PopStyleColorV,
	"PopStyleVar":               imgui.PopStyleVar,
	"PopStyleVarV":              imgui.PopStyleVarV,
	"PopTextWrapPos":            imgui.PopTextWrapPos,
	"ProgressBar":               imgui.ProgressBar,
	"ProgressBarV":              imgui.ProgressBarV,
	"PushFont":                  imgui.PushFont,
	"PushID":                    imgui.PushID,
	"PushItemWidth":             imgui.PushItemWidth,
	"PushStyleColor":            imgui.PushStyleColor,
	"PushStyleVarFloat":         imgui.PushStyleVarFloat,
	"PushStyleVarVec2":          imgui.PushStyleVarVec2,
	"PushTextWrapPos":           imgui.PushTextWrapPos,
	"PushTextWrapPosV":          imgui.PushTextWrapPosV,
	"RadioButton":               imgui.RadioButton,
	"Render":                    imgui.Render,
	"SameLine":                  imgui.SameLine,
	"SameLineV":                 imgui.SameLineV,
	"Selectable":                imgui.Selectable,
	"SelectableV":               imgui.SelectableV,
	"Separator":                 imgui.Separator,
	"SetAssertHandler":          imgui.SetAssertHandler,
	"SetColumnOffset":           imgui.SetColumnOffset,
	"SetColumnWidth":            imgui.SetColumnWidth,
	"SetCursorPos":              imgui.SetCursorPos,
	"SetCursorScreenPos":        imgui.SetCursorScreenPos,
	"SetDragDropPayload":        imgui.SetDragDropPayload,
	"SetDragDropPayloadV":       imgui.SetDragDropPayloadV,
	"SetItemDefaultFocus":       imgui.SetItemDefaultFocus,
	"SetKeyboardFocusHere":      imgui.SetKeyboardFocusHere,
	"SetKeyboardFocusHereV":     imgui.SetKeyboardFocusHereV,
	"SetMaxWaitBeforeNextFrame": imgui.SetMaxWaitBeforeNextFrame,
	"SetMouseCursor":            imgui.SetMouseCursor,
	"SetNextItemOpen":           imgui.SetNextItemOpen,
	"SetNextWindowBgAlpha":      imgui.SetNextWindowBgAlpha,
	"SetNextWindowContentSize":  imgui.SetNextWindowContentSize,
	"SetNextWindowFocus":        imgui.SetNextWindowFocus,
	"SetNextWindowPos":          imgui.SetNextWindowPos,
	"SetNextWindowPosV":         imgui.SetNextWindowPosV,
	"SetNextWindowSize":         imgui.SetNextWindowSize,
	"SetNextWindowSizeV":        imgui.SetNextWindowSizeV,
	"SetScrollHereY":            imgui.SetScrollHereY,
	"SetTabItemClosed":          imgui.SetTabItemClosed,
	"SetTooltip":                imgui.SetTooltip,
	"ShowDemoWindow":            imgui.ShowDemoWindow,
	"ShowUserGuide":             imgui.ShowUserGuide,
	"SliderFloat":               imgui.SliderFloat,
	"SliderFloat3":              imgui.SliderFloat3,
	"SliderFloat3V":             imgui.SliderFloat3V,
	"SliderFloatV":              imgui.SliderFloatV,
	"SliderInt":                 imgui.SliderInt,
	"SliderIntV":                imgui.SliderIntV,
	"Spacing":                   imgui.Spacing,
	"StyleColorsClassic":        imgui.StyleColorsClassic,
	"StyleColorsDark":           imgui.StyleColorsDark,
	"StyleColorsLight":          imgui.StyleColorsLight,
	"Text":                      imgui.Text,
	"TextLineHeight":            imgui.TextLineHeight,
	"TextLineHeightWithSpacing": imgui.TextLineHeightWithSpacing,
	"TreeNode":                  imgui.TreeNode,
	"TreeNodeToLabelSpacing":    imgui.TreeNodeToLabelSpacing,
	"TreeNodeV":                 imgui.TreeNodeV,
	"TreePop":                   imgui.TreePop,
	"Version":                   imgui.Version,
	"VertexBufferLayout":        imgui.VertexBufferLayout,
	"WindowHeight":              imgui.WindowHeight,
	"WindowWidth":               imgui.WindowWidth,

	"Alpha8Image":        qlang.StructOf((*imgui.Alpha8Image)(nil)),
	"CreateContext":      imgui.CreateContext,
	"CurrentContext":     imgui.CurrentContext,
	"glfw":               imgui.NewGLFW,
	"glfwclipboard":      imgui.NewGLFWClipboard,
	"CurrentIO":          imgui.CurrentIO,
	"ListClipper":        qlang.StructOf((*imgui.ListClipper)(nil)),
	"opengl3":            imgui.NewOpenGL3,
	"RGBA32Image":        qlang.StructOf((*imgui.RGBA32Image)(nil)),
	"Vec2":               qlang.StructOf((*imgui.Vec2)(nil)),
	"CalcTextSize":       imgui.CalcTextSize,
	"ContentRegionAvail": imgui.ContentRegionAvail,
	"CursorPos":          imgui.CursorPos,
	"CursorScreenPos":    imgui.CursorScreenPos,
	"CursorStartPos":     imgui.CursorStartPos,
	"GetItemRectMax":     imgui.GetItemRectMax,
	"GetItemRectMin":     imgui.GetItemRectMin,
	"GetItemRectSize":    imgui.GetItemRectSize,
	"WindowPos":          imgui.WindowPos,
	"WindowSize":         imgui.WindowSize,
	"Vec4":               qlang.StructOf((*imgui.Vec4)(nil)),
}
