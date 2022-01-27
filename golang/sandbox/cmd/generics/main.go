package main

import (
	"constraints"
	"fmt"
)

func sum[K constraints.Ordered, T []K](m T) K {
	var sum K
	for _, x := range m {
		sum += x
	}
	return sum
}

func some[K constraints.Ordered](list []K, f func(K) bool) bool {
	for _, x := range list {
		if f(x) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(sum([]int{1, 2, 3}))
	fmt.Println(sum([]string{"a", "b"}))

	fmt.Println(some([]int{1, 3, 4}, func(x int) bool { return x%2 == 0 }))
}
