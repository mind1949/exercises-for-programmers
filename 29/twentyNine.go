package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func prompt(msg string) (input string) {
	if !strings.HasSuffix(msg, " ") {
		msg = msg + " "
	}
	fmt.Print(msg)
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input)
	return
}

func main() {
	for {
		_input := prompt("What is the rate of return?")
		rate, err := strconv.ParseFloat(_input, 64)

		if err != nil {
			fmt.Println("Sorry. That's not a valid input")
			continue
		}

		if rate == 0 {
			fmt.Println("Rate cann't be zero!")
			continue
		}

		years := 72.0/rate
		fmt.Printf("It will take your %.0f years to double your initial investment.\n", years)
		break
	}
}
