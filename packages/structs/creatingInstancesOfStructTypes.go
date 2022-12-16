package structs

import "fmt"

type rectangle2 struct {
	length  int
	breadth int
	color   string

	geometry struct {
		area      int
		parimeter int
	}
}

func CreatingInstancesOfStructTypes() {
	var rect rectangle2
	rect.length = 10
	rect.breadth = 20

	rect.color = "Green"

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.parimeter = 2 * (rect.length + rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect)
	fmt.Println("Parimeter:", rect.geometry.parimeter)

}
