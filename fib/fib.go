package main

import "fmt"

func main() {
	fmt.Println(memoFib(5, map[int]int{}))
}

func fib(n int) int {

	if n < 0 || n == 0 || n == 1 {
		return n
	}
	fmt.Printf("computing fib %d\n", n)
	return fib(n-1) + fib(n-2)

}

func memoFib(n int, seen map[int]int) int {
	result := 0

	if n <= 1 {
		return n
	}
	if val, ok := seen[n]; ok {
		fmt.Printf("Getting Memo %d\n", val)
		return val
	}

	fmt.Printf("Computing fib %d\n", n)
	result = memoFib(n-1, seen) + memoFib(n-2, seen)

	seen[n] = result

	return result
}
