package main

import "fmt"

func main() {
	x := ((true && false) || (false && true) || !(false && false))
	fmt.Println(x)

	a := (true && false)
	fmt.Println(a)

	b := (false && true)
	fmt.Println(b)

	c := !(false && false)
	fmt.Println(c)

	d := (false && false)
	fmt.Println(d)

}
