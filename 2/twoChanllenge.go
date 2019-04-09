package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() string {
	var input string
	var getinput (func(input string) string)

	getinput = func(input string) string {
		fmt.Print("What is the input string?")
		input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		if len(input) == 0 || input == "\n" {
			return getinput(input)
		}
		fmt.Println("your input is", input)
		return input
	}

	input = getinput(input)
	return strings.TrimRight(input, "\n")
}

func main() {
	input := getInput()
	fmt.Printf("'%s' has %d charaters.\n", input, len(input))
}
