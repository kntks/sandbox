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

type Number interface {
	constraints.Integer | constraints.Float
}

func sum2[K Number, T []K](list T) K {
	return sum[K](list)
}

func main() {
	fmt.Println(sum([]int{1, 2, 3}))
	fmt.Println(sum([]string{"a", "b"}))

	// error: string does not implement Number
	// fmt.Println(sum2([]string{"a", "b"}))
	fmt.Println(sum2([]float64{1.2, 3.4}))

	fmt.Println(some([]int{1, 3, 4}, func(x int) bool { return x%2 == 0 }))
}
