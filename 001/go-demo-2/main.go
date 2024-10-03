package main

import "fmt"

func main() {
	transactions := [5]int{1, 2, 3, 4, 5}
	transactionsNew := transactions
	banks := [2]string{}

	banks[0] = "Aktif"
	fmt.Println(banks)
	partial := transactions[1:4]
	fmt.Println(partial)
	fmt.Println(transactions)
	fmt.Println(transactionsNew)
}
