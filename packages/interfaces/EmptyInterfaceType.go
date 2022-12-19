package interfaces

import "fmt"

func printType(i interface{}) {
	fmt.Println(i)
}

func EmptyInterfaceType() {
	var manyType interface{}
	manyType = 100

	fmt.Println(manyType) // 100

	manyType = 200.50

	fmt.Println(manyType) // 200.5

	manyType = "Germany"

	fmt.Println(manyType) // Germany

	printType("Go Programing language") // Go programming language
	var countries = []string{"india", "japan", "canada", "australia", "russia"}
	printType(countries) // [india japan canada australia russia]

	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	printType(employee) // map[Mark:10 Sandy:20]

	country := [3]string{"Japan", "Australia", "Germany"}
	printType(country) // [Japan Australia Germany]
}

/*
the manyTypes variables is declared to be of the type interface{} and it is
able to be assigned values of different types. the printtype() function
takes aparameter of the type interface{}, hance this function can take the
values of may any valid type.
*/
