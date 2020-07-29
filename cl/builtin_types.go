package qlang

import (
	"github.com/topxeq/qlang/exec"
	qlang "github.com/topxeq/qlang/spec"
)

// -----------------------------------------------------------------------------

func (p *Compiler) structInit() {

	arity := p.popArity()
	p.Code.Block(exec.StructInit((arity << 1) + 1))
}

func (p *Compiler) mapInit() {

	arity := p.popArity()
	p.Code.Block(exec.MapInit((arity << 1) + 1))
}

// -----------------------------------------------------------------------------

func (p *Compiler) tMap() {

	p.Code.Block(exec.Map)
}

func (p *Compiler) vMap() {

	arity := p.popArity()
	p.Code.Block(exec.Call(qlang.MapFrom, arity<<1))
}

// -----------------------------------------------------------------------------

func (p *Compiler) tSlice() {

	p.Code.Block(exec.Slice)
}

func (p *Compiler) vSlice() {

	hasSlice := p.popArity()
	hasInit := 0
	arityInit := 0
	if hasSlice > 0 {
		hasInit = p.popArity()
		if hasInit > 0 {
			arityInit = p.popArity()
		}
	}
	arity := p.popArity()

	if hasSlice > 0 {
		if arity > 0 {
			panic("must be []type")
		}
		if hasInit > 0 { // []T{a1, a2, ...}
			p.Code.Block(exec.SliceFromTy(arityInit + 1))
		} else { // []T
			p.Code.Block(exec.Slice)
		}
	} else { // [a1, a2, ...]
		p.Code.Block(exec.SliceFrom(arity))
	}
}

// -----------------------------------------------------------------------------
