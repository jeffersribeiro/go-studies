package operators

import "fmt"

func Assignment() {

	var (
		x int = 4
		y int = 2
	)

	x = y // Assign	x = y
	fmt.Println("=", x)

	x += y // Add and assign	x = x + y
	fmt.Println("+=", x)

	x -= y // Subtract and assign	x = x - y
	fmt.Println("-=", x)

	x *= y // Multiply and assign	x = x * y
	fmt.Println("*=", x)

	x /= y // Divide and assign quotient	x = x / y
	fmt.Println("/=", x)

	x %= y // Divide and assign modulus	x = x % y
	fmt.Println("%=", x)

}
