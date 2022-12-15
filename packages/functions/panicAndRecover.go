package functions

import (
	"fmt"

	"errors"
)

var result = 1

func chain(n int) {
	if n == 0 {
		panic(errors.New("cannot multiply a number by zero"))
	} else {
		result *= n
		fmt.Println("Output: ", result)
	}
}

func chainWithRecover(n int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if n == 0 {
		panic(errors.New("cannot multiply a number by zero"))
	} else {
		result *= n
		fmt.Println("Output: ", result)
	}
}

func TestPanicAndRecover() {
	chainWithRecover(5)
	chainWithRecover(2)
	chainWithRecover(0)
	chainWithRecover(8)

}
