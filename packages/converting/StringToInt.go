package converting

import (
	"fmt"
	"reflect"
	"strconv"
)

func StringToInt() {
	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}
