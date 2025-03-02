package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("Inside main without go routine")

	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println("For loop i: ", i)
		}()
	}

	time.Sleep(10 * time.Millisecond)
	// main go routine terminate before other go routine finishes
}
