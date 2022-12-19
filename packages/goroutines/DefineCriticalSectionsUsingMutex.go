/*
 A mutex is used to create a critical section around code thta ensures
 only one goroutine at a time can execure thta code section
*/

package goroutines

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
)

func increment2(lang string) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		mutex.Lock()
		{
			fmt.Println(lang)
			counter++
		}
		mutex.Unlock()
	}
}

func DefineCriticalSectionsUsingMutex() {
	wg.Add(3)
	go increment2("Python")
	go increment2("Go Programming Language")
	go increment2("Java")

	wg.Wait()

	fmt.Println("Counter", counter)
}

/*
Output:

Python
Python
Python
Java
Java
Java
Go Programming Language
Go Programming Language
Go Programming Language
Counter 9
*/

/*
 A critical section defined by tge calls to Loack() and unlock()
 protects the acrtions against the counter variable and reading the text of name variable.
*/
