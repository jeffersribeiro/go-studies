package converting

import (
	"fmt"
	"reflect"
	"strconv"
)

func StringToBool() {
	b := true
	s := fmt.Sprintf("%v", b)

	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}

func BoolToString() {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))

	var s string = strconv.FormatBool(true)

	fmt.Println(reflect.TypeOf(s))
}
