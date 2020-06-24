package doc2vec

import (
	"github.com/topxeq/doc2vec/doc2vec"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/topxeq/doc2vec/doc2vec",

	"EXP_TABLE_SIZE":          doc2vec.EXP_TABLE_SIZE,
	"MAX_EXP":                 doc2vec.MAX_EXP,
	"NEG_SAMPLING_TABLE_SIZE": doc2vec.NEG_SAMPLING_TABLE_SIZE,
	"PROGRESS_BAR_THRESHOLD":  doc2vec.PROGRESS_BAR_THRESHOLD,
	"THREAD_NUM":              doc2vec.THREAD_NUM,

	"DBC2SBC":                    doc2vec.DBC2SBC,
	"GetNegativeSamplingWordIdx": doc2vec.GetNegativeSamplingWordIdx,
	"GetSigmoidValue":            doc2vec.GetSigmoidValue,
	"If":                         doc2vec.If,
	"Max":                        doc2vec.Max,
	"Min":                        doc2vec.Min,
	"QuickSort":                  doc2vec.QuickSort,
	"SBC2DBC":                    doc2vec.SBC2DBC,

	"NewDoc2Vec": doc2vec.NewDoc2Vec,

	"SortItem":     qlang.StructOf((*doc2vec.SortItem)(nil)),
	"TDoc2VecImpl": qlang.StructOf((*doc2vec.TDoc2VecImpl)(nil)),
}
