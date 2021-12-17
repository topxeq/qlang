package exec

import (
	"errors"
	"fmt"
	"time"

	// "fmt"
	// "reflect"
	// "unsafe"

	qlang "github.com/topxeq/qlang/spec"
)

var (
	// ErrAssignWithoutVal is returned when variable assign without value
	ErrAssignWithoutVal = errors.New("variable assign without value")

	// ErrMultiAssignExprMustBeSlice is returned when expression of multi assignment must be a slice
	ErrMultiAssignExprMustBeSlice = errors.New("expression of multi assignment must be a slice")
)

// -----------------------------------------------------------------------------

type IRef struct {
	Name  int
	SName string
}

func (p *IRef) Exec(stk *Stack, ctx *Context) {
	snameT, _ := ctx.ReverseSymtbl[p.Name]

	p.SName = snameT

	stk.Push(ctx.getRef(p.Name))
}

func (p *IRef) ToVar() Instr {
	// fmt.Printf("iRef ToVar: %v\n", p.name)

	return &IVar{Name: p.Name}
}

func (p *Context) getRef(name int) interface{} {
	// fmt.Printf("getRef: %v\n", name)

	if name < symbolIndexMax {
		return p.FastGetVar(name)
	}
	scope := name >> symbolNameBits
	for scope > 0 {
		p = p.parent
		scope--
	}
	return p.FastGetVar(name & (symbolIndexMax - 1))
}

func (p *Context) getP(name int) interface{} {

	if name < symbolIndexMax {

		// valof := reflect.ValueOf(p.vars[name])

		// pof := reflect.ValueOf(&p.vars[name])

		// fmt.Printf("val: %T, %v, %#v, %v, pof: %T, %v, %#v, %v\n", valof, valof, valof, valof.Type(), pof, pof, pof, pof.Type())

		switch nv := p.vars[name].(type) {
		case int:
			return &nv
		case []int:
			return &nv
		case uint:
			return &nv
		case uint8: // byte
			return &nv
		case []uint8: // []byte
			return &nv
		case uint16:
			return &nv
		case uint32:
			return &nv
		case uint64:
			return &nv
		case int8:
			return &nv
		case int16:
			return &nv
		case int32: // rune
			return &nv
		case []int32: // []rune
			return &nv
		case int64:
			return &nv
		case []int64:
			return &nv
		case complex64:
			return &nv
		case []complex64:
			return &nv
		case complex128:
			return &nv
		case []complex128:
			return &nv
		case float32:
			return &nv
		case []float32:
			return &nv
		case float64:
			return &nv
		case []float64:
			return &nv
		case bool:
			return &nv
		case string:
			// fmt.Printf("%T %v %#v\n", p.vars[name], p.vars[name], p.vars[name])
			// fmt.Printf("%T %v %#v\n", &p.vars[name], &p.vars[name], &p.vars[name])
			return &nv //(*string)(unsafe.Pointer(&p.vars[name]))
		case []string:
			return &nv
		case map[string]string:
			return &nv
		case []map[string]string:
			return &nv
		case time.Time:
			return &nv
		case []interface{}:
			return &nv
		case map[string]interface{}:
			return &nv
		case []map[string]interface{}:
			return &nv
		case *interface{}:
			return &nv
		case interface{}:
			return &nv
		default:
			fmt.Printf("unknown type pointer: %T, %v, %#v\n", nv, nv, nv)
			return &nv
		}

		return &p.vars[name]
	}

	scope := name >> symbolNameBits
	for scope > 0 {
		p = p.parent
		scope--
	}

	switch nv := p.vars[name&(symbolIndexMax-1)].(type) {
	case int:
		return &nv
	case []int:
		return &nv
	case uint:
		return &nv
	case uint8: // byte
		return &nv
	case []uint8: // []byte
		return &nv
	case uint16:
		return &nv
	case uint32:
		return &nv
	case uint64:
		return &nv
	case int8:
		return &nv
	case int16:
		return &nv
	case int32: // rune
		return &nv
	case []int32: // []rune
		return &nv
	case int64:
		return &nv
	case []int64:
		return &nv
	case complex64:
		return &nv
	case []complex64:
		return &nv
	case complex128:
		return &nv
	case []complex128:
		return &nv
	case float32:
		return &nv
	case []float32:
		return &nv
	case float64:
		return &nv
	case []float64:
		return &nv
	case bool:
		return &nv
	case string:
		return &nv
	case []string:
		return &nv
	case map[string]string:
		return &nv
	case []map[string]string:
		return &nv
	case time.Time:
		return &nv
	case []interface{}:
		return &nv
	case map[string]interface{}:
		return &nv
	case []map[string]interface{}:
		return &nv
	case *interface{}:
		return &nv
	case interface{}:
		return &nv
	default:
		fmt.Printf("unknown type pointer: %T, %v, %#v\n", nv, nv, nv)
		return &nv
	}

	return &p.vars[name&(symbolIndexMax-1)]
}

func (p *Context) getPointer(name int) interface{} {

	if name < symbolIndexMax {
		return p.FastRefVar(name)
	}
	scope := name >> symbolNameBits
	for scope > 0 {
		p = p.parent
		scope--
	}
	return p.FastRefVar(name & (symbolIndexMax - 1))
}

// SymbolIndex returns symbol index value.
//
func SymbolIndex(id, scope int) int {
	return id | (scope << symbolNameBits)
}

// Ref returns an instruction that refers a variable.
//
func Ref(name int) Instr {
	// fmt.Printf("Ref: %v\n", name)
	return &IRef{Name: name}
}

// -----------------------------------------------------------------------------
// Var

type IVar struct {
	Name  int
	SName string
}

func (p *IVar) Exec(stk *Stack, ctx *Context) {
	snameT, _ := ctx.ReverseSymtbl[p.Name]

	p.SName = snameT

	stk.Push(&variable{Name: p.Name, SName: snameT})
}

// Var returns a Var instruction.
//
func Var(name int) Instr {
	// fmt.Printf("Var: %v\n", name)
	return &IVar{Name: name}
}

// -----------------------------------------------------------------------------

type iGet int

func (p iGet) Exec(stk *Stack, ctx *Context) {

	k, ok1 := stk.Pop()
	o, ok2 := stk.Pop()
	if !ok1 || !ok2 {
		panic("unexpected to call `Get` instruction")
	}
	stk.Push(qlang.Get(o, k))
}

func (p iGet) ToVar() Instr {

	return GetVar
}

// Get is the Get instruction.
//
var Get Instr = iGet(0)

// -----------------------------------------------------------------------------

type iGetVar int

func (p iGetVar) Exec(stk *Stack, ctx *Context) {

	k, ok1 := stk.Pop()
	o, ok2 := stk.Pop()
	if !ok1 || !ok2 {
		panic("unexpected to call `GetVar` instruction")
	}
	stk.Push(&qlang.DataIndex{Data: o, Index: k})
}

// GetVar is the Get instruction.
//
var GetVar Instr = iGetVar(0)

// -----------------------------------------------------------------------------

type iGfnRef struct {
	val   interface{}
	toVar func() Instr
	SName string
}

func (p *iGfnRef) Exec(stk *Stack, ctx *Context) {

	stk.Push(p.val)
}

func (p *iGfnRef) ToVar() Instr {

	return p.toVar()
}

// GfnRef returns an instruction that refers a fntable item.
//
func GfnRef(v interface{}, toVar func() Instr, sname string) Instr {
	return &iGfnRef{v, toVar, sname}
}

type iAddrOf struct {
	name int
}

func AddrOf(name int) Instr {
	return &iAddrOf{name}
}

// func ToPointer(pointerA *interface{}) interface{} {
// 	v := *pointerA

// 	fmt.Printf("ToPointer: %T\n", v)

// 	switch nv := v.(type) {
// 	case int:
// 		return (*int)(unsafe.Pointer(pointerA))
// 	case []int:
// 		return (*[]int)(unsafe.Pointer(pointerA))
// 	case uint:
// 		return (*uint)(unsafe.Pointer(pointerA))
// 	case uint8: // byte
// 		return (*uint8)(unsafe.Pointer(pointerA))
// 	case []uint8: // []byte
// 		return (*[]uint8)(unsafe.Pointer(pointerA))
// 	case uint16:
// 		return (*uint16)(unsafe.Pointer(pointerA))
// 	case uint32:
// 		return (*uint32)(unsafe.Pointer(pointerA))
// 	case uint64:
// 		return (*uint64)(unsafe.Pointer(pointerA))
// 	case int8:
// 		return (*int8)(unsafe.Pointer(pointerA))
// 	case int16:
// 		return (*int16)(unsafe.Pointer(pointerA))
// 	case int32: // rune
// 		return (*int32)(unsafe.Pointer(pointerA))
// 	case []int32: // []rune
// 		return (*[]int32)(unsafe.Pointer(pointerA))
// 	case int64:
// 		return (*int64)(unsafe.Pointer(pointerA))
// 	case []int64:
// 		return (*[]int64)(unsafe.Pointer(pointerA))
// 	case complex64:
// 		return (*complex64)(unsafe.Pointer(pointerA))
// 	case []complex64:
// 		return (*[]complex64)(unsafe.Pointer(pointerA))
// 	case complex128:
// 		return (*complex128)(unsafe.Pointer(pointerA))
// 	case []complex128:
// 		return (*[]complex128)(unsafe.Pointer(pointerA))
// 	case float32:
// 		return (*float32)(unsafe.Pointer(pointerA))
// 	case []float32:
// 		return (*[]float32)(unsafe.Pointer(pointerA))
// 	case float64:
// 		return (*float64)(unsafe.Pointer(pointerA))
// 	case []float64:
// 		return (*[]float64)(unsafe.Pointer(pointerA))
// 	case bool:
// 		return (*bool)(unsafe.Pointer(pointerA))
// 	case string:
// 		return (*string)(unsafe.Pointer(pointerA))
// 	case []string:
// 		return &nv
// 		// return (*[]string)(unsafe.Pointer(pointerA))
// 	case map[string]string:
// 		return (*map[string]string)(unsafe.Pointer(pointerA))
// 	case []map[string]string:
// 		return (*[]map[string]string)(unsafe.Pointer(pointerA))
// 	case time.Time:
// 		return (*time.Time)(unsafe.Pointer(pointerA))
// 	case []interface{}:
// 		return (*[]interface{})(unsafe.Pointer(pointerA))
// 	case map[string]interface{}:
// 		return (*map[string]interface{})(unsafe.Pointer(pointerA))
// 	case []map[string]interface{}:
// 		return (*[]map[string]interface{})(unsafe.Pointer(pointerA))
// 	case *interface{}:
// 		return (*interface{})(unsafe.Pointer(pointerA))
// 	case interface{}:
// 		return pointerA
// 	}

// 	return pointerA
// }

func (p *iAddrOf) Exec(stk *Stack, ctx *Context) {
	// NOTE:: To wrap Clang!
	val := ctx.getPointer(p.name)

	// tp := reflect.ValueOf(val)
	// typeT := reflect.TypeOf(val)
	// fmt.Printf("tp: %#v, %#v, %#v\n", val, tp, tp.String())
	// tty.Debugf("iAddrOf : tp [%s]\n", tp.String())
	// fmt.Printf(": %#v, %#v, %#v, %#v, %v, %#v\n", val, tp, tp.Elem(), typeT, typeT, tp.Elem().Type())

	// stk.Push(uintptr(unsafe.Pointer(val.(*interface{}))))

	// stk.Push(ctx.getP(p.name))
	// stk.Push(ToPointer(val.(*interface{})))

	stk.Push(val.(*interface{}))
	return

	// var result uintptr
	// switch tp.String() {
	// case "*exec.Uint8":
	// 	result = uintptr(unsafe.Pointer(val.(*uint8)))
	// case "*exec.Uint16":
	// 	result = uintptr(unsafe.Pointer(val.(*uint16)))
	// case "*exec.Uint32":
	// 	result = uintptr(unsafe.Pointer(val.(*uint32)))
	// case "*exec.Uint64":
	// 	result = uintptr(unsafe.Pointer(val.(*uint64)))
	// default:
	// 	panic(fmt.Sprintf("Only c variable could be address of: %#v, %#v", val, tp.String()))
	// }
	// stk.Push(result)
}

type iAddrxOf struct {
	name int
}

func AddrxOf(name int) Instr {
	return &iAddrxOf{name}
}

func (p *iAddrxOf) Exec(stk *Stack, ctx *Context) {
	// NOTE:: To wrap Clang!
	// val := ctx.getPointer(p.name)

	// tp := reflect.ValueOf(val)
	// typeT := reflect.TypeOf(val)
	// fmt.Printf("tp: %#v, %#v, %#v\n", val, tp, tp.String())
	// tty.Debugf("iAddrOf : tp [%s]\n", tp.String())
	// fmt.Printf(": %#v, %#v, %#v, %#v, %v, %#v\n", val, tp, tp.Elem(), typeT, typeT, tp.Elem().Type())

	// stk.Push(uintptr(unsafe.Pointer(val.(*interface{}))))

	stk.Push(ctx.getP(p.name))
	// stk.Push(ToPointer(val.(*interface{})))

	// stk.Push(val.(*interface{}))
	return

	// var result uintptr
	// switch tp.String() {
	// case "*exec.Uint8":
	// 	result = uintptr(unsafe.Pointer(val.(*uint8)))
	// case "*exec.Uint16":
	// 	result = uintptr(unsafe.Pointer(val.(*uint16)))
	// case "*exec.Uint32":
	// 	result = uintptr(unsafe.Pointer(val.(*uint32)))
	// case "*exec.Uint64":
	// 	result = uintptr(unsafe.Pointer(val.(*uint64)))
	// default:
	// 	panic(fmt.Sprintf("Only c variable could be address of: %#v, %#v", val, tp.String()))
	// }
	// stk.Push(result)
}

// -----------------------------------------------------------------------------
