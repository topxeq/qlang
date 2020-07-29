package qlang

import (
	"github.com/topxeq/qlang/exec"
	"github.com/topxeq/text/tpl/interpreter.util"
)

// -----------------------------------------------------------------------------

func (p *Compiler) fnGo(e interpreter.Engine) {

	src, _ := p.gstk.Pop()
	instr := p.Code.Reserve()
	fnctx := p.Fnctx
	p.exits = append(p.exits, func() {
		start, end := p.clBlock(e, "expr", src, fnctx)
		instr.Set(exec.Go(start, end))
	})
}

// -----------------------------------------------------------------------------

func (p *Compiler) chanIn() {

	p.Code.Block(exec.ChanIn)
}

func (p *Compiler) chanOut() {

	p.Code.Block(exec.ChanOut)
}

func (p *Compiler) tChan() {

	p.Code.Block(exec.Chan)
}

// -----------------------------------------------------------------------------
