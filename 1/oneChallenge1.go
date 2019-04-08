package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func mu(a ...interface{}) []interface{} {
	return a
}

func getInput() (string, error) {
	fmt.Print("wha is your name: ")
	return bufio.NewReader(os.Stdin).ReadString('\n')
}

func main() {
	fmt.Println("Hello, ", strings.Replace(mu(getInput())[0].(string), "\n","",-1), " nice to meet you!")
}
