package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin) // func assigments cannot be a const because it must run main()

const mainTitle = "BMI calculator"
const separator = "--------------------"
const weightPrompt = "Please, enter your weight (kg) "
const heightPrompt = "Please, enter your height (m) "

func main() {
	fmt.Println(mainTitle)

	fmt.Println(separator)

	fmt.Print(weightPrompt)
	weigthInput, _ := reader.ReadString('\n')

	fmt.Print(heightPrompt)
	heightInput, _ := reader.ReadString('\n')

	weigthInput = strings.Replace(weigthInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weigthInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := weight / (height * height)

	fmt.Printf("Your BMI is %.2f", bmi)
}
