package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight float64
	var userWeight float64
	fmt.Println("Body mass index calculation.")
	fmt.Print("Type in your height (sm): ")
	fmt.Scan(&userHeight)
	fmt.Print("Type in your weight: ")
	fmt.Scan(&userWeight)
	BMI := calculateBMI(userWeight, userHeight)
	outputResult(BMI)
}

func outputResult(bmi float64) {
	result := fmt.Sprintf("Your BMI: %.0f", bmi)
	fmt.Print(result)
}

func calculateBMI(userWeight float64, userHeight float64) float64 {
	const BMIPower = 2
	BMI := userWeight / math.Pow(userHeight/100, BMIPower)
	return BMI
}
