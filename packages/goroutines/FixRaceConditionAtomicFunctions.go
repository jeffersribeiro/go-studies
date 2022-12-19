package goroutines

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

var (
	counter int32
)

func FixRaceConditionAtomicFunctions() {
	wg.Add(3)

	go increment("Python")
	go increment("Java")
	go increment("Golang")

	wg.Wait()
	fmt.Println("Counter", counter)
}

func increment(name string) {
	defer wg.Done()

	for range name {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched()
	}
}

/*
 The AddInt32 function from the atomic package synchronizes the adding of integer
 values by enforcing that only one goroutine can perform and complete this add
 operation at a time. When goroutines attempt to call any atomic function, they're
 automatically synchronized against the variable that's referenced.
*/

/*
Note if you replace the code line atomic.AddInt32(&counter, 1) with
counter++, then you will see ther below output-
run go -race main.go

WARNING: DATA RACE
Read at 0x0000006072b0 by goroutine 7:
  main.increment()
      C:/Golang/goroutines/main.go:31 +0x76

Previous write at 0x0000006072b0 by goroutine 8:
  main.increment()
      C:/Golang/goroutines/main.go:31 +0x90

Goroutine 7 (running) created at:
  main.main()
      C:/Golang/goroutines/main.go:18 +0x7e

Goroutine 8 (running) created at:
  main.main()
      C:/Golang/goroutines/main.go:19 +0x96
==================
Counter: 15
Found 1 data race(s)
exit status 66

*/
