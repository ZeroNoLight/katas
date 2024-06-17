package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(SpinWords(sentence))
}

// test vars
var sentence string = "This is another test"

// write a function that spins words
// write second function that check if word > 5

func SpinWords(sentence string) string {
	// variable for words in sentence
	var words []string = strings.Split(sentence, " ")

	// test for len of words in sentence
	fmt.Printf("Len of arr words is: %v\n", len(words))

	// var holding reversed string in array
	var reversedStrArr []string
	// var holding reversed string

	for i, word := range words {
		// check if len of word >= 5
		if len(word) >= 5 {
			// string reverser
			for j := 0; j < len(word); j++ {
				reversedStrArr = append(reversedStrArr, string(word[len(word)-j-1]))
			}
			var reversedStr = strings.Join(reversedStrArr, "")
			// test for reversedStr value
			fmt.Printf("Reversed string is: %v\n", reversedStr)

			// assign reversed str to a word
			words[i] = reversedStr

			// clean reversed str and strArr
			reversedStr = ""
			reversedStrArr = nil

			// test for reversed str and strArr clean
			fmt.Printf("Clean reversed strArr: %v\n", reversedStr)
			fmt.Printf("Clean reversed strArr: %v\n", reversedStrArr)

			// test for reversed word
			fmt.Printf("Reversed word is: %v\n", word)

			// test for reversed sentence
			fmt.Printf("Reversed sentence is: %v\n", words)
		}
	}
	// turn slice into string var
	var strWords = strings.Join(words, " ")

	// test if return statement is str
	fmt.Printf("The return statement is: %v\n", strWords)

	return strWords
}
