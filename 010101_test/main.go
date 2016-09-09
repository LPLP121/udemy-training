package main

import "fmt"

func main() {
	fmt.Printf("%s", "Hello, 世界")
	fmt.Printf("%T: %v\n", "Hello, 世界", "Hello, 世界")

	fmt.Printf("%T %v\n", 0, 0)
	fmt.Printf("%T %v\n", 0.0, 0.0)
	fmt.Printf("%T %v\n", 'x', 'x')
	fmt.Printf("%T %v\n", 0i, 0i)

	type MyBool bool
	const True = true
	const TypedTrue bool = true
	var mb MyBool
	mb = true // OK
	mb = True // OK
	//	mb = TypedTrue // Bad
	fmt.Println(mb)
}
