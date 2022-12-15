package functions

import "fmt"

func Retangle(l int, b int) (area int) {

	parameter := 2 * (l + b)

	fmt.Println("Parameter: ", parameter)

	area = l * b
	return // return statement without specify variable name
}

/**
Output:

Parameter:  100
Area:  600
*/
