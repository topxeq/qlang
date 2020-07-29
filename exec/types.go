package exec

import (
	qlang "github.com/topxeq/qlang/spec"
)

// -----------------------------------------------------------------------------

type tchan int

func (p tchan) OptimizableGetArity() int {

	return 1
}

func (p tchan) Exec(stk *Stack, ctx *Context) {

	n := len(stk.Data) - 1
	stk.Data[n] = qlang.ChanOf(stk.Data[n])
}

// Chan is an instruction that returns chan T.
//
var Chan Instr = tchan(0)

// -----------------------------------------------------------------------------

type slice int

func (p slice) OptimizableGetArity() int {

	return 1
}

func (p slice) Exec(stk *Stack, ctx *Context) {

	n := len(stk.Data) - 1
	stk.Data[n] = qlang.Slice(stk.Data[n])
}

// Slice is an instruction that returns []T.
//
var Slice Instr = slice(0)

// -----------------------------------------------------------------------------

type sliceFrom int
type sliceFromTy int

var nilVarSlice = make([]interface{}, 0)

func (p sliceFrom) Exec(stk *Stack, ctx *Context) {

	if p == 0 {
		stk.Data = append(stk.Data, nilVarSlice)
		return
	}
	n := len(stk.Data) - int(p)
	stk.Data[n] = qlang.SliceFrom(stk.Data[n:]...)
	stk.Data = stk.Data[:n+1]
}

func (p sliceFromTy) Exec(stk *Stack, ctx *Context) {

	n := len(stk.Data) - int(p)
	stk.Data[n] = qlang.SliceFromTy(stk.Data[n:]...)
	stk.Data = stk.Data[:n+1]
}

// SliceFrom is an instruction that creates slice in [a1, a2, ...] form.
//
func SliceFrom(arity int) Instr {

	return sliceFrom(arity)
}

// SliceFromTy is an instruction that creates slice in []T{a1, a2, ...} form.
//
func SliceFromTy(arity int) Instr {

	return sliceFromTy(arity)
}

// -----------------------------------------------------------------------------
// StructInit

type iStructInit int

func (p iStructInit) Exec(stk *Stack, ctx *Context) {

	n := len(stk.Data) - int(p)
	stk.Data[n] = qlang.StructInit(stk.Data[n:]...)
	stk.Data = stk.Data[:n+1]
}

// StructInit returns a StructInit instruction that means `&StructType{name1: expr1, name2: expr2, ...}`.
//
func StructInit(arity int) Instr {
	return iStructInit(arity)
}

// -----------------------------------------------------------------------------
// MapInit

type iMapInit int

func (p iMapInit) Exec(stk *Stack, ctx *Context) {

	n := len(stk.Data) - int(p)
	stk.Data[n] = qlang.MapInit(stk.Data[n:]...)
	stk.Data = stk.Data[:n+1]
}

// MapInit returns a MapInit instruction that means `map[key]elem{key1: expr1, key2: expr2, ...}`.
//
func MapInit(arity int) Instr {
	return iMapInit(arity)
}

// -----------------------------------------------------------------------------

type tmap int

func (p tmap) OptimizableGetArity() int {

	return 2
}

func (p tmap) Exec(stk *Stack, ctx *Context) {

	n := len(stk.Data) - 1
	stk.Data[n-1] = qlang.Map(stk.Data[n-1], stk.Data[n])
	stk.Data = stk.Data[:n]
}

// Map is an instruction that returns `map[key]elem`.
//
var Map Instr = tmap(0)

// -----------------------------------------------------------------------------
