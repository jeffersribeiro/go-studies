package statements

import "fmt"

// Variables declaration can be grouped together into block for greater readability and code quality
var (
	product  = "Mobile"
	quantity = 50
	price    = 50.50
	inStock  = true
)

func Variables() {
	fmt.Println()
	fmt.Println("** Begin Variables **")

	fmt.Println(
		product,
		quantity,
		price,
		inStock,
	)

	var a, b = 1, 2

	var c = a + b

	fmt.Println(c)

	d := 10

	fmt.Println(d)

	e := "hello"

	fmt.Println(e)

	fmt.Println()
	fmt.Println("** End Variables **")
}
