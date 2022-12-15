package converting

import (
	"fmt"
	"reflect"
)

func Float32ToFloat64() {
	var f32 float32 = 10.6556

	fmt.Println(reflect.TypeOf(f32))

	i32 := int32(f32)

	fmt.Print(reflect.TypeOf(i32))
	fmt.Print(i32)

	f64 := float64(i32)
	fmt.Println(reflect.TypeOf(f64))
}
