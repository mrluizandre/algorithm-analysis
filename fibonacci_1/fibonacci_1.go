package main

import "fmt"

func main() {
	n := 10

	result := fibonacci(n)

	fmt.Println(result)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
