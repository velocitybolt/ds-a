package main

import (
	"fmt"
	"math"
)

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
*/
// LongestSubstringWithoutRepeatingCharacters ...
func LongestSubstringWithoutRepeatingCharacters(input string) int {

	start := 0              // This used to know which character to remove from the map and then incremented
	maxLen := math.MinInt32 // records min length
	frequency := make(map[string]bool)

	for i := 0; i < len(input); i++ {
		// first check if its in the map
		if frequency[string(input)] {
			// if it is a duplicate then we update the maxLen accordingly
			maxLen = max(i-start, maxLen)
			// delete characters from the map until you find the duplicate
			startChar := string(input[start])
			duplicateChar := string(input[i])
			for {
				delete(frequency, startChar)
				// keep deleting unless you find duplicate, if found then
				if startChar == duplicateChar {
					break
				}
			}
		}
		frequency[string(input)] = true
	}
	return maxLen

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	test := "abcabcbb"
	test1 := "bbbb"
	test2 := "pwwkew"
	lastchar := string(test[len(test)-1])
	fmt.Println(lastchar)
	fmt.Println(LongestSubstringWithoutRepeatingCharacters(test))
	fmt.Println(LongestSubstringWithoutRepeatingCharacters(test1))
	fmt.Println(LongestSubstringWithoutRepeatingCharacters(test2))
	fmt.Println(LongestSubstringWithoutRepeatingCharacters("abcd"))
	fmt.Println(LongestSubstringWithoutRepeatingCharacters("abcdae"))
}
