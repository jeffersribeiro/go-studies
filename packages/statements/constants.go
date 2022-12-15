package statements

import "fmt"

// By convention, constant names are usually written in uppercase letters.
// this is for their easy identification and differentiation from variables in the souce code

const (
	PRODUCT  = "Mobile"
	QUANTITY = 50
	PRICE    = 50.50
	STOCK    = true
)

func Constants() {
	fmt.Println()
	fmt.Println("** Begin Constants **")

	fmt.Println(
		PRODUCT,
		QUANTITY,
		PRICE,
		STOCK,
	)

	const a = 100
	fmt.Println(a)

	const s string = "Hello"
	fmt.Println(s)

	fmt.Println()
	fmt.Println("** End Constants **")

}
