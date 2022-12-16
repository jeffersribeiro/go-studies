package structs

import (
	"encoding/json"
	"fmt"
)

type Employee2 struct {
	FirstName string `json:"fistname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

/*
During the definition of a Struct type. optional String values may be added to each field declaration.

The tags are represented as raw string values (wrapped within a pair of `<code_string_inside>`) and ignored by normal code execution
*/
func UseFieldTags() {
	json_string := `
	{
		"lastname":"Rocky",
		"lastname": "Sting",
		"city": "London"
	}`

	emp1 := new(Employee2)
	json.Unmarshal([]byte(json_string), emp1)
	fmt.Println(emp1)

	emp2 := new(Employee2)

	emp2.FirstName = "Ramesh"
	emp2.LastName = "Soni"
	emp2.City = "Munbai"

	jsonStr, _ := json.Marshal(emp2)

	fmt.Printf("%s\n", jsonStr)
}
