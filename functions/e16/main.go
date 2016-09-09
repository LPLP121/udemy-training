package main

import "fmt"

// func main() {
//
// 	fmt.Println(test(2))
// 	fmt.Println(test(4))
// 	fmt.Println(test(5))
//
// 	yes, true := test(7)
// 	fmt.Println(yes, true)
// }
//
// func test(x int) (int, bool) {
// 	return x / 2, x%2 == 0
// }

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(halffloat(1))
	fmt.Println(halffloat(2))
	fmt.Println(half(5))
	fmt.Println(halffloat(5))
}

func halffloat(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0
}
