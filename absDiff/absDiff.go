package absDiff

import (
	"fmt"	
)

func diagDiff(arr [][]int) int{
	sum := 0
	n := len(arr)
	for i:=0; i<n; i++{
		sum = sum + arr[i][i] - arr[i][n-i-1]
	}

	if sum < 0 {
		return -sum
	}
	return sum		
}

func Main(){
	inputArr := [][]int{
					{11,2,4},
					{4,5,6},
					{10,8,-12},
				}

fmt.Println("Absolute Difference:",diagDiff(inputArr))
}