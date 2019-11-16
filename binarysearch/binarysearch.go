package main

import "fmt"

/* This file contains the iterative binary search O(log n) rt, O(1) space*/

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 6}, 6))
}

func binarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		midpoint := (start + end) / 2

		// check if target is midpoint
		if target == arr[midpoint] {
			return midpoint
		}
		// check if target is in left or right subarray
		if target > arr[midpoint] {
			start = midpoint + 1
		} else {
			end = midpoint - 1
		}

	}

	// not found in array
	return -1
}
