package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.8
	userWeight := 60.0
	BMI := userWeight / math.Pow(userHeight, 2)

	fmt.Println("Body mass index calculation.")
	fmt.Print(BMI)
}
