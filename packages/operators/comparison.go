package operators

import "fmt"

func Comparison() {
	var (
		x int = 4
		y int = 3
	)

	// Equal
	fmt.Println(x == y) //	True if x is equal to y

	// Not equal
	fmt.Println(x != y) //	True if x is not equal to y

	// Less than
	fmt.Println(x < y) //	True if x is less than y

	// Less than or equal to
	fmt.Println(x <= y) //	True if x is less than or equal to y

	// Greater than
	fmt.Println(x > y) //	True if x is greater than y

	// Greater than or equal to
	fmt.Println(x >= y) //	True if x is greater than or equal to y

}
