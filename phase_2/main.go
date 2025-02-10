package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println(time.Now(), "taking a nap!")
		time.Sleep(time.Second * 2)
		ch <- "hello"

	}()

	fmt.Println("Waiting for message!")
	v := <-ch
	fmt.Println("received:", v)
}
