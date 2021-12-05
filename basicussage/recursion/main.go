package main

import (
	"fmt"
)

func fib(n int) int {
	if n == 0 {
		return 1
	}

	return n * fib(n - 1)
}

func main() {
	fmt.Println(fib(7))
}