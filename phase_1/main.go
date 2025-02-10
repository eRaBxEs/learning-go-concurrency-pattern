package main

import (
	"fmt"
)

// This program panics because there is no goroutine
// outside of main interacting with the channel
//fatal error: all goroutine are asleep

func main() {
	var ch chan int
	ch = make(chan int)
	// the two lines above can be replaced with
	// ch := make(chan int)

	ch <- 10

	v := <-ch
	fmt.Println("received:", v)
}
