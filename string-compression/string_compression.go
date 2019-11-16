package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "aaAAbbcDDDdCCCaaBBbbb"
	fmt.Println(compressString(s))
}

func compressString(s string) string {
	result := ""
	current := rune(0)
	count := 0

	for _, c := range s {
		if c != current {
			if count > 0 {
				result += string(current)
				result += strconv.Itoa(count)
				count = 0
			}
			current = c
		}

		count++
	}

	if count > 0 {
		result += string(current)
		result += strconv.Itoa(count)
	}

	return result
}
