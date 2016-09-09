package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

	fmt.Println("++++++++++")

	var s2 []string
	var s2s [][]string
	fmt.Println(s2)
	fmt.Println(s2s)
	fmt.Println(s2 == nil)

	fmt.Println("++++++++++")

	s3 := make([]string, 35)
	s3s := make([][]string, 35)
	fmt.Println(s3)
	fmt.Println(s3s)
	fmt.Println(s3 == nil)

}
