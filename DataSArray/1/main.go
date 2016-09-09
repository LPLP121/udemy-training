package main

import "fmt"

func main() {
	var x [58]string
	fmt.Println(x)      //
	fmt.Println(len(x)) // 58
	fmt.Println(x[42])  // same as first print, but with 42, instead of 58

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}
