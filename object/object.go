package object

import (
	"fmt"
)

type ObjectType string

const (
	INTEGER_OBJ = "Integer"
	BOOLEAN_OBJ = "Boolean"
	NULL_OBJ    = "Null"
	STRING_OBJ  = "String"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// integer object
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// boolean object
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// NULL object
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// String Object
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }
