package functions

import "fmt"

// Higher-Order Functions is a functions that receives a function
// as an argument or return the function as output

// Passing Functions as Arguments to other Functions
func sum(x, y int) int {
	return x + y
}

func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}

}

// Returning Functions from other Functions
func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

// User Defined Function types in Golang

type First func(int) int
type Second func(int) First

func SquareSum(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func TestHighOrder() {
	partial := partialSum(3)
	fmt.Println(partial(7))

	fmt.Println(squareSum(5)(6)(7))
	fmt.Println(SquareSum(7)(6)(5))
}
