package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Body mass index calculation.")
	userHeight, userWeight := getUserInput()
	BMI := calculateBMI(userHeight, userWeight)
	isLean := BMI < 16
	if isLean {
		fmt.Println("You're underweight")
	}
	outputResult(BMI)
}

func outputResult(bmi float64) {
	result := fmt.Sprintf("Your BMI: %.0f", bmi)
	fmt.Print(result)
}

func calculateBMI(userHeight float64, userWeight float64) float64 {
	const BMIPower = 2
	BMI := userWeight / math.Pow(userHeight/100, BMIPower)
	return BMI
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Type in your height (sm): ")
	fmt.Scan(&userHeight)
	fmt.Print("Type in your weight: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}
