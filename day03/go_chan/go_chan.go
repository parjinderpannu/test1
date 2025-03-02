package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		// // Fix 2 : using loop body
		i := i
		go func() {
			fmt.Println(i) // read from line 14
		}()
		// // Fix 1 : Using parameter
		// go func(n int) {
		// 	fmt.Println(n)
		// }(i)

		// Earlier it was bug but looks like it is fixed in latest go version
		// go func() {
		// 	fmt.Println(i)
		// }()
	}

	ch := make(chan string)
	go func() {
		ch <- "hi" // send
	}()
	msg := <-ch // receive
	fmt.Println(msg)
	/*
		- send & receive will block until opposite operation (*)
		- receiving from closed channel will return the zero value without blocking
		- send to a closed channel will panic
		- closing a closed channel will panic
		- send/receive to a nil channel will block forever
	*/

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch) // this fixes the issue of deadlock
	}()
	// go doesn't know how many messages are going to come
	for msg := range ch {
		fmt.Println("got:", msg)
	}

	/* for/range does this
	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("got:", msg)
	}
	*/

	msg = <-ch // ch is closed
	fmt.Printf("closed: %#v\n", msg)

	msg, ok := <-ch // ch is closed
	fmt.Printf("closed: %#v (ok:%v)\n", msg, ok)

	// go runtime doesn't wait for go routine other than main go routine
	time.Sleep(10 * time.Millisecond)

	// shadowExample()
}

func shadowExample() {
	n := 7
	{
		n := 2 // from here to } this is 'n'
		//n = 2 // this is not shadowing
		fmt.Println("Inner :", n)
	}
	fmt.Println("Outer :", n)
}
