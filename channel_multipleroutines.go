package main 

import "fmt"

func main(){

	in := make(chan int)
	out := make(chan int)

	go double_it(in, out)
	go double_it(in, out)
	go double_it(in, out)

	in <- 100
	in <- 200
	in <- 300

	fmt.Println("Doubled number is: ", <-out)
	fmt.Println("Doubled number is: ", <-out)
	fmt.Println("Doubled number is: ", <-out)

}

func double_it(in <-chan int, out chan<- int){

	x := <-in

	fmt.Println("Goroutine is initializing with input", x,"...............")

	out <- (x*2)
}

/* 

If this function is used instead of the one above, DEADLOCK occurs 
beacuse two of the channel inputs are already used and the program falls short opf resources


func double_it(in <-chan int, out chan<- int){

	fmt.Println("Goroutine is initializing with input", <-in ,"...............")

	out <- (<-in*2)
}

*/