package main

import (
	"fmt"
)

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b = &a
	fmt.Println(b)
	fmt.Println(*b)

	c := 95
	fmt.Println(c)
	fmt.Println(&c)

	var d = &c
	fmt.Println(d)
	fmt.Println(*d)

	fmt.Printf("%b ; %d ; %T \n", b, b, b)
	fmt.Printf("%b ; %d ; %T \n", d, d, d)
}
