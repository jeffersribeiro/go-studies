package statements

import "fmt"

func Conditionals() {
	fmt.Println()
	fmt.Println("** Begin Conditionals **")

	if 3%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 5 > 3 {
		fmt.Println("5 is greater")
	} else if 5 < 3 {
		fmt.Println("5 is less then 3")
	} else {
		fmt.Println("5 is equal to 3")
	}

	fmt.Println()
	fmt.Println("** End Conditionals **")
}
