package structs

import "fmt"

/*
Indroduction

 A *struct*(short for "Structure") is a collection of data fields with declared data types. golang has the ability to
 declare and create own data types by combinig one or more types, including both built-in and use-defined
 types. Each data field in struct is declared with a known type, which could by a built-in type another
 user-defined type

 Structs are the only way to create concrete user-defined types in Golanf. Struct types are declred by
 composing a fixed set of unique fields. Structs cam improce modularity and allow to create ans pass complex
 data structures around the system. you can also consider Structs as a template for creatind a data record,
 like am employee record or an-ecommerce product.

 The declaration starts wiht the keyword *type*, then a name for new struct, and finally the keyword *struct*.
 within the curly brackets, a seris of data fields are specified with a name and type.

Example

 type identifier strcut {
 	field1 data_type
 	field2 data_type
 	field3 data_type
 }
*/

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func Struct() {
	fmt.Println(rectangle{10.5, 25.10, "red"})
}
