package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("BMI calculator")

	fmt.Println("--------------------")

	fmt.Print("Please, enter your weight (kg) ")
	weigthInput, _ := reader.ReadString('\n')

	fmt.Print("Please, enter your height (m) ")
	heightInput, _ := reader.ReadString('\n')

	weigthInput = strings.Replace(weigthInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weigthInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	fmt.Print(weight)
	fmt.Print(height)
}
