package main 

import "fmt"

func main(){

	n := 3

	in := make(chan int)
	out := make(chan int)

	go double_it(in, out)

	in <- n

	fmt.Println("Doubled number is: ", <-out)
}

func double_it(in <-chan int, out chan<- int){

	fmt.Println("Goroutine is initializing.........")

	out <- (<-in*2)
}