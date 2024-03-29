package main

import (
	"fmt"
	"time"
)

/*
	Using this select timeout pattern requires communicating results over channels.
	This is a good idea in general because other important Go features are based on
	channels and select.

*/
func timeout_example() {
	// Channel + goroutine
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	// Here’s the select implementing a timeout. res := <-c1 awaits the result and <-Time.
	// After awaits a value to be sent after the timeout of 1s. Since select proceeds with
	// the first receive that’s ready, we’ll take the timeout case if the operation takes
	// more than the allowed 1s.

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	// If we allow a longer timeout of 3s, then the receive from c2 will succeed and we’ll print the result.

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
