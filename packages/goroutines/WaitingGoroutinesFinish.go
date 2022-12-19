package goroutines

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

/*
Waiting for goroutines to finish execution

the waitGroup type of Sync package, us used to wait the program to finish all goroutines launched from
the main function. it uses a counter that specifies the number of gogroutines, and wait the execution of
the program until the WaitGroup counter is zero.

The Add method is used to add a counter to the WaitGroup

The Done method os WaitGroup is scheduled using a defer statement to decrement the WaitGroup counter.

The Wait method of the WaitGroup type waits for the program to finish all goroutines.

The wait method is called inside the mains function, which block execution until the WaitGroup counter
reached the value of zero and ensures that all goroutines are executed.
*/

var wg sync.WaitGroup

func responseSize2(url string) {

	defer wg.Done()

	fmt.Println("Step1: ", url)
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)

	defer response.Body.Close()

	fmt.Println("Step3: ", url)

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step 4: ", len(body))
}

func WaitingGoroutinesFinish() {

	wg.Add(3)

	fmt.Println("Start Goroutines")

	go responseSize2("https://www.golangprograms.com")
	go responseSize2("https://www.stackoverflow.com")
	go responseSize2("https://www.codewall.com")

	wg.Wait()
	fmt.Println("terminating Program")
}

/*
Output:

Start Goroutines
Step1:  https://coderwall.com
Step1:  https://www.golangprograms.com
Step1:  https://stackoverflow.com
Step2:  https://stackoverflow.com
Step3:  https://stackoverflow.com
Step4:  116749
Step2:  https://www.golangprograms.com
Step3:  https://www.golangprograms.com
Step4:  79801
Step2:  https://coderwall.com
Step3:  https://coderwall.com
Step4:  203842
Terminating Program
*/
