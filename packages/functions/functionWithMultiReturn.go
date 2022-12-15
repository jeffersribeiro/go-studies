package functions

import "fmt"

func FuncWithMultiReturn(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // return statement without specify variable name
}

func TestFuncWithMultiReturn() {
	a, p := FuncWithMultiReturn(20, 30)
	fmt.Println("Area: ", a)
	fmt.Println("Parameter: ", p)
}
