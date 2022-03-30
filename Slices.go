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
	
/*
We make a copy of the languages array using the [:] operator. The resulting copy is a slice. We assert that’s the case using the "reflect" package. 
Next, we create a slice called frameworks. Notice the blank entry in the square brackets responsible for the size. If we pass a parameter inside these brackets we are creating an array. Leaving it blank creates a slice. 
From there, we create another slice called jsFrameworks that selects JavaScript frameworks. 
Finally, we extend our frameworks slice by adding Meteor to the list of frameworks.

The 'append' function pushes new values to the end of a slice and returns a new slice with the same type as the original. 
In case the capacity of a slice is insufficient to store the new element, a new slice is created that can fit all the elements. 
In that case, the returned slice will refer to a different underlying array. 
*/	
  
	allLangs := languages[:]                             // copy of the array made earlier in Line:17
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

/*
# output
...
all frameworks: [React Vue Angular Svelte Laravel Django Flask Fiber Meteor]
js frameworks: [React Vue Angular Svelte]
*/

