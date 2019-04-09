package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"strings"
)

func getfloatinput(ques string) (float64) {
	var returninput float64
	for done := false; !done; {
		fmt.Print(ques)
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		finput, err := strconv.ParseFloat(input, 64)
		if err !=nil {
			continue
		}
		if finput > 0 {
			done = true
			returninput = finput
		}
	}
	return returninput
}

func main() {
	f1 := getfloatinput("what is first number? ")
	f2 := getfloatinput("what is second number? ")

	fmt.Printf("%f + %f = %f\n", f1, f2, f1 + f2)
	fmt.Printf("%f - %f = %f\n", f1, f2, f1 - f2)
	fmt.Printf("%f * %f = %f\n", f1, f2, f1 * f2)
	fmt.Printf("%f / %f = %f\n", f1, f2, f1 / f2)
}
