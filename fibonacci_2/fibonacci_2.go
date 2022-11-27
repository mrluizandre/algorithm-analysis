package main

import "fmt"

var table = make(map[uint64]uint64)

func main() {
	var n uint64 = 10

	result := fibonacci(n)

	fmt.Println(result)
}

func fibonacci(n uint64) uint64 {
	_, isPresent := table[n]
	if isPresent {
		return table[n]
	}

	if n < 2 {
		return n
	} else {
		value := fibonacci(n-1) + fibonacci(n-2)
		table[n] = value
		return value
	}
}
