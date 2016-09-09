package main

import "fmt"

var x int

func makeGreeter() func() string {
	return func() string {
		return "Hello world!!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Printf("%T\n", increment)
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
