package main

import "fmt"

func main() {
	transactions := [3]int{1, 5, 8}
	banks := [3]string{"AmeX", "Revolut"}

	fmt.Println(transactions[0])
	banks[2] = "Aktif"
	fmt.Println(banks)
}
