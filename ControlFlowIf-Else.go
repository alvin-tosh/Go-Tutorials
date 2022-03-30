/*
Let’s suppose we’re given a slice containing floats, and we’re interested in computing their average value. 
We’ll proceed by creating a function called average that takes a slice as a parameter and returns a float called avg. 
The example below shows a possible implementation using If-Else statement.
*/

package main

import "fmt"

/* Define a function to find the average of the floats contained in a slice */
func average(x []float64) (avg float64) {
	total := 0.0
	if len(x) == 0 {
		avg = 0
	} else {
		for _, v := range x {
			total += v
		}
		avg = total / float64(len(x))
	}
	return
}

func main() {
	x := []float64{2.15, 3.14, 42.0, 29.5}
	fmt.Println(average(x))   // 19.197499999999998
}
/*
#Output

19.197499999999998

*/
