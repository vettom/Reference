package main

/*
When defining channels, must pass value to channel for it to act on.
Channel has 2 part, sender and receiver.
c <- "Sending some data"
<- c "Receiving data"

Channel is a blocking code, meaning it will wait until data is received.
For each routine, there must be a receiver for channel, if not it will exist on first available.
*/
import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://facebook.com",
		"https://google.co.uk",
		"http://faildomain.com",
	}

	// Crete a channel so that go routing wait for child process to complete
	c := make(chan string)

	// Loop through links and check status
	for _, link := range links {
		// Calling with routine but there is nothing to return output
		// Pass channel also as argument.
		go checkLink(link, c)
	}
	// At end of routines, print channel output
	// Channel receive function is blocking code and required per routine
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	// Dont care about data, only checking status code
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error:", link, " may be down")
		// Send data to channel
		c <- "Site is down"
		return
	}
	// If success then return result
	fmt.Println("INFO: success", link, " is working")
	c <- "Site is up and running"

}
