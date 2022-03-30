/*
Most modern programming languages have a built-in implementation of a hash-map. For example, think of Python’s dictionary or JavaScript’s object. 
Fundamentally, a map is a data structure that stores key-value pairs with a constant look-up time. 
The efficiency of maps comes at the expense of randomizing the order of the keys and the associated values. 
In other words, we make no guarantees about the order of the elements in a map. The example below showcases this behavior.
*/

package main

import "fmt"

/*
We define a map called firstReleases containing several programming languages as the keys, and their release years as the corresponding values. 
We also write a loop to traverse the map and output each key-value pair. 
If we run the code, notice the random order of the elements displayed in the output.
*/
func main() {
  
	/* Define a map containing the release year of several languages */
	firstReleases := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "JavaScript": 1996, "Go": 2012,
	}

	/* Loop through each entry and output the name and release year */
	for k, v := range firstReleases {
		fmt.Printf("%s was first released in %d\n", k, v)
	}
}

/*
# output
Go was first released in 2012
C was first released in 1972
C++ was first released in 1985
Java was first released in 1996
Python was first released in 1991
JavaScript was first released in 1996

*/
