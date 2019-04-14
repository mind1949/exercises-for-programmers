package main

import (
	"fmt"
	"github.com/mind1949/getinput"
	"math"
)

func calculateMonthsUntilPaidOff(balance, APR, paymentPerMonth float64) int {
	dailyRate := APR / 100 / 356
	divisor := math.Log10(1+(balance/paymentPerMonth)*(1-math.Pow((1+dailyRate), 30)))
	dividend := math.Log10(1+dailyRate)
	months := int(math.Ceil((-1.0/30)*(divisor/dividend)))
	return months
}

func main() {
	balance := getinput.Call("What is your balance?","float64").(float64)
	APR := getinput.Call("What is the APR one the card(as a percent)?","float64").(float64)
	paymentPerMonth := getinput.Call("What is the monthly payment you can make?","float64").(float64)

	months := calculateMonthsUntilPaidOff(balance, APR, paymentPerMonth)

	fmt.Printf("It will take you %d months to pay off this card.\n", months)
}
