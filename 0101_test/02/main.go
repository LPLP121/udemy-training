package main

import "fmt"

func main() {

	var a int
	var b string
	var c float64
	var d bool
	var e int32

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)

	fmt.Println()

	s := "Hello World!"
	fmt.Println(s)
	s = "Goodbye cruel World!"

	fmt.Println("s = ", s)

}
