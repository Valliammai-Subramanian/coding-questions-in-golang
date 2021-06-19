package concatenate

import (
	"fmt"
	"log"
	"sort"
	"errors"
)

func Concatenation(ary1 []int, ary2 []int, output *[]int)(error) {
	if len(ary1) == 0 || len(ary2) == 0 {
		return errors.New("Empty array")
	}

	for _, v := range ary2 {
		ary1 = append(ary1, v)
	}
	sort.Ints(ary1)
	*output = ary1

	return nil
}

func Main() {
    ary1 := []int{1,3,5,7}
	ary2 := []int{2, 4, 6, 8}
	output := []int{}	
	
	log.SetPrefix("concatenate: ")
	log.SetFlags(0)
	 
	err := Concatenation(ary1, ary2, &output)

 	if err != nil {
	 log.Fatal(err)
 	}

	fmt.Println(output)	
}