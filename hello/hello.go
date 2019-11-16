package main

import (
	"errors"
	"fmt"
	"strings"
)

func divby(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("You idiot")
	}

	return a / b, nil
}

func main() {
	test()

	lolvar := []string{"aaple", "fdd", "SFf", "SFd"}
	for _, str := range lolvar {
		if strings.Contains(str, "aa") {
			lolvar = append(lolvar, "aa")
		}
	}
	fmt.Println(lolvar)

	res, err := divby(5, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
