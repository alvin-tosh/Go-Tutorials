package main

import "fmt"

func main() {
  
        /* User specified types */
        const w int32 = 150          // 32-bit integer
        const x float32 = 50.5       // 32-bit float
        var y complex128 = 1 + 4i    // 128-bit complex number
        var z uint16 = 18            // 16-bit unsigned integer

        /* Default types */
        n := 35                                  // int
        pi := 3.14                               // float64
        p, r := true, false                      // bool
        h := "Learning types in Go is AWESOME!"  // string

        fmt.Printf("user-specified types:\n %T %T %T %T\n", w, x, y, z)
        fmt.Printf("default types:\n %T %T %T %T %T\n", n, pi, p, r, h)
}
