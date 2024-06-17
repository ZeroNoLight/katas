package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(ValidBraces(sample))

}

// sample var
var sample string = "(){}[]["

func ValidBraces(str string) bool {
	// slice of a string
	strArr := strings.Split(str, "")

	for i := 0; i < (len(strArr) - 1); i++ {

		// switch statements
		switch strArr[i] {
		case "(":
			if strArr[i+1] == ")" {
				newStrArr := append(strArr[:i], strArr[i+2:]...)
				// assign newStrArr to strArr
				strArr = newStrArr
				// clear the counter
				i = -1
				// test
				fmt.Printf("strArr is: %v\n", strArr)
				fmt.Printf("len of strArr is now: %v\n", len(strArr))
			}
		case "[":
			if strArr[i+1] == "]" {
				newStrArr := append(strArr[:i], strArr[i+2:]...)
				// assign newStrArr to strArr
				strArr = newStrArr
				// clear the counter
				i = -1
				// test
				fmt.Printf("strArr is: %v\n", strArr)
				fmt.Printf("len of strArr is now: %v\n", len(strArr))
			}
		case "{":
			if strArr[i+1] == "}" {
				newStrArr := append(strArr[:i], strArr[i+2:]...)
				// assign newStrArr to strArr
				strArr = newStrArr
				// clear the counter
				i = -1
				// test
				fmt.Printf("strArr is: %v\n", strArr)
				fmt.Printf("len of strArr is now: %v\n", len(strArr))
			}
		}
	}
	// return true if len(strArr) == 0
	return len(strArr) == 0
}
