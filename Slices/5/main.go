package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2])
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2])
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:])
	fmt.Print("[:] ")
	fmt.Println(greeting[:])

	customerNumber := make([]int, 3)
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting2 := make([]string, 3, 5)
	greeting2[0] = "test1"
	greeting2[1] = "test2"
	greeting2[2] = "test3"
	greeting2 = append(greeting2, "test4")

	fmt.Println(greeting2[0])
	fmt.Println(greeting2[1])
	fmt.Println(greeting2[2])
	fmt.Println(greeting2[3])

	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8, 9}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)

	my2 := []string{"M", "Tu"}
	my3 := []string{"W", "Th", "F"}

	my2 = append(my2, my3...)
	fmt.Println(my2)

	my2 = append(my2[:2], my2[3:]...)
	fmt.Println(my2)

}
