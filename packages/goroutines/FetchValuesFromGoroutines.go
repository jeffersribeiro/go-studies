package goroutines

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

/*
The most natural way to fetch a value from a goroutine is channels, Channels are the pipes that connect
concurrent goroutines. You can send values into channels from one goroutine and receive those values into
another goroutines or in a synchoronous function.
*/

var wg2 sync.WaitGroup

func responseSize3(url string, nums chan int) {
	defer wg2.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Send value to the unbuffered channel
	nums <- len(body)
}

func FetchValuesFromGoroutines() {
	nums := make(chan int) // Declare a unbuffered channel

	wg2.Add(1)

	go responseSize3("https://golangprograms.com", nums)

	fmt.Println(<-nums) // read the value from unbuffered channel
	wg2.Wait()

	close(nums) // close the channel
}
