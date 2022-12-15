package operators

import "fmt"

func Logicals() {
	var (
		x float32 = 0.36
		y float32 = 4.25
		z float32 = 2
	)

	// &&	Logical And	Returns true if both statements are true
	fmt.Println(x < y && x > z)
	// ||	Logical Or	Returns true if one of the statements is true
	fmt.Println(x < y || x > z)
	// !	Logical Not	Reverse the result, returns false if the result is true
	fmt.Println(!(x == y && x > z))

}
