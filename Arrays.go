/*
Storing a number of elements in a list can be achieved using arrays, slices, and maps (like hash-maps but in Go). All three are in the example below. 
Arrays are defined by their fixed size and a common data type for all elements. The size of the array is part of the type, thus arrays cannot grow or shrink, otherwise, they would have a different type. 
Array elements are accessed using square brackets. 
The example below shows how to declare an array containing strings and how to loop through its elements.
*/

package main

import "fmt"

func main() {
	/* Define an array of size 4 that stores deployment options */
	var DeploymentOptions = [4]string{"R-pi", "AWS", "GCP", "Azure"}

	/* Loop through the deployment options array */
  
  //Notice the lack of parentheses around the looping condition. Here, we traverse the array outputting the current index and the value stored at that index. 
	for i := 0; i < len(DeploymentOptions); i++ {
		option := DeploymentOptions[i]
		fmt.Println(i, option)
	}
}

/*
# output
0 R-pi
1 AWS
2 GCP
3 Azure
*/
