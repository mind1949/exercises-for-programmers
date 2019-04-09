package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() string {
	fmt.Print("What is the input string?")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimRight(input, "\n")
}

func main() {
	input := getInput()
	fmt.Printf("'%s' has %d charaters.\n", input, len(input))
}
