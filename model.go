package main

import "fmt"

var thriftIndent = "    "

type TService struct {
	//map[methodName]Method
	Methods map[string]*TMethod
	Name    string
}

func (t *TService) formatMethods() string {
	if len(t.Methods) == 0 {
		return ""
	}
	str := "\n"
	for _, m := range t.Methods {
		str += fmt.Sprintf("%v\n", m.TString())
	}
	return str
}
func (t *TService) TString() string {
	return fmt.Sprintf("service %v {%v}", t.Name, t.formatMethods())
}

type TMethod struct {
	//map[fieldIndex]Field
	Parameters map[int]*TField
	Response   ThriftType

	Name      string
	Exception *TException
}

func (t *TMethod) formatRequestFields() string {
	if len(t.Parameters) == 0 {
		return ""
	}
	str := "\n"
	for idx, f := range t.Parameters {
		str += fmt.Sprintf("%v%v%v: %v %v,\n", thriftIndent, thriftIndent, idx, f.Type.TName(), f.Name)
	}
	return str + thriftIndent
}

func (t *TMethod) TString() string {
	return fmt.Sprintf(`%v%v %v(%v) throws (1: %v);`, thriftIndent, t.Response.TName(), t.Name, t.formatRequestFields(), t.Exception.TName())
}

type TField struct {
	Name string
	Type ThriftType
}

type ThriftSyntax interface {
	TString() string
}

type ThriftType interface {
	TName() string
}

type TBinary struct{}

func (t *TBinary) TName() string {
	return "binary"
}

type TVoid struct{}

func (t *TVoid) TName() string {
	return "void"
}

type TBool struct{}

func (t *TBool) TName() string {
	return "bool"
}

type TByte struct{}

func (t *TByte) TName() string {
	return "byte"
}

type TDouble struct{}

func (t *TDouble) TName() string {
	return "double"
}

type TI8 struct{}

func (t *TI8) TName() string {
	return "i8"
}

type TI16 struct{}

func (t *TI16) TName() string {
	return "i16"
}

type TI32 struct{}

func (t *TI32) TName() string {
	return "i32"
}

type TI64 struct{}

func (t *TI64) TName() string {
	return "i64"
}

type TString struct{}

func (t *TString) TName() string {
	return "string"
}

type TMap struct {
	KeyType   ThriftType
	ValueType ThriftType
}

func (t *TMap) TName() string {
	return "map<" + t.KeyType.TName() + ", " + t.ValueType.TName() + ">"
}

type TList struct {
	Type ThriftType
}

func (t *TList) TName() string {
	return "list<" + t.Type.TName() + ">"
}

type TSet struct {
	Type ThriftType
}

func (t *TSet) TName() string {
	return "set<" + t.Type.TName() + ">"
}

type TEnumField struct {
	Name string
}

func (t *TEnumField) TName() string {
	return t.Name
}

type TEnum struct {
	//map[index]EnumField
	Fields map[int]*TEnumField
	Name   string
}

func (t *TEnum) formatFields() string {
	if len(t.Fields) == 0 {
		return ""
	}
	str := "\n"
	for idx, f := range t.Fields {
		str += fmt.Sprintf("%v%v = %v\n", thriftIndent, f.TName(), idx)
	}
	return str
}

func (t *TEnum) TString() string {
	return fmt.Sprintf(`enum %v {%v}`, t.Name, t.formatFields())
}

func (t *TEnum) TName() string {
	return t.Name
}

type TStruct struct {
	Name   string
	Fields map[int]*TField
}

func (t *TStruct) formatFields() string {
	if len(t.Fields) == 0 {
		return ""
	}
	str := "\n"
	for idx, f := range t.Fields {
		str += fmt.Sprintf("%v%v: %v %v\n", thriftIndent, idx, f.Type.TName(), f.Name)
	}
	return str
}

func (t *TStruct) TString() string {
	return fmt.Sprintf("struct %v {%v}", t.TName(), t.formatFields())
}

func (t *TStruct) TName() string {
	return t.Name
}

type TException struct {
	*TStruct
}

func (t *TException) TString() string {
	return fmt.Sprintf("exception %v {%v}", t.TName(), t.formatFields())
}

func (t *TException) TName() string {
	return t.Name
}
