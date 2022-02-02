package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/shalveena/bmicalculator/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserWeightAndHeight() (height float64, weight float64) {

	height = getUserInput(info.HeightPrompt)
	weight = getUserInput(info.WeightPrompt)
	return
}

func getUserInput(promtText string) (userInputNumber float64) {

	fmt.Print(promtText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInputNumber, _ = strconv.ParseFloat(userInput, 64)
	return
}
