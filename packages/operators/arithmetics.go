package operators

import "fmt"

var (
	increment int
	decrement int
)

// Sum of x and y
func Addition() int { return 5 + 4 }

// Subtracts one value from another
func Subtraction() int { return 5 - 4 }

// Multiplies two values
func Multiplication() int { return 5 * 4 }

// Quotient of x and y
func Division() int { return 5 / 4 }

// Remainder of x divided by y
func Modulus() int { return 5 % 4 }

func Others() {
	increment++ // Increases the value of a variable by 1
	decrement-- // Decreases the value of a variable by 1
	fmt.Println(
		increment,
		decrement,
	)

}
