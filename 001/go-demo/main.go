package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Body mass index calculation.")
	for {

		userHeight, userWeight := getUserInput()
		BMI, err := calculateBMI(userHeight, userWeight)
		if err != nil {
			fmt.Println("Please write correct height or weight")
			continue
			// panic("Please write correct height or weight")
		}
		outputResult(BMI)
		isRepeateCalcalation := checkRepeatCalculation()
		if !isRepeateCalcalation {
			break
		}
	}
}

func outputResult(bmi float64) {
	result := fmt.Sprintf("Your BMI: %.0f", bmi)
	fmt.Println(result)
	switch {
	case bmi < 16:
		fmt.Println("You're severely underweight")
	case bmi < 18.5:
		fmt.Println("You're underweight")
	case bmi < 25:
		fmt.Println("You have a normal body weight")
	case bmi < 30:
		fmt.Println("You are overweight")
	case bmi >= 30:
		fmt.Println("You are obese")
	}
}

func calculateBMI(userHeight float64, userWeight float64) (float64, error) {
	if userHeight <= 0 || userWeight <= 0 {
		return 0, errors.New("incorrect height or weight")
	}
	const BMIPower = 2
	BMI := userWeight / math.Pow(userHeight/100, BMIPower)
	return BMI, nil
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

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Do you want to do another calculation (Y/N)?")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
