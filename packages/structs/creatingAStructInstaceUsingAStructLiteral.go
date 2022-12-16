package structs

import "fmt"

type rectangle3 struct {
	length  int
	breadth int
	color   string
}

func CreatingAStructInstaceUsingAStructLiteral() {
	var rect1 = rectangle3{10, 20, "Green"}

	fmt.Println(rect1)

	var rect2 = rectangle3{length: 10, color: "Green"} // breadth value skipped
	fmt.Println(rect2)

	rect3 := rectangle3{10, 20, "Green"}
	fmt.Println(rect3)

	rect4 := rectangle3{length: 10, breadth: 20, color: "green"}
	fmt.Println(rect4)

	rect5 := rectangle3{breadth: 20, color: "Green"} // length value skipped

	fmt.Println(rect5)
}
