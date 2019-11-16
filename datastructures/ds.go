package main

import "fmt"

type stack []string

func main() {
	m := make(map[string]int)
	m["route"] = 66
	fmt.Println(m)

	var stack stack
	//push to back
	stack = append(stack, "(")
	stack = append(stack, ")")
	fmt.Println(stack)
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack)
}

func (s *stack) pop() string {
	n := len(*s) - 1
	el := (*s)[n]
	*s = (*s)[:n]
	return el

}
