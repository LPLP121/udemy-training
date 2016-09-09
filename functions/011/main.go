package main

import "fmt"

func main() {
	age := "Test"
	fmt.Println(&age)

	changeMe(&age)

	fmt.Println(&age)
	fmt.Println(age)
}

func changeMe(z *string) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = "test2"
	fmt.Println(z)
	fmt.Println(*z)

}
