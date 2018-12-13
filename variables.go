/* Packages and Variables */

package main 

import(

	//importing multiple packages at once
	
	"fmt"
	"math"
)

// Declaration and Assignment of variables outside functions
var i int
var j int, k string 
var m int = 10

// For assignment, both the values should be mentioned on RHS
var n, p int = 10, 20 

// These declaration methods can be used in a function too
var c, d, e = 1, 2.93, "String!" 


func main(){
	fmt.Println(i, m, n, p, c, d, e)

	new_variable := "Hi, I'm the new member"
	fmt.Println(new_variable)

	// When importing a package, you can refer only to its exported names which start with a capital letter
	fmt.Println(math.Pi)
}