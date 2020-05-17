package sort

import (
	"sort"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "sort",

	"Float64s":          sort.Float64s,
	"Float64sAreSorted": sort.Float64sAreSorted,
	"Ints":              sort.Ints,
	"IntsAreSorted":     sort.IntsAreSorted,
	"IsSorted":          sort.IsSorted,
	"Search":            sort.Search,
	"SearchFloat64s":    sort.SearchFloat64s,
	"SearchInts":        sort.SearchInts,
	"SearchStrings":     sort.SearchStrings,
	"Slice":             sort.Slice,
	"SliceIsSorted":     sort.SliceIsSorted,
	"SliceStable":       sort.SliceStable,
	"Sort":              sort.Sort,
	"Stable":            sort.Stable,
	"strings":           sort.Strings,
	"Strings":           sort.Strings,
	"StringsAreSorted":  sort.StringsAreSorted,

	"Reverse": sort.Reverse,
}
