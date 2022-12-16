package structs

/*
Two instances of the retangle struct are instantiated, rect1 points to the address of the instantiated
struct and rect2 is the name of a strcut it represents.
*/
func StructInstantiationUsingnewKeyword() {
	rect1 := new(rectangle) // rect1 is a pointer to an instace of rectangle

	rect1.length = 10
	rect1.breadth = 10
	rect1.color = "Green"

	println(rect1)

	var rect2 = new(rectangle) // rect2 is an instance of retangle

	rect2.length = 10
	rect2.color = "Red"
	println(rect2)
}
