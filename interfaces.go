package main  

import "fmt"

type cakes interface{
	decorate() int
}

type tooti_frooti struct{
	A string
}

type vanilla struct{
	B string
}

func (a tooti_frooti) decorate() int{
	//fmt.Println("I'm a colourful cylinder")
	return 10
}

func (b vanilla) decorate() int{
	//fmt.Println("I'm the best of all, certainly!")
	return 20
}

func main(){
	var cake cakes

	alpha := tooti_frooti{}
	beta := vanilla{"fair and lovely"}

	cake = alpha
	cake.decorate()
	fmt.Println("I'm a colourful cylinder")

	cake = beta
	cake.decorate()
	fmt.Println("I'm the best of all, certainly!")

}