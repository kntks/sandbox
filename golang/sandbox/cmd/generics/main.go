package main

import "fmt"

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

func main() {
	// Example1()
	Print(A{hoge: "AAA"})
	Print(B{hoge: "BBB"})

	Print1("a")
	Print1(1)
}
