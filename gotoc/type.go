package gotoc

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"code.google.com/p/go.tools/go/types"
)

func dimFuncPtr(name string, dim []string) string {
	if len(dim) == 0 {
		return name
	}
	for i, d := range dim {
		switch d[0] {
		case '*':
			if i == len(dim)-1 || dim[i+1][0] == '*' {
				name = "*" + name
			} else {
				name = "(*" + name + ")"
			}

		case '[':
			name += d

		default:
			name = "(*" + name + ")" + d
		}
	}
	return name
}

func (cdd *CDD) Type(w *bytes.Buffer, typ types.Type) (dim []string, acds []*CDD) {
	direct := true

writeType:
	switch t := typ.(type) {
	case *types.Basic:
		if t.Kind() == types.UnsafePointer {
			w.WriteString("unsafe$Pointer")
		} else {
			types.WriteType(w, nil, t)
		}

	case *types.Named:
		tn := t.Obj()
		if tn.Name() == "error" {
			w.WriteString("error")
		} else {
			cdd.Name(w, tn, direct)
		}

	case *types.Pointer:
		typ = t.Elem()
		direct = false
		dim = append(dim, "*")
		goto writeType

	case *types.Struct:
		if t.NumFields() == 0 {
			w.WriteString("empty")
			break
		}
		w.WriteString("struct {\n")
		cdd.il++
		for i, n := 0, t.NumFields(); i < n; i++ {
			f := t.Field(i)
			cdd.indent(w)
			if tag := t.Tag(i); tag != "" {
				w.WriteString(reflect.StructTag(tag).Get("C"))
				w.WriteByte(' ')
			}
			d, a := cdd.Type(w, f.Type())
			acds = append(acds, a...)
			if !f.Anonymous() {
				w.WriteByte(' ')
				name := dimFuncPtr(f.Name(), d)
				if name == "_" {
					name += strconv.Itoa(i) + "$"
				}
				w.WriteString(name)
			}
			w.WriteString(";\n")
		}

		cdd.il--
		cdd.indent(w)
		w.WriteByte('}')

	case *types.Array:
		dim = append(dim, "["+strconv.FormatInt(t.Len(), 10)+"]")
		d, a := cdd.Type(w, t.Elem())
		dim = append(dim, d...)
		acds = append(acds, a...)

	case *types.Slice:
		w.WriteString("slice")

	case *types.Map:
		w.WriteString("map")

	case *types.Chan:
		w.WriteString("chan")

	case *types.Signature:
		res, params := cdd.signature(t, true, noNames)
		w.WriteString(res.typ)
		dim = append(dim, params)
		dim = append(dim, res.dim...)
		acds = append(acds, res.acds...)

	case *types.Interface:
		if t.NumMethods() == 0 {
			w.WriteString("interface")
			break
		}
		w.WriteString("struct {\n")
		cdd.il++
		cdd.indent(w)
		w.WriteString("interface I$;\n")
		for i := 0; i < t.NumMethods(); i++ {
			cdd.indent(w)
			f := t.Method(i)
			d, a := cdd.Type(w, f.Type())
			acds = append(acds, a...)
			w.WriteString(" " + dimFuncPtr(f.Name(), d) + ";\n")
		}
		cdd.il--
		cdd.indent(w)
		w.WriteByte('}')

	default:
		fmt.Fprintf(w, "<%T>", t)
	}
	return
}

func (cdd *CDD) TypeStr(typ types.Type) (string, []string, []*CDD) {
	buf := new(bytes.Buffer)
	dim, acds := cdd.Type(buf, typ)
	return buf.String(), dim, acds
}

type retVar struct {
	name, typ string
}

type results struct {
	fields   []*types.Var
	names    []string
	typ      string
	dim      []string
	hasNames bool
	acds     []*CDD
}

func (cdd *CDD) results(tup *types.Tuple) (res results) {
	if tup == nil {
		res.typ = "void"
		return
	}

	n := tup.Len()

	res.names = make([]string, n)
	for i := 0; i < n; i++ {
		name := tup.At(i).Name()
		switch name {
		case "":
			name = "_" + strconv.Itoa(n)

		case "_":
			res.hasNames = true

		default:
			name += "$"
			res.hasNames = true
		}
		res.names[i] = name
	}
	if n > 1 {
		res.typ, res.fields, res.acds = cdd.tupleName(tup)
		return
	}
	v := tup.At(0)
	field0 := types.NewField(v.Pos(), v.Pkg(), "_0", v.Type(), false)
	res.fields = []*types.Var{field0}
	res.typ, res.dim, res.acds = cdd.TypeStr(v.Type())
	return
}

const (
	noNames = iota
	numNames
	orgNames
)

func (cdd *CDD) signature(sig *types.Signature, recv bool, pnames int) (res results, params string) {
	params = "("
	res = cdd.results(sig.Results())
	if r := sig.Recv(); r != nil && recv {
		var (
			typ  string
			dim  []string
			acds []*CDD
		)
		if _, ok := r.Type().Underlying().(*types.Interface); ok {
			typ = "uintptr"
		} else {
			typ, dim, acds = cdd.TypeStr(r.Type())
			res.acds = append(res.acds, acds...)
		}
		var pname string
		switch pnames {
		case numNames:
			pname = "_0"
		case orgNames:
			pname = cdd.NameStr(r, true)
		}
		if pname == "" {
			params += typ + dimFuncPtr("", dim)
		} else {
			params += typ + " " + dimFuncPtr(pname, dim)
		}
		if sig.Params() != nil {
			params += ", "
		}
	}
	if p := sig.Params(); p != nil {
		for i, n := 0, p.Len(); i < n; i++ {
			if i != 0 {
				params += ", "
			}
			v := p.At(i)
			typ, dim, acds := cdd.TypeStr(v.Type())
			res.acds = append(res.acds, acds...)
			var pname string
			switch pnames {
			case numNames:
				pname = "_" + strconv.Itoa(i+1)
			case orgNames:
				pname = cdd.NameStr(v, true)
				if pname == "_$" {
					pname = "unused" + cdd.gtc.uniqueId()
				}
			}
			if pname == "" {
				params += typ + dimFuncPtr("", dim)
			} else {
				params += typ + " " + dimFuncPtr(pname, dim)
			}
		}
	}
	params += ")"
	return
}

// BUG: this mapping can be ambiguous.
func symToDol(r rune) rune {
	switch r {
	case '*', '(', ')', '[', ']':
		return '$'
	}
	return r
}

func (cdd *CDD) tupleName(tup *types.Tuple) (tupName string, fields []*types.Var, acds []*CDD) {
	n := tup.Len()
	for i := 0; i < n; i++ {
		if i != 0 {
			tupName += "$$"
		}
		name, dim, a := cdd.TypeStr(tup.At(i).Type())
		tupName += dimFuncPtr(name, dim)
		acds = append(acds, a...)

	}
	tupName = strings.Map(symToDol, tupName)

	fields = make([]*types.Var, n)
	for i := 0; i < n; i++ {
		v := tup.At(i)
		fields[i] = types.NewField(
			v.Pos(), v.Pkg(), "_"+strconv.Itoa(i), v.Type(), false,
		)
	}

	if _, ok := cdd.gtc.tupNames[tupName]; ok {
		return
	}

	cdd.gtc.tupNames[tupName] = struct{}{}

	s := types.NewStruct(fields, nil)
	o := types.NewTypeName(tup.At(0).Pos(), cdd.gtc.pkg, tupName, s)
	acd := cdd.gtc.newCDD(o, TypeDecl, 0)
	acd.structDecl(new(bytes.Buffer), tupName, s)
	cdd.DeclUses[o] = true
	acds = append(acds, acd)

	return
}
