package main

import "fmt"

func main() {
	/* Given string s => cbaebabacd"
	   Find the start indices of anagrams in p */

	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s, p))

}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	mismatches := make([]int, 26)

	// Anagram starts at first index
	for i := 0; i < len(p); i++ {
		mismatches[int(p[i]-'a')]++
	}
	for i := 0; i < len(p); i++ {
		mismatches[int(s[i]-'a')]--
	}

	// This means anagram starts at first index
	ans := []int{}
	if check(mismatches) {
		ans = append(ans, 0)
	}

	// iterate through remaining characters across both strings disregarding first element
	for i := 1; i+len(p) <= len(s); i++ {
		end := i + len(p) - 1
		mismatches[int(s[end]-'a')]--
		mismatches[int(s[i-1]-'a')]++
		fmt.Println("This is the end values", end, string(s[end]))
		fmt.Println("This is the values after the first", i, string(s[i-1]))
		if check(mismatches) {
			ans = append(ans, i)
		}

	}
	return ans
}

func check(mismatches []int) bool {
	for _, mismatch := range mismatches {
		if mismatch != 0 {
			return false
		}
	}
	return true
}
