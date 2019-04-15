package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func promptAndGetInput(msg string) (input float64) {
	if !strings.HasSuffix(msg, " ") {
		msg = msg + " "
	}
	fmt.Print(msg)

	_input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	_input = strings.TrimSpace(_input)
	input, err = strconv.ParseFloat(_input, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func main() {
	count := promptAndGetInput("Enter a count:")
	var sum float64
	for ; count >= 1.0; count-- {
		sum += promptAndGetInput("Enter a number:")
	}

	fmt.Printf("The total is %.2f\n", sum)
}

