package slices

import (
	"fmt"
	"reflect"
)

/*
 * Introduction
 *
 * A slice is a flexible and extensible data structure to implement and manage collections of data, Slices are
 * made up of multiple elements, all of the same type. A slice is a segment of dynamic arrays that can grow and
 * shrink as you see fit. Like arrays, slices are index-able and have a length. Slices have a capacity and length
 * property.
 */

func emptySlice() {
	var intSlice []int
	var strSlice []string

	fmt.Println(reflect.ValueOf(intSlice).Kind()) //slice
	fmt.Println(reflect.ValueOf(strSlice).Kind()) //slice
}

func declareSliceUsingMake() {
	var intSlice = make([]int, 10)        // when length and capacity is same
	var strSlice = make([]string, 10, 20) // when length and capacity is different

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice)) // intSlice        Len: 10         Cap: 10
	fmt.Println(reflect.ValueOf(intSlice).Kind())                              // slice

	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice)) // strSlice        Len: 10         Cap: 20
	fmt.Println(reflect.ValueOf(strSlice).Kind())                              // slice
}

func declareSliceUsingNewKey() {
	var intSlice = new([50]int)[0:10]

	fmt.Println(reflect.ValueOf(intSlice).Kind())                              //slice
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice)) // intSlice        Len: 10         Cap: 50
	fmt.Println(intSlice)                                                      // [0 0 0 0 0 0 0 0 0 0]

}

func usingLiteralSlice() {
	var intSlice = []int{10, 20, 30, 40, 50}
	var strSlice = []string{"India", "Canada", "Japan"}

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
}

func additems() {
	a := make([]int, 2, 5)

	a[0] = 10
	a[1] = 20

	fmt.Println("Slice A:", a)                                  //Slice A: [10 20]
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a)) //Length is 2 Capacity is 5

	a = append(a, 30, 40, 50, 60, 70, 80, 90)

	fmt.Println("Slice A after appending data:", a)             // Slice A after appending data: [10 20 30 40 50 60 70 80 90]
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a)) // Length is 9 Capacity is 12

}

func removeFromSlice() {
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Println(strSlice)

	strSlice = removeIndex(strSlice, 3)

	fmt.Println(strSlice)

}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func copySlice() {
	a := []int{5, 6, 7}

	fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a), cap(a))

	b := make([]int, 5, 10)
	copy(b, a)
	fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))

	fmt.Println("Slice B after copying:", b)
	b[3] = 8
	b[4] = 9
	fmt.Println("Slice B after adding elements:", b)
}

func tricksOfSlicing() {
	var countries = []string{"india", "japan", "canada", "australia", "russia"}

	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(":2 %v\n", countries[:2])

	fmt.Printf("1:3 %v\n", countries[1:3])

	fmt.Printf("2: %v\n", countries[2:])

	fmt.Printf("2:5 %v\n", countries[2:5])

	fmt.Printf("0:3 %v\n", countries[0:3])

	fmt.Printf("Last element: %v\n", countries[4])
	fmt.Printf("Last element: %v\n", countries[len(countries)-1])
	fmt.Printf("Last element: %v\n", countries[4:])

	fmt.Printf("All elements: %v\n", countries[0:len(countries)])

	fmt.Printf("Last two elements: %v\n", countries[3:len(countries)])
	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])

	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])

}

func loopThroughASlice() {
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Printf("\n---------------Example 1 --------------------\n")
	for index, element := range strSlice {
		fmt.Println(index, "--", element)
	}

	fmt.Printf("\n---------------Example 2 --------------------\n")

	for _, value := range strSlice {
		fmt.Println(value)
	}

	fmt.Printf("\n---------------Example 3 --------------------\n")
	j := 0
	for range strSlice {
		fmt.Println(strSlice[j])
		j++
	}

}

func appendAScliceToAnotherSlice() {
	var slice1 = []string{"india", "japan", "canada"}
	var slice2 = []string{"australia", "russia"}

	slice2 = append(slice2, slice1...)
}

func checkIfitemExists() {
	strString := []string{"India", "Canada", "Japan", "Germany", "Italy"}
	itemExists(strString, "India")
}

func itemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func TestSlices() {
	emptySlice()
	declareSliceUsingMake()
	declareSliceUsingNewKey()
	usingLiteralSlice()
	additems()
	removeFromSlice()
	checkIfitemExists()
}
