package main

import "fmt"

/* Function that finds the ways to calculate n */
func waysToGetN(n int) int {
	dp := make([]int, n+1)
	// handle base cases
	dp[0] = 1
	dp[1] = 1
	dp[2] = 1
	dp[3] = 2

	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}

func main() {
	fmt.Println(waysToGetN(5))
}
