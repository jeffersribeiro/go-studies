package statements

import "fmt"

// Global variable declaration
var (
	m int
	n int
)

func DataTypes() {

	var x int = 1 // interger data type
	var y int     // interger data type

	fmt.Println(x, y)

	var a, c, b = 5.25, 25.25, 14.15 // multiple float32 variables declaration

	fmt.Println(a, b, c)

	city := "Berlin"     // string variable declaration
	Country := "Germany" // variable names are case sensitive

	fmt.Println(city, Country)

	food, drink, price := "Pizza", "Pepsi", 125 // multiple type of variabel declaration in same line
	fmt.Println(food, drink, price)
	m, n = 1, 2
	fmt.Println(m, n)
}
