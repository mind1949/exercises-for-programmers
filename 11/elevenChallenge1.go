package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

var (
	COUNTRIES = []string{"eurios", "china",}
	EXCHANGE_RATE = map[string]float64{
		COUNTRIES[0]: 127.123455,
		COUNTRIES[1]: 140.122353,
	}
)

func getCountry() string {
	fmt.Print("Country: ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return getCountry()
	}
	input = strings.TrimSpace(input)
	var validCountry bool
	for _, country := range COUNTRIES {
		if country == input {
			validCountry = true
		}
	}
	if validCountry {
		return input
	}
	return getCountry()
}

func getAmount() float64 {
	fmt.Print("Amount: ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return getAmount()
	}
	input = strings.TrimSpace(input)
	retInput, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
		return getAmount()
	}
	if retInput > 0.0 {
		return retInput
	}
	return getAmount()
}

func calAmountTo(country string, amountfrom float64) float64 {
	return EXCHANGE_RATE[country] * amountfrom / 100
}

func fireOutput(country string, amountFrom float64, amountTo float64) {
	fmt.Println(amountFrom ," amount at ", country, " exchange to ", amountTo)
}

func main() {
	country := getCountry()
	amount := getAmount()
	amountTo := calAmountTo(country, amount)
	fireOutput(country, amount, amountTo)
}
