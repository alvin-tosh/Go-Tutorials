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

/*
First, we create a channel that will be common to all goroutines. Then we start 5 goroutines and pass the channel as an argument. 
Inside each goroutine, we write the gopher id to standard output as soon the gopher starts cooking the dish. We then send the gopher id from the goroutine back to the caller. 
From there, we’re back to the body of the main function where we receive the gopher id and record their finishing time.
*/
func main() {
	c := make(chan int) 		// Create a channel to pass ints
	for i := 0; i < 5; i++ {
		go cookingGopher(i, c)  // Start a goroutine
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

/*
Since we’re dealing with concurrent code, we lose the ability to predict the order of the output, however, we can observe how the channel blocks the execution, as a goroutine has to wait until the channel is available before it can send an id. 
One POSIBLE output is included below.
Keep in mind that we’re probably using more goroutines than the number of cores on our machine, hence it’s likely that a single core is time-multiplexed to simulate the concurrency.
*/

/*
# Output
gopher 0 started cooking
gopher 4 started cooking
gopher 3 started cooking
gopher 0 finished the dish
gopher 2 started cooking
gopher 1 started cooking
gopher 4 finished the dish
gopher 3 finished the dish
gopher 2 finished the dish
gopher 1 finished the dish

*/
