/* Packages, Variables and Functions */

package main 

import(

	//importing multiple packages at once
	
	"fmt"
	"math"
 )

var i int
var j, k string // Declaration and Assignment of variables outside functions
var m int = 10
var n, p int = 10, 20 // For assignment, both the values should be mentioned on RHS
var c, d, e = 1, 2, 3 // These declaration methods can be used in a function too

// A function can return any number of results
func multi_return(a int, b string) (string, int, int){
	a = 100
	b = "TejVR" // Assignment of variables within a function
	return b, a, (a+10)
}

// Named return values
// Naked return
func try_named()(name1 string, name2 int){
	name2 = 5
	name1 = "named Return"
	return
}

func main(){
	// A name is exported if it begins with a capital letter
	// When importing a package, you can refer only to its exported names.
	// Here Println starts with a capital letter, hence it is an exported name
	fmt.Println("This is module 1")
	fmt.Println(math.Pi)
	new_variable := "Hi, I'm the new member"
	fmt.Println(new_variable)
	x, y, z := multi_return(10, "Tej") // Declaration and Assignment of new variables within a function
	fmt.Println(x, y, z)
	fmt.Println(i, m, n, p, c, d, e)
	a1, a2 := try_named()
	fmt.Println(a1, a2)
}