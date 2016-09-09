package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlide := []int{1, 2, 3, 4}
	foo(aSlide...)
	foo()
	foo2(1, 2)
	foo2(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo2(aSlice...)
	foo2()
}

func foo(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

func foo2(num2 ...int) {
	fmt.Println(num2)
}
