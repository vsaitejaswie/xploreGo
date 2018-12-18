package main 

import "fmt"

func main(){

	// 1
	m := make(map[string]int)

	m["first"] = 1
	m["second"] = 2
	m["third"] = 3

	for _, v := range m{
		fmt.Println(" My value is ", v)
	}

	// 2 map literal
	type bunch struct{
		former int
		latter string
	} 

	n := map[int]bunch{

		1 : bunch{100, "hundred"},
		2 : bunch{200, "hundred doubled"},
	}

	for i := range n{

		fmt.Println(n[i].former, "   ", n[i].latter)

	}

	// 3 dropping the struct name for each assignment 

	p := map[int]bunch{

		1 : {100, "hundred"},
		2 : {200, "hundred doubled"}, // this comma after the final key-value pair is important
	}

	for i := range p{

		fmt.Println(p[i].former, "   ", p[i].latter)

	}

}