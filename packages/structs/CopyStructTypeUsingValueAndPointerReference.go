package structs

import "fmt"

/*
r2 will be the same as r1, it is a copy r1 tather than a reference to it. Any changes made to r2 will not
be applied to r1 and vice versa. When r3 is updated, the undelying memory thta uis assigned to r1 is updated
*/
func CopyStructTypeUsingValueAndPointerReference() {
	r1 := rectangle{10, 20, "Green"}

	fmt.Println(r1) // {10 20 Green}

	r2 := r1
	r2.color = "Pink"
	fmt.Println(r2) // {10 20 Pink}

	r3 := &r1

	r3.color = "Red"
	fmt.Println(r3) // &{10 20 Red}

	fmt.Println(r1) // {10 20 Red}
}

/*
As  both r1 and r3 both reference the same undelying memorym their values are the same, Printing the
values of r3 and r1 shows that values are the same.
*/
