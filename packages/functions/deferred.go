package functions

import "fmt"

// Deferred Functions Calls
// Go has a special statement called defer that schedules a function call to be run after the function completes.
// consider the folling example:

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func third() {
	fmt.Println("Third")
}

type File struct {
	Open  func(name string)
	Close func()
}

var (
	file     = new(File)
	failureX = true
	failureY = true
)

// Without defer
func ReadWriteWithoutDefer() bool {
	file.Open("file")

	if failureX {
		file.Close() // And here
		return false
	}

	if failureY {
		file.Close() // And here
	}

	file.Close() // And here
	return true
}

// With defer
func ReadWriteWithDefer() bool {
	file.Open("file")

	defer file.Close() // And here

	if failureX {
		return false
	}

	if failureY {
		return false
	}

	return true
}

func TestDeferred() {
	defer third()
	first()
	second()
}

/*
 * A defer statement is often used with paired operatrions like open and close,
 * connect and disconnect, or lock and unlock to ensure that resorces are released in all cases,
 * no matter how complex the control flow. The right place for a defer statement that releases
 * a resouce is immeadiately after the resouce has benn sucessfuly acquired.
 *
 * Below is the example to opne a file perform read/write action on it.
 * in thes example there are often spots where you want to return early
 */
