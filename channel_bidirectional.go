package main  

import "fmt"

func main(){

	store := make(chan int)
	num := 10

	go double_it(num, store) 

	fmt.Println(<-store)
}

func double_it(num int, out chan<- int){

	out <- (num*2)
}