package main

import "fmt"

type person struct {
	age  int
	next *person
	prev *person
}

func main() {
	fmt.Println("To learn...")
}
