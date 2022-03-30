/*
 A struct is simply a collection of fields. 
 In the next example, we’ll use what we’ve learned about pointers, learn how to use a struct, and build a stack from scratch.
*/

package main

import "fmt"

/* Define a stack type using a struct...
To achieve the stack functionality, we need an array to store the stack elements, and an index to point to the last item in the stack. 
For the sake of the example, let’s fix our stack size to 5 elements. 
Inside the body of the struct, we specify an index field which is of type int, and a field called data, which is an array of 5 int elements.
*/

type stack struct {
	index int
	data  [5]int
}

/* Define push and pop methods...
Next we define the push and pop methods. A method is a special kind of function that takes a receiver argument between the func keyword and the method name. 
Notice the type of the parameter s. In this case, it is a stack pointer instead of a stack.
*/
func (s *stack) push(k int) {
	s.data[s.index] = k
	s.index++
}

/* Notice the stack pointer s passed as an argument */
func (s *stack) pop() int {
	s.index--
	return s.data[s.index]
}

/*
In the body of our stack methods, we access the stack fields using the dot notation. In the push method, we write a given integer k to the first available index (recall the default value of a declared int is 0), and increment the index by 1. 
In the pop method we decrement the index by 1, and return the last item in the stack. 
In the body of the main function, we use new() to create a pointer to a newly allocated stack. We then push 2 items and write the result to standard output.
*/
func main() {
	/* Create a pointer to the new stack and push 2 values */
	s := new(stack)
	s.push(23)
	s.push(14)
	fmt.Printf("stack: %v\n", *s)
}

/*
# Output

stack: {2 [23 14 0 0 0]}

*/
