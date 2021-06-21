package reformat

import (
	"fmt"
	"regexp"
)

func reformat(S string) string {

	reg, _ := regexp.Compile("[^0-9]+")

	numericStr := reg.ReplaceAllString(S, "")

	n := len(numericStr)
	remainder := n%3
	quotient := n/3

	if (remainder == 1){
		formattedStr := numericStr[:3]
		for i := 1; i < quotient - 1; i++ {
			formattedStr = formattedStr + "-" + numericStr[i*3:(i+1)*3]			
		}
		numericStr = formattedStr + "-" + numericStr[n-4:n-2] + "-" + numericStr[n-2:]		
	}	else {
		for i := 3; i < len(numericStr); i += 4 {
			numericStr = numericStr[:i] + "-" + numericStr[i:]
		}
	}
	return numericStr
}

func Main() {
	fmt.Println("One block with two digits:", reformat("00-44  48 5555 8361"))
	fmt.Println("Two blocks with two digits:", reformat("0 - 22 1985--324"))
	fmt.Println("No block with two digits:", reformat("555372654"))
}