package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	list := rand.Perm(100)

	fmt.Println("Before:", list)
	sorted_list := mergeSort(list)

	fmt.Println("After:", sorted_list)
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	left := mergeSort(items[:len(items)/2])
	right := mergeSort(items[len(items)/2:])
	return merge(left, right)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
