package main

import (
	"fmt"
	"math"
)

/*
You are given coins of different denominations and a total amount of money amount.
Write a function to compute the fewest number of coins that you need to make up that amount.
If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:
Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:
Input: coins = [2], amount = 3
Output: -1
*/

func main() {
	coinSet := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChangeMemo(coinSet, amount))
}

func coinChangeMemo2(coins []int, targetAmount int) int {
	coinsNeeded := map[int]int{}

	for lowerTargetAmount := 1; lowerTargetAmount <= targetAmount; lowerTargetAmount++ {
		for _, coin := range coins {
			if lowerTargetAmount < coin {
				continue
			}

			if lowerTargetAmount == coin {
				coinsNeeded[lowerTargetAmount] = 1
				continue
			}

			needed, ok := coinsNeeded[lowerTargetAmount-coin]
			if !ok {
				continue
			}

			neededForLTA, ok := coinsNeeded[lowerTargetAmount]
			if ok && neededForLTA <= needed+1 {
				continue
			}

			coinsNeeded[lowerTargetAmount] = needed + 1
		}
	}

	needed := coinsNeeded[targetAmount]
	if needed == 0 {
		return -1
	}

	return needed
}

func coinChangeMemo(coinSet []int, amount int) int {
	seen := make([]int, amount+1)
	seen[0] = 0

	// intialize values in array
	for i := 1; i <= amount; i++ {
		seen[i] = math.MaxInt32
	}
	for i := 0; i <= amount; i++ {
		for _, c := range coinSet {
			//fmt.Println(i, c)
			if i-c >= 0 {
				fmt.Println(seen[i], seen[i-c]+1)
				seen[i] = int(math.Min(float64(seen[i]), float64(seen[i-c]+1)))
				fmt.Println(seen[i])
			}
		}
	}
	if seen[amount] == math.MaxInt32 {
		return -1
	}

	return seen[amount]

}
