package main

import (
	"concatenate"
	"hello"
	"fmt"
)

func callConcatenate(){
	ary1 := []int{1, 4, 5, 7}
	ary2 := []int{2, 4, 6, 8}
	output := []int{}
	concatenate.Concatenation(ary1, ary2, &output)
	fmt.Println(output)
}

func callHello(){
	hello.Hello()
	fmt.Println("Done!")
}

func main(){
	callConcatenate()
	callHello()
}
