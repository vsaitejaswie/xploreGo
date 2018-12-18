package main  

import "fmt"

func main(){

	array := []string{"one", "two", "three", "four", "five", "six"}

	// i refers to the key and v refers to the value
	for i, v := range array{
		fmt.Println(" The key being referenced is ", i)
		fmt.Println(" The value being referenced is ", v, "\n")
	}

	// You can skip the index or value by assigning to underscore

	for _, v := range array{
		fmt.Println(" Only the value ", v, "\n")
	}
}