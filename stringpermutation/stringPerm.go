package main

func stringPerm(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	mismatches := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		mismatches[int(s1[i]-'a')]++
	}
	for i := 0; i < len(s1); i++ {
		mismatches[int(s2[i]-'a')]--
	}
	return true
}

func main() {
	stringPerm("ab", "eidbaooo")
}

/* Finding the permutation in a string
Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1.
In other words, one of the first string's permutations is the substring of the second string.

Example 1:
Input: s1 = "ab" s2 = "eidbaooo"
Output: True
Explanation: s2 contains one permutation of s1 ("ba").

Example 2:
Input:s1= "ab" s2 = "eidboaoo"
Output: False

make array of size 26 to check all the letters of alphabet
*/
