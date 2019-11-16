package main

import (
	"fmt"
	"unicode"
)

func main() {
	input := "A man, a plan, a canal: Panama"
	input2 := "race a car"
	fmt.Println(validPalindrome(input))
	fmt.Println(validPalindrome(input2))

}

func validPalindrome(input string) bool {
	start := 0
	end := len(input) - 1

	for start < end {
		a := unicode.ToLower(rune(input[start]))
		b := unicode.ToLower(rune(input[end]))

		if !unicode.IsLetter(a) && !unicode.IsDigit(a) {
			start++
			continue
		}

		if !unicode.IsLetter(b) && !unicode.IsDigit(b) {
			end--
			continue
		}

		if a == b {
			start++
			end--
		} else {
			return false
		}
	}
	return true

}
