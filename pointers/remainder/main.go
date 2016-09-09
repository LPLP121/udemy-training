package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Println(x)
	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
	for i := 1; i < 70; i++ {
		if i%2 == 1 {
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}
	}

	for t := 87; t <= 101; t++ {
		if t+1 <= 97 {
			fmt.Println(t, "yes")
		} else {
			fmt.Println(t, "nope")
		}
	}

}
