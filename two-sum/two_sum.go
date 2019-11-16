package main

import (
	"errors"
	"fmt"
)

/* This file contains the twoSum function which returns the indices of two values which add up to the target */
func main() {
	values := []int{1, 2, 5, 3, 0, -56, 7}
	target := -49

	a, b, err := searchForTwoSum(values, target)
	if err == nil {
		fmt.Printf("values[%d] + values[%d] = %d\n", a, b, target)
	} else {
		fmt.Println(err)
	}
}

func searchForTwoSum(values []int, target int) (int, int, error) {
	seen := make(map[int]int)

	for i, value := range values {
		if j, ok := seen[target-value]; ok {
			return i, j, nil
		}

		seen[value] = i
	}

	return 0, 0, errors.New("No pair found")
}
