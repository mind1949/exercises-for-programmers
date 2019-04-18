package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func prompt(msg string) (input string) {
	if !strings.HasSuffix(msg, " ") {
		msg += " "
	}
	fmt.Print(msg)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	return
}

func sliceAtoi(slice []string) (retSlice []int) {
	for _, v := range slice {
		retV, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		retSlice = append(retSlice, retV)
	}
	return
}

func isEven(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}

func retEvenSlice(slice []int) (retSlice []string) {
	for _, v := range slice {
		if isEven(v) {
			retSlice = append(retSlice, strconv.Itoa(v))
		}
	}
	return
}
func main() {
	numbersString := prompt("Enter a list of numbers, seperated by space:")
	_numbersSlice := strings.Split(numbersString, " ")
	numbersSlice := sliceAtoi(_numbersSlice)

	EvenSlice := retEvenSlice(numbersSlice)

	fmt.Printf("The Even numbers is %s \n", strings.Join(EvenSlice, " "))
}
