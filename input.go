package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Shakhboz2021/BMI/info"
)

var reader = bufio.NewReader(os.Stdin) // func assigments cannot be a const because it must run main()

func GetUserMetrics() (weight float64, height float64) {

	weight = getUserInput(info.WeightPrompt)
	height = getUserInput(info.HeightPrompt)

	return
}

func getUserInput(promtText string) (value float64) {

	fmt.Print(promtText)

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	value, _ = strconv.ParseFloat(userInput, 64)

	return
}
