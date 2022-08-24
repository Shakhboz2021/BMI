package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Shakhboz2021/BMI/info"
)

func main() {
	fmt.Println(info.MainTitle)

	fmt.Println(info.Separator)

	fmt.Print(info.WeightPrompt)
	weigthInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	weigthInput = strings.Replace(weigthInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weigthInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := weight / (height * height)

	fmt.Printf("Your BMI is %.2f", bmi)
}
