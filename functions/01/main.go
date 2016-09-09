package main

import "fmt"

func main() {
	greet("Jane")
	greet("John")
	greet2("Jimbo", "James")
	greet3("James", "Jimbo")
}

func greet(name string) {
	fmt.Println(name)
}

func greet2(fname, lname string) {
	fmt.Println(fname, lname)
}

func greet3(lname, fname string) {
	fmt.Println(lname, ",", fname)
}
