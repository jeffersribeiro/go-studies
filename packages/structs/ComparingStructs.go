package structs

import "fmt"

func ComparingStructs() {
	var rect1 = rectangle{10, 20, "green"}
	rect2 := rectangle{length: 20, breadth: 10, color: "red"}

	if rect1 == rect2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	/*
		rect3 := new(rectangle)
		var rect4 = rectangle{length: 20, breadth: 10, color: "red"}

		if rect3 == rect4 {
		  fmt.Println(true)
		} else {
		  fmt.Println(false)
		}
	*/
}
