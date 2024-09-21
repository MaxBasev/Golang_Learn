package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight = 1.8
	var userWeight float64 = 60
	var BMI = userWeight / math.Pow(userHeight, 2)

	fmt.Println("Body mass index calculation.")
	fmt.Print(BMI)
}
