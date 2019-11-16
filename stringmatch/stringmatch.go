package main

import "fmt"

func main() {
	input := "DDI"
	fmt.Println(stringMatch(input))
}
func stringMatch(input string) []int {
	min := 0
	max := 0
	result := []int{0}

	for _, c := range input {
		if string(c) == "I" {
			max++
			result = append(result, max)
		}
		if string(c) == "D" {
			min--
			result = append(result, min)
		}
	}
	for i := 0; i < len(result); i++ {
		result[i] = result[i] - min
	}

	return result

}
