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

func SpinWords(sentence string) string {
	// variable for words in sentence
	var sentenceArr []string = strings.Split(sentence, " ")

	// var holding reversed string in array
	var reversedStrArr []string
	// var holding reversed string

	for i, word := range sentenceArr {
		// check if len of word >= 5
		if len(word) >= 5 {
			// string reverser
			for j := 0; j < len(word); j++ {
				reversedStrArr = append(reversedStrArr, string(word[len(word)-j-1]))
			}
			var reversedStr = strings.Join(reversedStrArr, "")

			// assign reversed str to a word
			sentenceArr[i] = reversedStr

			// clean reversed str and strArr
			reversedStr = ""
			reversedStrArr = nil

		}
	}
	// turn slice into string var
	var sentenceStr = strings.Join(sentenceArr, " ")

	return sentenceStr
}
