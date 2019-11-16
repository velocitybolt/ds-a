package main

import (
	"fmt"
	"strconv"
)

func main() {
	test1 := []string{"a", "a", "b", "b", "c", "c", "c"}
	test2 := []string{"a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"}
	test3 := []string{"a"}
	fmt.Println(compString(test1))
	fmt.Println(compString(test2))
	fmt.Println(compString(test3))
}

func compString(input []string) string {

	result := ""
	current := string(0)
	count := 0

	for _, c := range input {
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
	// add the remaining portion
	if count > 0 {
		result += string(current)
		result += strconv.Itoa(count)
	}
	return result
}
