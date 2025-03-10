package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("Inside main without go routine")

	for i := 0; i < 3; i++ {
		// i assign to i again still work as well
		i := i // i shadow i
		go func() {
			fmt.Println("For loop i: ", i)
		}()
		// go func() {
		// 	fmt.Println("For loop i: ", i)
		// }()

		// parameter approach still works
		// go func(i int) {
		// 	fmt.Println("For loop i: ", i)
		// }(i)
	}

	time.Sleep(10 * time.Millisecond)
	// main go routine terminate before other go routine finishes

	ch := make(chan string)
	go func() {
		ch <- "hi" // send
	}()
	msg := <-ch // receive
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("got : ", msg)
	}

	// receive from the closed channel
	msg, ok := <-ch // ch is closed
	fmt.Printf("closed: %#v (ok=%v)\n", msg, ok)
	// getting zero value without blocking

}
