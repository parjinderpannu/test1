package main

import "time"

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- 2
	}()
}
