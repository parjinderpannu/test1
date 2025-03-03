package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it finishes
	fmt.Printf("Worker %d: Starting\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d: Done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Launch 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)         // Increment the counter
		go worker(i, &wg) // Launch goroutine
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers completed")
}
