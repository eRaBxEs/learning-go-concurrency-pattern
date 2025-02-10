package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	exit := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")

			time.Sleep(time.Second * 1)
		}

		fmt.Println(time.Now(), "all completed, leaving")

		close(ch)
	}()

	// Below is the consumer
	go func() {
		// N.B This is overcomplicated becos its a only one channel,
		// select shines when using multiple channels
		// for {
		// 	select {
		// 	case v, open := <-ch:
		// 		if !open {
		// 			close(exit)
		// 			return
		// 		}
		// 		fmt.Println(time.Now(), "received", v)

		// 		// default:
		// 		// 	fmt.Println("Nothing is happening")
		// 	}
		// }
		// using range: in cases where only one channel is used
		for v := range ch {
			fmt.Println(time.Now(), v)
		}

		close(exit)
	}()

	fmt.Println(time.Now(), "waiting for everything to complete")

	<-exit

	fmt.Println(time.Now(), "exiting")
}
