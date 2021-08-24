package exec

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime/debug"
	"sort"
	"strings"

	qlang "github.com/topxeq/qlang/spec"
	"github.com/topxeq/tk"
)

// -----------------------------------------------------------------------------
// type Stack

// A Stack represents a FILO container.
//
type Stack struct {
	Data []interface{}
}

// NewStack returns a new Stack.
//
func NewStack() *Stack {

	data := make([]interface{}, 0, 16)
	return &Stack{data}
}

func (v Stack) String() string {

	bufT := make([]string, 0, len(v.Data))
	for i, iv := range v.Data {
		bufT = append(bufT, fmt.Sprintf("%v: (%T) %v", i, iv, iv))
	}

	return "[" + strings.Join(bufT, ", ") + "]"
}

// Push pushs a value into this stack.
//
func (p *Stack) Push(v interface{}) {

	p.Data = append(p.Data, v)
}

// Top returns the last pushed value, if it exists.
//
func (p *Stack) Top() (v interface{}, ok bool) {

	n := len(p.Data)
	if n > 0 {
		v, ok = p.Data[n-1], true
	}
	return
}

// Pop pops a value from this stack.
//
func (p *Stack) Pop() (v interface{}, ok bool) {

	n := len(p.Data)
	if n > 0 {
		v, ok = p.Data[n-1], true
		p.Data = p.Data[:n-1]
	}
	return
}

// PushRet pushs a function call result.
//
func (p *Stack) PushRet(ret []reflect.Value) error {

	switch len(ret) {
	case 0:
		p.Push(nil)
	case 1:
		p.Push(ret[0].Interface())
	default:
		slice := make([]interface{}, len(ret))
		for i, v := range ret {
			slice[i] = v.Interface()
		}
		p.Push(slice)
	}
	return nil
}

// PopArgs pops arguments of a function call.
//
func (p *Stack) PopArgs(arity int) (args []reflect.Value, ok bool) {

	pstk := p.Data
	n := len(pstk)
	if n >= arity {
		args, ok = make([]reflect.Value, arity), true
		n -= arity
		for i := 0; i < arity; i++ {
			args[i] = reflect.ValueOf(pstk[n+i])
		}
		p.Data = pstk[:n]
	}
	return
}

// PopNArgs pops arguments of a function call.
//
func (p *Stack) PopNArgs(arity int) []interface{} {

	pstk := p.Data
	n := len(pstk)
	if n >= arity {
		args := make([]interface{}, arity)
		n -= arity
		for i := 0; i < arity; i++ {
			args[i] = pstk[n+i]
		}
		p.Data = pstk[:n]
		return args
	}
	panic("PopNArgs: unexpected argument count")
}

// PopFnArgs pops argument names of a function call.
//
func (p *Stack) PopFnArgs(arity int) []string {

	ok := false
	pstk := p.Data
	n := len(pstk)
	if n >= arity {
		args := make([]string, arity)
		n -= arity
		for i := 0; i < arity; i++ {
			if args[i], ok = pstk[n+i].(string); !ok {
				panic("function argument isn't a symbol")
			}
		}
		p.Data = pstk[:n]
		return args
	}
	panic("PopFnArgs: unexpected argument count")
}

// BaseFrame returns current stack size.
//
func (p *Stack) BaseFrame() int {

	return len(p.Data)
}

// SetFrame sets stack to new size.
//
func (p *Stack) SetFrame(n int) {

	p.Data = p.Data[:n]
}

// -----------------------------------------------------------------------------

type variable struct {
	Name  int
	SName string
}

type variables struct {
	vars          []interface{}
	symtbl        map[string]int // symbol table
	ReverseSymtbl map[int]string // symbol table
}

func (v variables) String() string {
	return fmt.Sprintf("vars: %v, symtbl: %v", v.vars, v.symtbl)
}

// initVars initializes the variable table.
//
func (p *variables) initVars(symtbl map[string]int) {

	n := len(symtbl)
	if n > 0 {
		p.vars = make([]interface{}, n, 10000)
	}
	p.symtbl = symtbl

	p.ReverseSymtbl = make(map[int]string, len(symtbl))

	for k, v := range symtbl {
		p.ReverseSymtbl[v] = k
	}
}

// ResizeVars is reserved for internal use.
//
func (p *variables) ResizeVars() {

	vars := p.vars
	n := len(p.symtbl)
	i := len(vars)
	if i >= n {
		return
	}
	if i == 0 {
		vars = make([]interface{}, 0, n+10000)
	}
	for ; i < n; i++ {
		vars = append(vars, qlang.Undefined)
	}
	p.vars = vars

	p.ReverseSymtbl = make(map[int]string, len(p.symtbl))

	for k, v := range p.symtbl {
		p.ReverseSymtbl[v] = k
	}

}

// Vars is deprecated. please use `CopyVars` method.
//
func (p *variables) Vars() []interface{} {

	return p.vars
}

// CopyVars copies and returns all variables.
//
func (p *variables) CopyVars() map[string]interface{} {

	vars := make(map[string]interface{})
	for k, name := range p.symtbl {
		vars[k] = p.vars[name]
	}
	return vars
}

func (p *variables) VarsInfo() string {

	var sb strings.Builder

	for k, name := range p.symtbl {
		sb.WriteString(fmt.Sprintf("%v(%v, %v, %T): %v, ", k, name, &p.vars[name], p.vars[name], p.vars[name]))
	}

	return sb.String()
}

// ResetVars resets all variables.
//
func (p *variables) ResetVars(vars map[string]interface{}) {

	for k, name := range p.symtbl {
		p.vars[name] = vars[k]
	}

	p.ReverseSymtbl = make(map[int]string, len(p.symtbl))

	for k, v := range p.symtbl {
		p.ReverseSymtbl[v] = k
	}

}

// Var returns a variable value.
//
func (p *variables) Var(name string) interface{} {

	if k, ok := p.symtbl[name]; ok {
		return p.vars[k]
	}
	return qlang.Undefined
}

// GetVar returns a variable value.
//
func (p *variables) GetVar(name string) (v interface{}, ok bool) {

	if k, ok := p.symtbl[name]; ok {
		return p.vars[k], true
	}
	return qlang.Undefined, false
}

func (p *variables) GetVarWithIndex(name string) (v interface{}, idx int, ok bool) {

	if k, ok := p.symtbl[name]; ok {
		return p.vars[k], k, true
	}
	return qlang.Undefined, -1, false
}

func (p *variables) GetVarName(name int) string {

	if k, ok := p.ReverseSymtbl[name]; ok {
		return k
	}
	return ""
}

// SetVar sets a variable value.
//
func (p *variables) SetVar(name string, v interface{}) {

	k, ok := p.symtbl[name]
	if !ok {
		k = len(p.symtbl)
		if k != len(p.vars) {
			panic("ERROR: Variables need to resize (call `ResizeVars` method please)")
		}
		p.symtbl[name] = k
		p.vars = append(p.vars, qlang.Undefined)
	}
	p.vars[k] = v
}

// UnsetVar deletes a variable.
//
func (p *variables) UnsetVar(name string) {

	if k, ok := p.symtbl[name]; ok {
		p.vars[k] = qlang.Undefined
	}
}

// FastGetVar returns a variable value.
//
func (p *variables) FastGetVar(name int) interface{} {

	return p.vars[name]
}

// FastRefVar returns a variable address.
//
func (p *variables) FastRefVar(name int) *interface{} {

	return &p.vars[name]
}

// FastSetVar sets a variable value.
//
func (p *variables) FastSetVar(name int, v interface{}) {

	p.vars[name] = v
}

// -----------------------------------------------------------------------------
// type Context

type theDefer struct {
	next  *theDefer
	start int
	end   int
}

// A Context represents the context of an executor.
//
type Context struct {
	variables
	Stack  *Stack
	Code   *Code
	parent *Context
	defers *theDefer
	modmgr *moduleMgr
	Recov  interface{}
	ret    interface{}
	export []string
	ip     int
	base   int
	onsel  bool // on select
	noextv bool // don't cache extern var

	NoExit           bool
	LineStack        []string
	CurrentLineCount int
	CurrentLine      string
}

// func (v Context) String() string {
// 	return fmt.Sprintf("variables")
// }

// NewContextEx returns a new context of an executor.
//
func NewContextEx(symtbl map[string]int) *Context {

	mods := make(map[string]*importMod)
	modmgr := &moduleMgr{
		mods: mods,
	}
	p := &Context{modmgr: modmgr}
	p.initVars(symtbl)
	return p
}

// Exports returns a module exports.
//
func (p *Context) Exports() map[string]interface{} {

	export := make(map[string]interface{}, len(p.export))
	for _, name := range p.export {
		export[name], _ = p.GetVar(name)
	}
	return export
}

// ExecBlock executes an anonym function.
//
func (p *Context) ExecBlock(ip, ipEnd int, symtbl map[string]int) {

	mod := NewFunction(nil, ip, ipEnd, symtbl, nil, false)
	mod.ExtCall(p)
}

// ExecDefers executes defer blocks.
//
func (p *Context) ExecDefers() {

	d := p.defers
	if d == nil {
		return
	}

	p.defers = nil
	code := p.Code
	stk := p.Stack
	for {
		code.Exec(d.start, d.end, stk, p)
		d = d.next
		if d == nil {
			break
		}
	}
}

// -----------------------------------------------------------------------------

// A Error represents a qlang runtime error.
//
type Error struct {
	Err   error
	File  string
	Line  int
	Stack []byte
}

func (p *Error) Error() string {

	var sep string
	var stk []byte
	if qlang.DumpStack {
		stk = p.Stack
		sep = "\n\n"
	}

	if p.Line == 0 {
		return fmt.Sprintf("%v%s%s", p.Err, sep, stk)
	}
	if p.File == "" {
		return fmt.Sprintf("line %d: %v%s%s", p.Line, p.Err, sep, stk)
	}
	return fmt.Sprintf("%s:%d: %v%s%s", p.File, p.Line, p.Err, sep, stk)
}

// -----------------------------------------------------------------------------
// type Code

// A Instr represents a instruction of the executor.
//
type Instr interface {
	Exec(stk *Stack, ctx *Context)
}

// func (v Instr) String() string {
// 	return fmt.Sprintf("%v", v)
// }

// RefToVar converts a value reference instruction into a assignable variable instruction.
//
type RefToVar interface {
	ToVar() Instr
}

type optimizableArityGetter interface {
	OptimizableGetArity() int
}

type IpFileLine struct {
	Ip   int
	Line int
	File string
}

func (v IpFileLine) String() string {
	return fmt.Sprintf("ip: %v, line: %v, file: %v", v.Ip, v.Line, v.File)
}

// A Code represents generated instructions to execute.
//
type Code struct {
	Data  []Instr
	Lines []*IpFileLine // ip => (file,line)
	Debug bool
}

func (v Code) String() string {
	var sb strings.Builder

	for i, vi := range v.Data {
		sb.WriteString(fmt.Sprintf("\n [%v] -> %#v; ", i, vi))
	}

	return fmt.Sprintf("data: %v\n, lines: \n%v", sb.String(), v.Lines)
}

func (p *Code) GetCurrentLine(idxA int) int {
	rs := 0

	if p.Lines == nil {
		return rs
	}

	for _, v := range p.Lines {
		if idxA < v.Ip {
			return v.Line
		}

		rs = v.Line
	}

	return rs
}

// A ReservedInstr represents a reserved instruction to be assigned.
//
type ReservedInstr struct {
	Code *Code
	Idx  int
}

// New returns a new Code object.
//
func New(data ...Instr) *Code {

	return &Code{data, nil, (os.Getenv("GOXDEBUG") == "true") || (os.Getenv("GOXVERBOSE") == "true")}
}

// CodeLine informs current file and line.
//
func (p *Code) CodeLine(file string, line int) {
	p.Lines = append(p.Lines, &IpFileLine{Ip: len(p.Data), File: file, Line: line})
}

// Line returns file line of a instruction position.
//
func (p *Code) Line(ip int) (file string, line int) {

	idx := sort.Search(len(p.Lines), func(i int) bool {
		return ip < p.Lines[i].Ip
	})
	if idx < len(p.Lines) {
		t := p.Lines[idx]
		return t.File, t.Line
	}
	return "", 0
}

// Len returns code length.
//
func (p *Code) Len() int {

	return len(p.Data)
}

// Reserve reserves an instruction and returns it.
//
func (p *Code) Reserve() ReservedInstr {

	idx := len(p.Data)
	p.Data = append(p.Data, nil)
	return ReservedInstr{p, idx}
}

// Set sets a reserved instruction.
//
func (p ReservedInstr) Set(code Instr) {

	p.Code.Data[p.Idx] = code
}

// Next returns next instruction position.
//
func (p ReservedInstr) Next() int {

	return p.Idx + 1
}

// Delta returns distance from b to p.
//
func (p ReservedInstr) Delta(b ReservedInstr) int {

	return p.Idx - b.Idx
}

// CheckConst returns the value, if code[ip] is a const instruction.
//
func (p *Code) CheckConst(ip int) (v interface{}, ok bool) {

	if instr, ok := p.Data[ip].(*IPush); ok {
		return instr.V, true
	}
	return
}

func appendInstrOptimized(data []Instr, instr Instr, arity int) []Instr {

	n := len(data)
	base := n - arity
	for i := base; i < n; i++ {
		if _, ok := data[i].(*IPush); !ok {
			return append(data, instr)
		}
	}
	args := make([]interface{}, arity)
	for i := base; i < n; i++ {
		args[i-base] = data[i].(*IPush).V
	}
	stk := &Stack{Data: args}
	instr.Exec(stk, nil)
	return append(data[:base], Push(stk.Data[0]))
}

// Block appends some instructions to code.
//
func (p *Code) Block(code ...Instr) int {

	for _, instr := range code {
		if g, ok := instr.(optimizableArityGetter); ok {
			arity := g.OptimizableGetArity()
			p.Data = appendInstrOptimized(p.Data, instr, arity)
		} else {
			p.Data = append(p.Data, instr)
		}
	}
	return len(p.Data)
}

// ToVar converts the last instruction from ref to var.
//
func (p *Code) ToVar() {

	data := p.Data
	idx := len(data) - 1

	// fmt.Printf("ToVar: %#v, %#v, %#v\n", idx, data[idx], data)
	if cvt, ok := data[idx].(RefToVar); ok {
		data[idx] = cvt.ToVar()
	} else {
		panic("expr is not assignable")
	}
}

// // GetStatus show file and line
// func (p *Code) GetStatus() string {
// 	file, line := p.Line(ctx.ip - 1)

// 	return fmt.Sprintf("File: %v, Line: %v", file, line)
// }

// Exec executes a code block from ip to ipEnd.
//
var codeRingG *tk.StringRing = nil

func (p *Code) Exec(ip, ipEnd int, stk *Stack, ctx *Context) {

	if codeRingG == nil {
		codeRingG = tk.NewStringRing(10)
	}

	defer func() {
		if e := recover(); e != nil {
			if p.Debug {
				if e != ErrReturn {
					fmt.Println("\n--- error trace ---\n")

					startT := ctx.ip - 20
					if startT < ip {
						startT = ip
					}

					for i := startT; i <= ctx.ip; i++ {
						instr := p.Data[i]

						fmt.Printf("[%v, line: %v] %s %v \n", i, p.GetCurrentLine(i), instrName(instr), instr)
					}

					ep, ok := e.(*Error)
					if !ok {
						fmt.Printf("e: %#v, len: %v, end: %v, line: %v\n", e, p.Len(), ipEnd, p.GetCurrentLine(ctx.ip))

					} else {
						fmt.Printf("e: %#v, file: %v, len: %v, end: %v, line: %v, stack: %v\n", ep.Err, ep.File, p.Len(), ipEnd, p.GetCurrentLine(ctx.ip), string(ep.Stack))

					}

				}
			}

			if e == ErrReturn {
				panic(e)
			}

			// fmt.Printf("Current line: %v\n", p.GetCurrentLine(ctx.ip))
			lineT := p.GetCurrentLine(ctx.ip)
			// fmt.Printf("Current line: %v\n", lineT)

			if err, ok := e.(*Error); ok {
				if err.Line <= 0 && lineT > 0 {
					err.Line = lineT
				}
				fmt.Printf("Code stack: \n%v\n", codeRingG.GetString())
				panic(err)
			}
			err, ok := e.(error)
			if !ok {
				if s, ok := e.(string); ok {
					err = errors.New(s)
				} else {
					// fmt.Printf("Current line: %v\n", lineT)
					fmt.Printf("Code stack: \n%v\n", codeRingG.GetString())
					panic(e)
				}
			}
			file, line := p.Line(ctx.ip - 1)
			err = &Error{
				Err:   err,
				File:  file,
				Line:  line,
				Stack: debug.Stack(),
			}
			fmt.Printf("Code stack: \n%v\n", codeRingG.GetString())
			panic(err)
		}
	}()

	if p.Debug {
		fmt.Printf("%v, %v\n", p.Lines, p.GetCurrentLine(ip))
	}

	ctx.ip = ip
	data := p.Data
	for ctx.ip != ipEnd {
		instr := data[ctx.ip]

		if ctx.ip >= p.Len() {
			break
		}

		instr.Exec(stk, ctx)

		codeRingG.Push(fmt.Sprintf("[%v/%v] %s %#v", ctx.ip, p.Len(), instrName(instr), instr))
		if p.Debug {
			fmt.Printf("[%v/%v] %s %#v\n", ctx.ip, p.Len(), instrName(instr), instr)
		}

		if p.Debug {
			fmt.Println("STACK:", stk)
			fmt.Println("VARS:", ctx.VarsInfo())
		}

		ctx.ip++
	}
}

// Dump dumps code instructions within a range.
//
func (p *Code) Dump(ranges ...int) {

	start := 0
	end := len(p.Data)
	if len(ranges) > 0 {
		start = ranges[0]
		if len(ranges) > 1 {
			end = ranges[1]
		}
	}
	for i, instr := range p.Data[start:end] {
		fmt.Printf("==> %04d: %s %v\n", i+start, instrName(instr), instr)
	}
}

func instrName(instr interface{}) string {

	if instr == nil {
		return "<nil>"
	}
	t := reflect.TypeOf(instr).String()
	if strings.HasPrefix(t, "*") {
		t = t[1:]
	}
	if strings.HasPrefix(t, "exec.") {
		t = t[5:]
	}
	if strings.HasPrefix(t, "i") {
		t = t[1:]
	}
	return t
}

// -----------------------------------------------------------------------------
