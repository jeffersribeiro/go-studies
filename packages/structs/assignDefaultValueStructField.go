package structs

import "fmt"

type Employee3 struct {
	Name string
	Age  int
}

func (obj *Employee3) Info() {
	if obj.Name == "" {
		obj.Name = "John Doe"
	}
	if obj.Age == 0 {
		obj.Age = 25
	}
}

func AssignDefaultValueStructField() {
	emp1 := Employee3{Name: "Mr. Fred"}
	emp1.Info()

	fmt.Println(emp1)

	emp2 := Employee3{Age: 26}
	emp2.Info()

	fmt.Println(emp2)
}
