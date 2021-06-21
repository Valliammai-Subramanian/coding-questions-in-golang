package substring

import (
	"fmt"
	"strings"
)

func subString(A []string, B []string, P string) string {
    result := ""

    for i:=0 ; i<len(B) ; i++ {
        if strings.Contains(B[i], P) == true {              
                name := A[i]
				//fmt.Println("Enters the loop:",name)

				if (result == "") || (name < result) {
					result = A[i] 
					//fmt.Println("String compare:",result)
				}
				
            }       
        }

	if (result == ""){
		result = "NO CONTACT"
	}

    return result    
}

func Main() {
	names1 := []string {"pim","pom"}
	phone1 := []string {"999999999","777888999"}
	findStr1 := "88999"
	fmt.Println("Single match:",subString(names1, phone1, findStr1))

	names2 := []string {"sander","amy","ann","michael"}
	phone2 := []string {"123456789","234567890","789123456","123123123"}
	findStr2 := "1"
	fmt.Println("Multiple match:",subString(names2, phone2, findStr2))
	
	names3 := []string {"adam","eva","leo"}
	phone3 := []string {"121212121","111111111","444555666"}
	findStr3 := "112"
	fmt.Println("No match:",subString(names3, phone3, findStr3))	
}