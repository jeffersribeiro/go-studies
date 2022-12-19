package interfaces

import "fmt"

/*
  Two or more interfaces can have one or more common method in lis of method sets. Here,
  Structure is a common method between two interfaces Vehicle and Human
*/

type Vehicle interface {
	Structure() []string
	Speed() string
}

type Human interface {
	Structure() []string
	Perfomance() string
}

type Car string

func (c Car) Structure() []string {
	var parts = []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Task"}
	return parts
}

func (c Car) Speed() string {
	return "200 Km/Hrs"
}

type Man string

func (m Man) Structure() []string {
	var parts = []string{"Brain", "Heart", "Nose", "Eyelashes", "Stomach"}
	return parts
}

func (m Man) Perfomance() string {
	return "8 Hrs/Day"
}

func InterfaceswithcommonMethod() {
	var bmw Vehicle
	bmw = Car("World Top Brand")

	var labour Human
	labour = Man("Software Developer")

	for i, j := range bmw.Structure() {
		fmt.Printf("%-15s <=====> %15s\n", j, labour.Structure()[i])
	}

}
