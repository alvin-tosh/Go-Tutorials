package main

import "fmt"

// Declare a single variable 
var x int

// Declare a group of variables
var (
    y bool
    z float32
    p string
)

func main() {
	x = 20                                                // Assign single value
	y, z = true, 75.0                                     // Assign multiple values
  p = "Follow me on Twitter @_Lord_Heathen"             // Strings must contain double quotes
	fmt.Println(x, y, z, p)                               // 20 true 75 Follow me on Twitter @_Lord_Heathen
}
