package main

import (
	"fmt"
)

// main function
func NumberToLetterIncrement(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += NumberToLetterIncrement(firstLetter) + string('A'+number%26)
	} else {
		letters += string('A' + number)
	}
	return
}

// helpers for testing
func main() {
	fmt.Println("Hello, playground")
	fmt.Println(NumberToLetterIncrement(110))
}
