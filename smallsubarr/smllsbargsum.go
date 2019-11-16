package main

import (
	"fmt"
	"math"
)

/* Objective: Given an array of n positive integers and a positive integer s,
find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

Example:

Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.

*/

func smallestSubArray(input []int, target int) int {

	currentSum := 0          // this records the current sum of the elements
	start := 0               // start index of window will need to incremented after each update of minSize
	minSize := math.MaxInt32 // records the size of the minimum subarray

	for i := 0; i < len(input); i++ {
		currentSum += input[i] // add next int to the curr sum

		for currentSum >= target {
			minSize = min(minSize, i-start+1)
			currentSum -= input[start]
			start++
		}

	}
	return minSize

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	testArr := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(smallestSubArray(testArr, 7))
}
