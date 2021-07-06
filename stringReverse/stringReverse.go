package stringReverse

import (
	"fmt"	
)

func reverse(v []rune, c0 chan int) {	

	c1:= make (chan int)

	length := len(v)

	if(length == 0){
    	c0 <- 1
		return
	}

	if(length == 1){
		c0 <- 1
		return
	}
	
	go reverse(v[1:length-1], c1)
	v[0], v[length-1] = v[length-1], v[0]
	<-c1
	c0 <- 1
		
}

func Main() {

	str:= "abcdefgh"
	char:= []rune(str)

	c0:= make (chan int, 1)
	
	reverse (char, c0)
	<-c0

	fmt.Println("Given string:", str)
	fmt.Println("Reverse string:", string(char))
}