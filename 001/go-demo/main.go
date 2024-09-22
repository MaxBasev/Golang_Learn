package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower = 2
	var userHeight float64
	var userWeight float64
	fmt.Println("Body mass index calculation.")
	fmt.Print("Type in your height (meter): ")
	fmt.Scan(&userHeight)
	fmt.Print("Type in your weight: ")
	fmt.Scan(&userWeight)

	BMI := userWeight / math.Pow(userHeight, BMIPower)
	fmt.Printf("Your BMI: %.0f", BMI)
	//	fmt.Print(BMI)
}
