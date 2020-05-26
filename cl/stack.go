package qlang

import (
	"strconv"

	"github.com/topxeq/qlang/exec"
)

// -----------------------------------------------------------------------------

func (p *Compiler) pushInt(v int) {
	// fmt.Printf("pushInt: %v\n", v)

	p.code.Block(exec.Push(v))
}

func (p *Compiler) pushFloat(v float64) {
	// fmt.Printf("pushFloat: %v\n", v)

	p.code.Block(exec.Push(v))
}

func (p *Compiler) pushByte(lit string) {
	// fmt.Printf("pushbyte: %v\n", lit)

	v, multibyte, tail, err := strconv.UnquoteChar(lit[1:len(lit)-1], '\'')
	if err != nil {
		panic("invalid char `" + lit + "`: " + err.Error())
	}
	if tail != "" || multibyte {
		panic("invalid char: " + lit)
	}
	p.code.Block(exec.Push(byte(v)))
}

func (p *Compiler) pushString(lit string) {
	// fmt.Printf("pushString: %v\n", lit)

	v, err := strconv.Unquote(lit)
	if err != nil {
		panic("invalid string `" + lit + "`: " + err.Error())
	}
	p.code.Block(exec.Push(v))
}

func (p *Compiler) pushID(name string) {
	// fmt.Printf("push id: %v\n", name)

	p.code.Block(exec.Push(name))
}

// -----------------------------------------------------------------------------

func (p *Compiler) popArity() int {

	if v, ok := p.gstk.Pop(); ok {
		if arity, ok := v.(int); ok {
			// fmt.Printf("poparity: %v\n", arity)
			return arity
		}
	}
	panic("no arity")
}

func (p *Compiler) popName() string {

	if v, ok := p.gstk.Pop(); ok {
		if name, ok := v.(string); ok {
			// fmt.Printf("pop name: %v\n", name)
			return name
		}
	}
	panic("no ident name")
}

func (p *Compiler) pushCode(code interface{}) {
	// fmt.Printf("pushCode: %v\n", code)

	p.gstk.Push(code)
}

func (p *Compiler) arity(arity int) {
	// fmt.Printf("pusharity: %v\n", arity)

	p.gstk.Push(arity)
}

func (p *Compiler) pushName(name string) {
	// fmt.Printf("push name: %v\n", name)
	p.gstk.Push(name)
}

// -----------------------------------------------------------------------------
