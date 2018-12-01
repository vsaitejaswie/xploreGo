/* Pointers */

package main 

import "fmt" 

// Unlike C, Go has no pointer arithmetic
var ptr *int
var p = 10

var dlt *string
var str string
func main(){
	ptr = &p 
	// A pointer can be initialised only inside a function --- [why???]

	fmt.Println(ptr)
	fmt.Println("Address of 'ptr' pointer is", ptr)

	fmt.Println("The default value of a pointer is", dlt)
	fmt.Println("The default value of a string is ", str) // Blank Space
	// <nil> can cause Panic in the control flow

	str = "String"
	dlt = &str
	fmt.Println("The assigned string is", *dlt)
}

// Refer : https://dave.cheney.net/2014/03/17/pointers-in-go
// Refer : https://blog.golang.org/defer-panic-and-recover