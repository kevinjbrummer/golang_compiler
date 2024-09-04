package object

import (
	"bytes"
	"fmt"
	"goblin/ast"
	"goblin/color"
	"hash/fnv"
	"strings"
)

type ObjectType string
type BuiltinFunction func(args ...Object) Object

var hashCache = make(map[string]HashKey)

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	STRING_OBJ = "STRING"
	NULL_OBJ = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ = "ERROR"
	FUNCTION_OBJ = "OBJECT"
	BUILTIN_OBJ = "BUILTIN"
	ARRAY_OBJ = "ARRAY"
	HASH_OBJ = "HASH"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

type Null struct {}

func (n *Null) Inspect() string {
	return "null"
}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

type Error struct {
	Message string
}

func (e *Error) Inspect() string {
	return color.ColorWrapper(color.RED, "ERROR: ") + e.Message
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

type Function struct {
	Parameters []*ast.Identifier
	Body *ast.BlockStatement
	Env *Environment
}

func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	
	return out.String()
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

type String struct {
	Value string
}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

type Builtin struct {
	Fn BuiltinFunction
}

func(b *Builtin) Inspect() string {
	return "builtin function"
}

func(b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}

type Array struct {
	Elements []Object
}

func(a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}

	for _, el := range a.Elements {
		elements = append(elements, el.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func(a *Array) Type() ObjectType {
	return ARRAY_OBJ
}

type HashKey struct {
	Type ObjectType
	Value uint64
}

func getHashKey(objType ObjectType, value uint64) HashKey {
	key := fmt.Sprintf("%s-%d", objType, value)
	hashKey, ok := hashCache[key]; 
	if ok { 
		return hashKey
	}

	hashKey = HashKey{Type: objType, Value: value} 
	hashCache[key] = hashKey
	return hashKey
}

func (b *Boolean) HashKey() HashKey {
	var value uint64
	
	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return getHashKey(b.Type(), value)
}

func (i *Integer) HashKey() HashKey {
	return getHashKey(i.Type(), uint64(i.Value))
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return getHashKey(s.Type(), h.Sum64())
}

type HashPair struct {
	Key Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

func (h *Hash) Type() string {
	return HASH_OBJ
}

type Hashable interface {
	HashKey() HashKey
}