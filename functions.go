/* Functions */

package main 

import "fmt"

// A function can take zero arguments and to use return method in a function is 

// A function with single return value
func single_return(v1, v2 int) int{
	v1 = 1997
	v2 = 7991
	return v2 
}

// A function can return any number of results
func multi_return(a int, b string) (string, int, int){

	// Assignment of variables within a function
	a = 100
	b = "TejVR"
	return b, a, (a+10)
}

// Named return values or Naked return
func try_named()(name1 string, name2 int){
	
	name2 = 5
	name1 = "named_Return"
	return
}

func main(){
	
	fmt.Println(single_return(10, 20))

	fmt.Println(multi_return(10, "Tej"))

	// Declaration and Assignment of new variables within a function
	x, y, z := multi_return(10, "Tej") 
	fmt.Println(x, y, z)

	a1, a2 := try_named()
	fmt.Println(a1, a2)
}