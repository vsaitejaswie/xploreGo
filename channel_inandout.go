package main 

import "fmt"

func main(){

	n := 3

	in := make(chan int)
	out := make(chan int)

	go double_it(in, out)

	in <- n

	fmt.Println(<-out)
}

func double_it(in <-chan int, out chan<- int){
	x := <-in*2
	out <- x
}