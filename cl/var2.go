package qlang

import (
	"github.com/topxeq/qlang/exec"
	qlang "github.com/topxeq/qlang/spec"
)

// -----------------------------------------------------------------------------

func (p *Compiler) inc() {

	p.Code.Block(exec.IncEx)
}

func (p *Compiler) dec() {

	p.Code.Block(exec.DecEx)
}

func (p *Compiler) addAssign() {

	p.Code.Block(exec.AddAssignEx)
}

func (p *Compiler) subAssign() {

	p.Code.Block(exec.SubAssignEx)
}

func (p *Compiler) mulAssign() {

	p.Code.Block(exec.MulAssignEx)
}

func (p *Compiler) quoAssign() {

	p.Code.Block(exec.QuoAssignEx)
}

func (p *Compiler) modAssign() {

	p.Code.Block(exec.ModAssignEx)
}

func (p *Compiler) xorAssign() {

	p.Code.Block(exec.XorAssignEx)
}

func (p *Compiler) bitandAssign() {

	p.Code.Block(exec.BitAndAssignEx)
}

func (p *Compiler) bitorAssign() {

	p.Code.Block(exec.BitOrAssignEx)
}

func (p *Compiler) andnotAssign() {

	p.Code.Block(exec.AndNotAssignEx)
}

func (p *Compiler) lshrAssign() {

	p.Code.Block(exec.LshrAssignEx)
}

func (p *Compiler) rshrAssign() {

	p.Code.Block(exec.RshrAssignEx)
}

func (p *Compiler) multiAssign() {

	nval := p.popArity()
	arity := p.popArity() + 1
	if nval == 1 {
		p.Code.Block(exec.MultiAssignFromSliceEx(arity))
	} else if arity != nval {
		panic("argument count of multi assignment doesn't match")
	} else {
		p.Code.Block(exec.MultiAssignEx(arity))
	}
}

func (p *Compiler) assign() {

	p.Code.Block(exec.AssignEx)
}

func (p *FuncCtx) requireSymbol(name string) (id int, fnew bool) {
	id, ok := p.getSymbol(name)
	if ok {
		return
	}
	return p.newSymbol(name), true
}

func (p *Compiler) ref(name string) {

	if val, ok := p.gvars[name]; ok {
		p.Code.Block(exec.Push(val))
		return
	}

	fnctx := p.Fnctx
	id, ok := fnctx.getSymbol(name)
	if !ok {
		if val, ok := qlang.Fntable[name]; ok {
			p.Code.Block(exec.GfnRef(val, func() exec.Instr {
				id := fnctx.newSymbol(name)
				return exec.Var(id)
			}))
			return
		}
		id = fnctx.newSymbol(name)
	}
	p.Code.Block(exec.Ref(id))
}

func (p *Compiler) index() {

	arity2 := p.popArity()
	arityMid := p.popArity()
	arity1 := p.popArity()

	if arityMid == 0 {
		if arity1 == 0 {
			panic("call operator[] without index")
		}
		p.Code.Block(exec.Get)
	} else {
		p.Code.Block(exec.Op3(qlang.SubSlice, arity1 != 0, arity2 != 0))
	}
}

func (p *Compiler) toVar() {

	p.Code.ToVar()
}

func (p *Compiler) addrof(name string) {
	// tty.Debugf("Compiler.addrof(%s)\n", name)
	if _, ok := p.gvars[name]; ok {
		panic("only C variable could be addressable")
		return
	}
	fnctx := p.Fnctx
	id, ok := fnctx.getSymbol(name)
	if !ok {
		panic("variable %s not defined yet")
	}
	p.Code.Block(exec.AddrOf(id))
}

// -----------------------------------------------------------------------------
