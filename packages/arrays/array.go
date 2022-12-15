package arrays

import (
	"fmt"
	"reflect"
)

/**
 * Introduction
 *
 * An Array is a data structure thta consists of a collection of elements of a single type or simply you can say a
 * special variable, which can hold more than one value at a time. The values an array holds are called its
 * ELEMENTS or ITEMS. An array holds a specific number of elements, and it cannot grow or shrink. Different
 * data type can be handled as elements in arrays such as Int, String, Boolean, and others. the index of the
 * first elements of any dimension of an array is 0, the index of the second elements of any array dimension is 1,
 * and so on.
 */

// var <name of the array>[<length>] <data type>
func declaring() {
	var intArray [2]int
	var strArray [1]string

	intArray[0] = 10
	intArray[0] = 20

	strArray[0] = "hello"

	fmt.Println(intArray[0])
	fmt.Println(strArray[0])

	fmt.Println(reflect.ValueOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())
}

/*
 * Output:
 * [10 20 30 40 50]
 * [10 20 30 0 0]
 */
func literalInitializing() {
	x := [5]int{10, 20, 30, 40, 50}   // initialized with values
	var y [5]int = [5]int{10, 20, 30} // partial assigment

	fmt.Println(x)
	fmt.Println(y)
}

/*
 * Output:
 * array
 * 3
 */
func ellipses() {
	x := [...]int{10, 20, 30}

	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(len(x))
}

/*
 * Output:
 * [0 10 0 30 0]
 */
func specifyValuesElements() {
	x := [5]int{1: 10, 3: 30}

	fmt.Println(x)
}

func loopThroughIndexed() {
	intArray := [5]int{10, 20, 30, 40, 50}

	fmt.Println("------------------Example 1-----------------")

	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	fmt.Println("------------------Example 2-----------------")

	for index, element := range intArray {
		fmt.Println(index, "=>", element)
	}

	fmt.Println("------------------Example 3-----------------")

	for _, value := range intArray {
		fmt.Println(value)
	}

	fmt.Println("------------------Example 4-----------------")
	j := 0

	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}

func copyAnArray() {
	strArray1 := [3]string{"Japan", "Australia", "Germany"}

	strArray2 := strArray1  // data is passed by value
	strArray3 := &strArray1 // data is passed by refrence

	fmt.Printf("strArray1 %v\n", strArray1) // strArray1: [Japan Australia Germany]
	fmt.Printf("strArray2 %v\n", strArray2) // strArray2: [Japan Australia Germany]

	strArray1[0] = "Canada"

	fmt.Printf("strArray1: %v\n", strArray1)   // strArray1: [Canada Australia Germany]
	fmt.Printf("strArray2: %v\n", strArray2)   // strArray2: [Japan Australia Germany]
	fmt.Printf("*strArray3: %v\n", *strArray3) // *strArray3: [Canada Australia Germany]
}

func checkIfElementExists() {
	strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strArray, "Canada"))
	fmt.Println(itemExists(strArray, "Africa"))
}

func itemExists(arrayType [5]string, item string) bool {
	arr := reflect.ValueOf(arrayType)
	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

/*
 * Output
 * Countries: [India Canada Japan Germany Italy]
 * :2 [India Canada]
 * 1:3 [Canada Japan]
 * 2: [Japan Germany Italy]
 * 2:5 [Japan Germany Italy]
 * 0:3 [India Canada Japan]
 * Last element: Italy
 * All elements: [India Canada Japan Germany Italy]
 * [India Canada Japan Germany Italy]
 * [India Canada Japan Germany Italy]
 * [India Canada Japan Germany Italy]
 * Last two elements: [Germany Italy]
 */
func filterArrayElement() {
	countries := [...]string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Printf("countries: %v/n", countries)

	fmt.Printf(":2 %v\n", countries[:2])
	fmt.Printf("1:3 %v\n", countries[1:3])
	fmt.Printf("2: %v\n", countries[2:])
	fmt.Printf("2:5 %v\n", countries[2:5])
	fmt.Printf("0:3 %v\n", countries[0:3])
	fmt.Printf("Last Element %v\n", countries[len(countries)-1])

	fmt.Printf("All Element %v\n", countries[0:len(countries)])
	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])

	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])
}

func TestArrays() {
	declaring()
	literalInitializing()
	ellipses()
	specifyValuesElements()
	loopThroughIndexed()
	copyAnArray()
	checkIfElementExists()
	filterArrayElement()
}
