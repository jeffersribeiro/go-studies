package functions

import "fmt"

func Update(a *int, t *string) {
	*a = *a + 5
	*t = *t + " Doe"

}

func TestUpdate() {
	var age = 20
	var name = "John"
	fmt.Println("Before: ", name, age)

	Update(&age, &name)

	fmt.Println("After: ", name, age)

}

/**
Ouput:
Before: John 20
After : John Doe 25
*/
