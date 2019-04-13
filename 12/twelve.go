package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func getFloatInput(msg string) (retInput float64) {
	if !strings.HasSuffix(msg, " ") {
		msg = msg + " "
	}
	fmt.Print(msg)
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return getFloatInput(msg)
	}
	input = strings.TrimSpace(input)
	retInput, err = strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
		return getFloatInput(msg)
	}
	return retInput
}

func calInvestment(years float64, rate float64, principal float64) (investment float64) {
	investment = principal*(1 + rate/100*years)
	return
}

func Output(years float64, rate float64, investment float64) {
	fmt.Printf("after %.0f year at %.2f%%, the investment will be worth $%.2f.\n", years, rate, investment)
}

func main() {
	principal := getFloatInput("Principal:")
	years := getFloatInput("years:")
	rate := getFloatInput("rate:")

	investment := calInvestment(years, rate, principal)

	Output(years, rate, investment)
}
