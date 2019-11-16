package main

import (
	"fmt"
	"math"
)

// Find max sum subarray of the value k

func maxSubArr(input []int, sumT int) int {

	currentSum := 0         // record the sum of the current window
	maxSum := math.MinInt32 // record the highest possible sum

	for i := 0; i < len(input); i++ {
		currentSum += input[i] // add the current input i to currentsum

		if i >= sumT-1 { // check if the current index has exceeded the size of the window
			maxSum = max(currentSum, maxSum) // calculate which sum is bigger between current and max, then set it to max
			currentSum -= input[i-(sumT-1)]  // subtract the unneeded number out of the window
		}

	}
	return maxSum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	testArr := []int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}
	fmt.Println(maxSubArr(testArr, 3))
}
