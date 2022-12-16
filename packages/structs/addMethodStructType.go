package structs

import "fmt"

func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)

	for _, salary := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(salary.Basic)
		fmt.Println(salary.HRA)
		fmt.Println(salary.TA)
	}

	return "----------------------"

}

func AddMethodStructType() {
	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "Mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			}, {
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			}, {
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	fmt.Println(e.EmpInfo())
}
