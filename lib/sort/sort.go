package sort

import (
	"sort"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "sort",

	"float64s":          sort.Float64s,
	"float64sAreSorted": sort.Float64sAreSorted,
	"ints":              sort.Ints,
	"intsAreSorted":     sort.IntsAreSorted,
	"isSorted":          sort.IsSorted,
	"search":            sort.Search,
	"searchFloat64s":    sort.SearchFloat64s,
	"searchInts":        sort.SearchInts,
	"searchStrings":     sort.SearchStrings,
	"slice":             sort.Slice,
	"sliceIsSorted":     sort.SliceIsSorted,
	"sliceStable":       sort.SliceStable,
	"sort":              sort.Sort,
	"stable":            sort.Stable,
	"strings":           sort.Strings,
	"Strings":           sort.Strings,
	"stringsAreSorted":  sort.StringsAreSorted,

	"reverse": sort.Reverse,
}
