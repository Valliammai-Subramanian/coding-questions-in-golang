package main

import (
    "fmt"

    "github.com/valliammai-subramanian/coding-questions-in-golang/concatenate"
)

func main() {
    // Get a greeting message and print it.
	ary1 := []int{1, 4, 5, 7}
	ary2 := []int{2, 4, 6, 8}
	output := []int{}
	concatenate.Concatenation(ary1, ary2, &output)
	fmt.Println(output)
}