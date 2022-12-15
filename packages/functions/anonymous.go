package functions

import "fmt"

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func Anonymous() {
	func(name string) string {
		fmt.Println(`My name is` + name)
		return `My name is` + name
	}("Jefferson")

	fmt.Println(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)

	fmt.Println(area(20, 30))
}
