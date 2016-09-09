package main

//var x = 42
import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, "-", j)
		}
	}
	foo()
}

func foo() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

/*
func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)

}
*/
