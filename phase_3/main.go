package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")
		}

		// There could be cases where this message is not completed!
		// This is solved in future examples
		fmt.Println(time.Now(), "all completed!")
	}()

	time.Sleep(time.Second * 2)

	fmt.Println("waiting for messages")
	fmt.Println(time.Now(), "receieved", <-ch)
	fmt.Println(time.Now(), "receieved", <-ch)
	fmt.Println(time.Now(), "receieved", <-ch)
	fmt.Println("exiting!")
}
