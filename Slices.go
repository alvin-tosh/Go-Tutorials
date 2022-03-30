/*
Slices can be thought of as dynamic arrays. Slices always refer to an underlying array and can grow when new elements are added. 
The number of elements that are visible through a slice determines its length. 
If a slice has an underlying array that is larger, the slice may still have the capacity to grow. 
When it comes to slices, think of the LENGTH as the current number of elements, and think of the CAPACITY as the maximum number of elements that can be stored. 
*/

package main

import (
    "fmt"
    "reflect"
)

func main() {
	/* Define an array containing programming languages */
	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust",   // Must include the trailing comma
	}

	/* Define slices */
	classics := languages[0:3]  // alternatively languages[:3]
	modern := make([]string, 4) // len(modern) = 4
	modern = languages[3:7]     // include 3 exclude 7
	new := languages[7:9]       // alternatively languages[7:]
  
  fmt.Printf("classic languagues: %v\n", classics) // classic languagues: [C Lisp C++]
	fmt.Printf("modern languages: %v\n", modern)     // modern languages: [Java Python JavaScript Ruby]
	fmt.Printf("new languages: %v\n", new)           // new languages: [Go Rust]
  
  allLangs := languages[:]                             // copy of the array
        fmt.Println(reflect.TypeOf(allLangs).Kind())   // slice

        /* Create a slice containing web frameworks */
        frameworks := []string{
            "React", "Vue", "Angular", "Svelte",
            "Laravel", "Django", "Flask", "Fiber",
        }

        jsFrameworks := frameworks[0:4:4]          // length 4 capacity 4
        frameworks = append(frameworks, "Meteor")  // not possible with arrays

        fmt.Printf("all frameworks: %v\n", frameworks)
        fmt.Printf("js frameworks: %v\n", jsFrameworks)
}

