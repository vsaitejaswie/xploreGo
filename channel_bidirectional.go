package main  

import "fmt"

func main(){

	store := make(chan int)
	num := 10

	go double_it(num, store) 

	fmt.Println("Doubled number is: ", <-store)
}

func double_it(num int, out chan<- int){

	fmt.Println("Goroutine is initializing.........")

	out <- (num*2)
}