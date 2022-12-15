package statements

import (
	"fmt"
	"time"
)

func Switch() {
	today := time.Now()

	switch today.Day() {
	case today.Day():
		fmt.Println("Today is 5th. clean your house.")
		fallthrough

	case 10, 15:
		fmt.Println("Today is 10th. buy some wine.")
		fallthrough

	case 20:
		fmt.Println("Today is 5th.Visit a doctor.")

	default:
		fmt.Println("No information available for that day.")
	}
}

func SwitchWithInitializer() {
	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("Today is 5th. clean your house.")
		fallthrough

	case today.Day() >= 10:
		fmt.Println("Today is 10th. buy some wine.")
		fallthrough

	case today.Day() == 20:
		fmt.Println("Today is 5th.Visit a doctor.")

	default:
		fmt.Println("No information available for that day.")
	}
}
