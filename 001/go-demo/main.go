package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower = 2
	userHeight := 1.8
	userWeight := 60.0
	BMI := userWeight / math.Pow(userHeight, BMIPower)

	fmt.Println("Body mass index calculation.")
	fmt.Print(BMI)
}
