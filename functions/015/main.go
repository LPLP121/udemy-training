package main

import "fmt"

func main() {
	if !true {
		fmt.Println("this is true")
	}

	if !false {
		fmt.Println("this is false")
	}
	test()
}

func test() {
	if true || false {
		fmt.Println("func test - this ran")
	}
}
