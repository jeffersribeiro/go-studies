package structs

import (
	"fmt"
	"reflect"
)

type rectangle4 struct {
	length  float64
	breadth float64
	color   string
}

func FindTypeStruct() {

	var rec1 = rectangle4{10, 20, "green"}

	fmt.Println(reflect.ValueOf(rec1))        // FindTypeStruct.rectangle4
	fmt.Println(reflect.ValueOf(rec1).Kind()) // struct

	rect2 := rectangle4{length: 10, breadth: 20, color: "Green"}
	fmt.Println(reflect.TypeOf(rect2))         // FindTypeStruct.rectangle4
	fmt.Println(reflect.ValueOf(rect2).Kind()) // struct

	rect3 := new(rectangle4)
	fmt.Println(reflect.TypeOf(rect3))         // *FindTypeStruct.rectangle4
	fmt.Println(reflect.ValueOf(rect3).Kind()) // ptr

	var rect4 = &rectangle4{}
	fmt.Println(reflect.TypeOf(rect4))         // *FindTypeStruct.rectangle4
	fmt.Println(reflect.ValueOf(rect4).Kind()) // ptr
}
