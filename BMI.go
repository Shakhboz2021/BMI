package main

import (
	"github.com/Shakhboz2021/BMI/info"
)

func main() {

	// Output information
	info.Printwelcome()

	weight, height := GetUserMetrics()

	bmi := calculateBMI(weight, height)

	printBMI(bmi)
}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}
