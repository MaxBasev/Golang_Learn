package main

import "fmt"

func main() {
	transactions := [5]int{1, 2, 3, 4, 5}
	banks := [2]string{}

	fmt.Println(transactions[0])
	banks[0] = "Aktif"
	fmt.Println(banks)
	partial := transactions[1:4]
	fmt.Println(partial)
}
