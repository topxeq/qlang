package utils

import (
	"github.com/topxeq/gods/utils"
	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/gods/utils",

	"ByteComparator":    utils.ByteComparator2,
	"Float32Comparator": utils.Float32Comparator2,
	"Float64Comparator": utils.Float64Comparator2,
	"Int16Comparator":   utils.Int16Comparator2,
	"Int32Comparator":   utils.Int32Comparator2,
	"Int64Comparator":   utils.Int64Comparator2,
	"Int8Comparator":    utils.Int8Comparator2,
	"IntComparator":     utils.IntComparator2,
	"RuneComparator":    utils.RuneComparator2,
	"Sort":              utils.Sort,
	"StringComparator":  utils.StringComparator2,
	"StringComparator2": utils.StringComparator2,
	"TimeComparator":    utils.TimeComparator2,
	"ToString":          utils.ToString,
	"UInt16Comparator":  utils.UInt16Comparator2,
	"UInt32Comparator":  utils.UInt32Comparator2,
	"UInt64Comparator":  utils.UInt64Comparator2,
	"UInt8Comparator":   utils.UInt8Comparator2,
	"UIntComparator":    utils.UIntComparator2,

	"Comparator": qlang.StructOf((*utils.Comparator)(nil)),
}
