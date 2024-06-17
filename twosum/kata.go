package main

import "fmt"

func main() {

	fmt.Println(TwoSum(intArr, intTarget))
}

// sample array
var intArr []int = []int{1, 2, 3, 4}

// sample target
var intTarget int = 5

func TwoSum(arr []int, target int) [2]int {

	// first counter
	for i := 0; i < len(arr); i++ {
		// second counter i+1
		for j := i + 1; j < (len(arr)); j++ {
			// check if i + j = target num
			if arr[i]+arr[j] == target {

				return [2]int{i, j}
			}
		}
	}
	return [2]int{0, 0}
}
