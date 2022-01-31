package main

import (
	"constraints"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type A struct {
	hoge string
}

func (a A) Hoge() string {
	return a.hoge
}

type B struct {
	hoge string
}

func (b B) Hoge() string {
	return b.hoge
}

func Print[T A | B](x T) {
	switch (interface{})(x).(type) {
	case A:
		fmt.Println("a is ", x)
	case B:
		fmt.Println("b is ", x)
	}
}

func Print1[T string | int](x T) {
	switch (interface{})(x).(type) {
	case string:
		fmt.Println("arg is string", x)
	case int:
		fmt.Println("arg is integer", x)
	}
}

func Filter[K constraints.Ordered](list []K, f func(K) bool) []K {
	result := make([]K, 0, len(list))
	for _, x := range list {
		if f(x) {
			result = append(result, x)
		}
	}
	return result
}

func Example_Filter() {
	fmt.Println("=====  Filter example1 ========")
	input1 := []int{1, 2, 3, 4}
	output1 := Filter(input1, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("input: %#v\n", input1)
	fmt.Printf("output: %#v\n", output1)

	fmt.Println("=====  Filter example2 ========")
	re := regexp.MustCompile(`^1_`)
	input2 := []string{"1_1", "1_2", "2_3", "2_4"}
	output2 := Filter(input2, re.MatchString)
	fmt.Printf("input: %#v\n", input2)
	fmt.Printf("output: %#v\n", output2)
}

func Map[K, T any](list []K, f func(K) T) []T {
	result := make([]T, 0, len(list))
	for _, x := range list {
		result = append(result, f(x))
	}
	return result
}

type stuff struct {
	price int
}

func Example_Map() {
	fmt.Println("=====  Map example1 ========")
	input1 := []int{1, 2, 3, 4}
	output1 := Map(input1, strconv.Itoa)
	fmt.Printf("input: %#v\n", input1)
	fmt.Printf("output: %#v\n", output1)

	fmt.Println("=====  Map example2 ========")
	input2 := []stuff{
		{price: 10},
		{price: 20},
		{price: 30},
	}

	output2 := Map(input2, func(s stuff) stuff {
		return stuff{price: int(math.Ceil(float64(s.price) * 1.5))}
	})
	fmt.Printf("input: %+v\n", input2)
	fmt.Printf("output: %+v\n", output2)
}

func Reduce[K, T any](list []K, f func(acc T, value K, index int) T, initialValue T) T {
	acc := initialValue
	for i, x := range list {
		acc = f(acc, x, i)
	}
	return acc
}

func Example_Reduce() {
	fmt.Println("=====  Reduce example1 ========")

	input1 := []stuff{
		{price: 10},
		{price: 20},
		{price: 30},
	}

	output1 := Reduce(input1, func(acc int, value stuff, index int) int {
		return acc + value.price
	}, 0)
	fmt.Printf("input: %+v\n", input1)
	fmt.Printf("output: %+v\n", output1)
}

func Every[K any](list []K, f func(K) bool) bool {
	for _, x := range list {
		if !f(x) {
			return false
		}
	}
	return true
}

func Some[K any](list []K, f func(K) bool) bool {
	for _, x := range list {
		if f(x) {
			return true
		}
	}
	return false
}

func Includes[K comparable](list []K, elem K) bool {
	return Some(list, func(x K) bool { return x == elem })
}

func IncludesAll[K comparable](list []K, elems []K) bool {
	return Every(elems, func(x K) bool { return Includes(list, x) })
}

var permissionList = []string{
	"create", "read", "update", "delete",
}

type User struct {
	permission []string
}

func (u *User) CanDelete() bool {
	return Some(u.permission, func(x string) bool {
		return x == "delete"
	})
}

func (u *User) IsAdmin() bool {
	return IncludesAll(u.permission, permissionList)
}

func Example_Some() {
	user1 := User{
		permission: []string{"read"},
	}
	fmt.Printf("person1 can delete xxx: %v\n", user1.CanDelete())
	fmt.Printf("person1 is admin: %v\n", user1.IsAdmin())

	user2 := User{
		permission: []string{"create", "delete"},
	}
	fmt.Printf("person2 can delete xxx: %v\n", user2.CanDelete())
	fmt.Printf("person2 is admin: %v\n", user2.IsAdmin())

}

func main() {
	// Example1()
	// Print(A{hoge: "AAA"})
	// Print(B{hoge: "BBB"})

	// Example_Filter()
	// Example_Map()
	// Example_Reduce()
	Example_Some()
}
