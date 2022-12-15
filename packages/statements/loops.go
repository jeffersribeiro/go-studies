package statements

import "fmt"

func Loops() {
	fmt.Println()
	fmt.Println("** Begin Loops **")

	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	i := 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println()
	fmt.Println("** End Loops **")
}

func ForRange() {
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println(" Index: ", index, " Element: ", element)
	}
}

/**
Output:
	Index:  Japan  Element:  Tokyo
	Index:  China  Element:  Beijing
	Index:  Canada  Element:  Ottawa
*/
