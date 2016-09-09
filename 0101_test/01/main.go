package main

import "fmt"

func main() {

	z := 10
	fmt.Printf("%v \n", z)
	fmt.Printf("%T \n", z)

	x := "test"
	fmt.Printf("%v \n", x)
	fmt.Printf("%T \n", x)

	y := 3.2
	fmt.Printf("%v \n", y)
	fmt.Printf("%T \n", y)

	w := 'A'
	fmt.Printf("%v \n", w)
	fmt.Printf("%T \n", w)

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
}
