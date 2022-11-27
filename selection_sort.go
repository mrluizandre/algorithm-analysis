package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	list := rand.Perm(20)
	n := len(list)
	var min, aux int

	fmt.Println("Before:", list)

	for i := 0; i < (n - 1); i++ {
		min = i

		for j := i + 1; j < n; j++ {
			if list[j] < list[min] {
				min = j
			}
		}

		if list[i] != list[min] {
			aux = list[i]
			list[i] = list[min]
			list[min] = aux
		}
	}

	fmt.Println("After:", list)
}
