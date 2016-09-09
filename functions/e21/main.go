package main

import "fmt"

// func main() {
// 	t := 0
// 	for i := 0; i < 1000; i++ {
// 		if i%3 == 0 {
// 			fmt.Println(i)
// 			t += i
// 		}
// 		if i%5 == 0 {
// 			fmt.Println(i)
// 			t += i
// 		}
//
// 	}
// 	fmt.Println(t)
// }

func main() {
	counter := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}
