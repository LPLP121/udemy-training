package main

import "fmt"

func main() {
	//var name string
	//fmt.Scan(&name)
	//fmt.Println("Hello", name)

	var LargeNum int
	var SmallNum int
	//var Remainder int
	//LargeNum = 5
	//SmallNum = 3

	fmt.Print("Type a Large Number: ")
	fmt.Scan(&LargeNum)
	fmt.Print("Type a Smaller Number: ")
	fmt.Scan(&SmallNum)
	//Remainder = LargeNum % SmallNum
	fmt.Println(LargeNum, "divided by", SmallNum, "is", LargeNum/SmallNum)
	fmt.Println("The remainder of", LargeNum, "divided by", SmallNum, "is", LargeNum%SmallNum)
}
