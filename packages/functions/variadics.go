package functions

import (
	"fmt"
	"reflect"
)

func variadicExample1(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
}

func variadicExample2(s ...string) {
	fmt.Println(s)
}

func calculation(str string, y ...int) int {
	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}

	return area
}

func manyVariadicTypes(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

func TestVariadics() {
	variadicExample1()                       // null null
	variadicExample2("red", "green")         // [red green]
	variadicExample2("red", "green", "blue") // [red green blue]

	fmt.Println(calculation("Rectangle", 20, 30)) // 600
	fmt.Println(calculation("Square", 20))        // 400

	manyVariadicTypes(1, "red", true, 10.5, []string{"foo", "bar", "baz"}, map[string]int{"apple": 23, "tomato": 13})
	/**
	int     -- 1
	string  -- red
	boolean -- true
	float64 -- 10.5
	slice   -- [foo, bar, baz]
	map     -- map[apple:23 tomato:13]
	*/
}
