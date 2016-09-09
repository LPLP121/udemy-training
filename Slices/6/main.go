package main

import "fmt"

func main() {
	records := make([][]string, 0)
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"
	//store the records
	records = append(records, student1)
	//student 2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"
	// store the records
	records = append(records, student2)
	// Print
	fmt.Println(records)

	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		tr := make([]int, 0)
		for j := 0; j < 4; j++ {
			tr = append(tr, j)
		}
		transactions = append(transactions, tr)
	}
	fmt.Println(transactions)

}
