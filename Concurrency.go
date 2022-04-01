/*
Things get interesting when multiple threads are run in parallel so that a program can make use of multiple CPU cores. 
Goroutines(Go’s version of threads) are initiated using the go keyword. In addition to goroutines, Go has built-in channels which are used to share data between goroutines. 
In general, send and receive operations across a channel block the execution until the other side is ready.

In the example below, we’ll consider 5 goroutines that run in parallel. Let’s suppose we organize a cooking contest between 5 gofer chefs. 
This is a timed contest and whoever finishes their dish first wins. Let’s see how we can simulate this contest using Go’s concurrency features.

*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // Create a channel to pass ints
	for i := 0; i < 5; i++ {
		go cookingGopher(i, c) // Start a goroutine
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c // Receive a value from a channel
		fmt.Println("gopher", gopherID, "finished the dish")
	} // All goroutines are finished at this point
}

/* Notice the channel as an argument */
func cookingGopher(id int, c chan int) {
	fmt.Println("gopher", id, "started cooking")
	c <- id // Send a value back to main
}
