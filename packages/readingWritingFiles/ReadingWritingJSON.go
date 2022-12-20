package reading_files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CatlogNodes struct {
	CatlogNodes []Catlog `json:"catlog_nodes"`
}

type Catlog struct {
	Product_id string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

func ReadingJSON() {
	file, _ := ioutil.ReadFile("test.json")

	data := CatlogNodes{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.CatlogNodes); i++ {
		fmt.Println("Product Id: ", data.CatlogNodes[i].Product_id)
		fmt.Println("Quantity: ", data.CatlogNodes[i].Quantity)
	}

}

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func WritingJSON() {

	data := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	file, _ := json.MarshalIndent(data, "", "  ")

	_ = ioutil.WriteFile("test.json", file, 0644)
}
