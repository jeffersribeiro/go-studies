package interfaces

import "fmt"

/*
An intereface is an abstract type.

interface describes all the method set adn provides the signarture for each method.

to create interface use interface keyword, followed by curly braces containing a list of emthod names,
along with any parameters or return values the methods are expected to have.
*/

type Address struct {
	Id     int
	Street string
}

type Employee interface {
	PrintName(name string)
	PrintAddress(id int) Address        // Method with int parameter
	PrintSalary(basic int, tax int) int // Method with parameters and return type
}

type Emp int

func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id: \t", e)
	fmt.Println("Employee Name: \t", name)
}

func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

func (e Emp) PrintAddress(id int) Address {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var addresses = []Address{
		{
			Id:     1,
			Street: "Beaver St",
		},
		{
			Id:     2,
			Street: "Hanover St",
		},
	}

	var j int

	for index, address := range addresses {
		if address.Id == id {
			j = index
			break
		} else {
			panic("Not Found, change id and try again")
		}

	}
	return addresses[j]
}

func PrintMethods() {
	e1 := Emp(1)
	e1.PrintName("John Doe")
	fmt.Println("Employee Address: ", e1.PrintAddress(5))
	fmt.Println("Employee Salary: ", e1.PrintSalary(25000, 5))
}
