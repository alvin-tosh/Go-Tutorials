/*
Numerical types are the most versatile, with 8, 16, 32, and 64-bit variants for both signed (int) and unsigned (uint) integers.

A byte is an alias for uint8. A rune is an alias for int32. Floats (or floating-point numbers) are either float32 or float64. 
Complex numbers are also supported and can be represented as complex128 or complex64.
When a variable is declared it is assigned to the natural “null” value of the corresponding type. 
For example, in var x int, x has the value 0. In var y string, y has the value "“. 
The example below shows the difference between user-specified types and the default types assigned with a short variable declaration.
*/
package main

import "fmt"

func main() {
  
        // User specified types 
        const w int32 = 150          // 32-bit integer
        const x float32 = 50.5       // 32-bit float
        var y complex128 = 1 + 4i    // 128-bit complex number
        var z uint16 = 18            // 16-bit unsigned integer

        // Default types 
        n := 35                                  // int
        pi := 3.14                               // float64
        p, r := true, false                      // bool
        h := "Learning types in Go is AWESOME!"  // string

        fmt.Printf("user-specified types:\n %T %T %T %T\n", w, x, y, z)
        fmt.Printf("default types:\n %T %T %T %T %T\n", n, pi, p, r, h)
}
