package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("Inside main without go routine")

	time.Sleep(10 * time.Millisecond)
	// main go routine terminate before other go routine finishes
}
