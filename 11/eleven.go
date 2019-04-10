package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//	"math"
)

func getFloatInput(ques string) (retInput float64) {
	for undone := true; undone; {
		fmt.Print(ques)
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			continue
		}
		input = strings.TrimSpace(input)
		retInput, err = strconv.ParseFloat(input, 64)
		if err != nil  || retInput <= 0 {
			continue
		}
		undone = false
	}
	return
}

func main() {
	amountFrom := getFloatInput("How many euros are you exchanging?")
	rateFrom := getFloatInput("What is the exchange rate?")
	amountTo := amountFrom * rateFrom / 100
	// ToDo: 保留向上保留两位小数
	fmt.Printf("%.2f euros at exchange rate of %.3f is %.3f U.S. dollars", amountFrom, rateFrom, amountTo)
}
