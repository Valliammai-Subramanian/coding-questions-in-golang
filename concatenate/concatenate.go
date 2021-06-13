package concatenate

import (
	"sort"
)

func Concatenation(ary1 []int, ary2 []int, output *[]int) {
	for _, v := range ary2 {
		ary1 = append(ary1, v)
	}
	sort.Ints(ary1)
	*output = ary1
}