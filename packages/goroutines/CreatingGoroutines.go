package goroutines

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
Goroutines

Concurrency in Golang is the ability for functions to run independent of each other. goroutines are functions
that are run concurrently. Golan provides Goroutines as a way to handle operations concurrently

New goroutines are created by the go statement

top run a function as a goroutine, call thta function prefixed with go statement. Here is the example code block:

Example:

sum()    // A normal function call that executes sium synchronously and waits for completing it
go sum() // A goroutine thta executes sum  asynchronously and doesn't wait for completing it

The go keyword makes the function call to return immediately, While the function starts running in the
background as a goroutine and rest of the program continues its execution. the main function of every
Golang program is starts using a goroutine, so every golang progrma runs at least one goroutine.
*/

/*
Added the go keyword before each call function responseSize. the three responseSize
*/

func responseSize(url string) {
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

	fmt.Println("Step4: ", len(body))
}

func CreatingGoroutines() {
	go responseSize("https:www/golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	time.Sleep(10 * time.Second)
}

/*
  Output:

  Step1:  https://www.golangprograms.com
  Step1:  https://stackoverflow.com
  Step1:  https://coderwall.com
  Step2:  https://stackoverflow.com
  Step3:  https://stackoverflow.com
  Step4:  116749
  Step2:  https://www.golangprograms.com
  Step3:  https://www.golangprograms.com
  Step4:  79551
  Step2:  https://coderwall.com
  Step3:  https://coderwall.com
  Step4:  203842
*/
