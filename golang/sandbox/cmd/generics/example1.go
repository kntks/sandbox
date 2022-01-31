package main

import (
	"constraints"
	"fmt"
)

func Sum[K Number](list []K) K {
	var sum K
	for _, x := range list {
		sum += x
	}
	return sum
}

type Number interface {
	constraints.Integer | constraints.Float
}

func Example1() {
	fmt.Println(Sum([]int{1, 2, 3}))

	fmt.Println(Sum([]float64{1.2, 3.4}))
}
