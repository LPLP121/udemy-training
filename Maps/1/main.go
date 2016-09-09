package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)

	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
	fmt.Println(myGreeting["Jenny"])

	var myG2 = map[string]string{"Tom": "yes"}
	myG2["John"] = "yo"

	fmt.Println(myG2)

	G3 := map[string]string{
		"Hey":    "man",
		"What's": "up",
	}
	fmt.Println(G3)
	fmt.Println(len(G3))
	G3["That's"] = "right"
	fmt.Println(G3)
	fmt.Println(len(G3))
	G3["That's"] = "cool"
	fmt.Println(G3)
	fmt.Println(len(G3))
	delete(G3, "Hey")
	fmt.Println(G3)
	fmt.Println(len(G3))

	fmt.Println("*******")

	newGreet := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(newGreet)
	//	delete(newGreet, 2)
	if val, exists := newGreet[2]; exists {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(newGreet)

}
