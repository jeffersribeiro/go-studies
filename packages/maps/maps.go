package maps

import (
	"fmt"
	"sort"
)

/*
Golang Maps

in this tutorial you will learn what is a map data type and when to use it in Golang.

A map os a data structure that provides you with an unordered collection of  key/value pair(maps are also
sometimes called associative arrays in Php, hash tables in Java, or dictionaries in Python). Maps are used to
look up a value by its associated key. you store values into the maps based on a key.

The strngth of a map is its ability to retrive data quickly based on the key. A key works like an index,
pointing to the value you associate with that key.

A maps is implemented using a hash table, which is providing
*/

func mapInit() {
	var employee = map[string]int{"mark": 10, "Sandy": 20}

	fmt.Println(employee)
}

func emptyMapDeclaration() {
	var emplyee = map[string]int{}
	fmt.Println(emplyee)
	fmt.Printf("%T\n", emplyee)

}

func declareUsingMake() {
	var employee = make(map[string]int)

	employee["Mark"] = 10
	employee["Sandy"] = 20

	fmt.Println(employee)

	employeeList := make(map[string]int)

	employeeList["Mark"] = 10
	employeeList["Sandy"] = 20

	fmt.Println(employeeList)

}

func accessingItems() {
	var employee = map[string]int{"Mark": 10}

	fmt.Println(employee["Mark"])
}

func deleteItems() {
	var employee = make(map[string]int)
	employee["Mark"] = 20
	employee["Sandy"] = 30
	employee["Rocky"] = 40
	employee["Josef"] = 50

	fmt.Println(employee)

	delete(employee, "Mark")
	fmt.Println(employee)
}

func iterateOverMap() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50}

	for key, element := range employee {
		fmt.Println("key: ", key, "=>", "element: ", element)
	}

}

func truncateMap() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50}

	// Method - I
	for k := range employee {
		delete(employee, k)
	}

	// Method - II
	employee = make(map[string]int)
}

func sortMapKeys() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	keys := make([]string, 0, len(unSortedMap))

	for k := range unSortedMap {
		keys = append(keys, k)

	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, unSortedMap[k])
	}
}

func sortMapValues() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	values := make([]int, 0, len(unSortedMap))

	for _, v := range unSortedMap {
		values = append(values, v)
	}

	sort.Ints(values)

	for _, v := range values {
		fmt.Println(v)
	}
}

func mergeMaps() {
	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4}

	for k, v := range second {
		first[k] = v
	}

	fmt.Println(first)
}

func TestMaps() {
	mapInit()
	emptyMapDeclaration()
	declareUsingMake()
	accessingItems()
	deleteItems()
	iterateOverMap()
	truncateMap()
	sortMapKeys()
	sortMapValues()
	mergeMaps()
}
