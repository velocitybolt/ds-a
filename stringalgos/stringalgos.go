package main

import "fmt"

func main() {
	arr := make([]int, 26)
	test := "aaabb"

	for i := 0; i < len(test); i++ {
		arr[test[i]-'a']++
	}
	fmt.Println(arr)

}
