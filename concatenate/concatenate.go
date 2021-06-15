package concatenate

import (
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