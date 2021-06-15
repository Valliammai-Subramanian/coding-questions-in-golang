package main

import (
    "fmt"
	"log"

    "github.com/valliammai-subramanian/coding-questions-in-golang/concatenate"
)

func main() {
    ary1 := []int{1,3,5,7}
	ary2 := []int{2, 4, 6, 8}
	output := []int{}	
	
	log.SetPrefix("concatenate: ")
	log.SetFlags(0)
	 
	err := concatenate.Concatenation(ary1, ary2, &output)

 	if err != nil {
	 log.Fatal(err)
 	}

	fmt.Println(output)
}