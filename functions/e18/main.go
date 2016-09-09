package main

import "fmt"

// func greatestNumber(args ...int) int {
// 	max := args[0]
// 	for _, arg := range args {
// 		if arg > max {
// 			max = arg
// 		}
// 	}
// 	return max
// }
//
// func main() {
// 	fmt.Println(greatestNumber(-12, -3, -84, -1, -100))
// }

func max(numbers ...int) int {
	var largest int
	for i, v := range numbers {
		if v > largest || i == 0 {
			largest = v
		}
	}
	return largest
}

func main() {
	greatest := max(-4, -7)
	fmt.Println(greatest)
}
