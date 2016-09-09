package main

import "fmt"

func main() {
	for x := 0; x <= 100; x++ {
		if x%3 == 0 && x%5 == 0 {
			fmt.Println(x, "FizzBuzz")
		} else if x%3 == 0 {
			fmt.Println(x, "Fizz")
		} else if x%5 == 0 {
			fmt.Println(x, "Buzz")
		}

	}

}
