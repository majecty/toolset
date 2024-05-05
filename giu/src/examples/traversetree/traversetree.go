package main

import (
	"fmt"
	"reflect"
)

type Parent struct {
	Children    []Child
	ParentField string
}

type Child struct {
	ChildField string
}

func main() {
	fmt.Println("Hello, world!")

	parent := Parent{
		Children: []Child{
			Child{
				ChildField: "child1",
			},
			Child{
				ChildField: "child2",
			},
		},
		ParentField: "parent",
	}

	fmt.Printf("Parent: %v\n", parent)
	(&Traverser{value: reflect.ValueOf(parent)}).Traverse()

}

type Traverser struct {
	value reflect.Value
}

func (t *Traverser) Traverse() {
	switch t.value.Kind() {
	case reflect.Struct:
	case reflect.Interface:
	case reflect.Ptr:

	case reflect.String:
	case reflect.Map:
	case reflect.Array, reflect.Slice:

	case reflect.Bool:
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Complex64, reflect.Complex128:
	case reflect.Float32, reflect.Float64:

	case reflect.Chan:
	case reflect.Func:
	case reflect.UnsafePointer:
	case reflect.Invalid:
	default:
	}
}
