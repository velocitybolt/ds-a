package main

import "fmt"

func main() {
	test := []string{"S", "T", "R"}
	for _, c := range test {
		defer fmt.Println(c)
	}
}
