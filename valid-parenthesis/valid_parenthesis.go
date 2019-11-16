package main

import (
	"fmt"
	"strings"
)

func main() {
	tests := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	for _, test := range tests {
		fmt.Printf("%s: %t\n", test, isValid(test))
	}
}

var oppositeRunes = map[rune]rune{')': '(', '}': '{', ']': '['}

func isValid(s string) bool {
	var seen []rune
	for _, c := range s {
		if strings.ContainsRune("[({", c) {
			seen = append(seen, c)
		} else if len(seen) > 0 && seen[len(seen)-1] == oppositeRunes[c] {
			seen = seen[:len(seen)-1]
		} else {
			return false
		}
	}
	return len(seen) == 0
}
