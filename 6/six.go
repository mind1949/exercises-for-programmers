package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getInput(question string) (input int) {
	for done := false; !done; {
		fmt.Print(question)
		answer, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		answer = strings.TrimSpace(answer)
		forinput, err := strconv.ParseInt(answer, 10, 0)
		if err != nil || forinput <= 0 {
			continue
		}
		done = true
		input = forinput
	}
	return input
}

func retirementYear(age int, retirementAget int) int {
	afterYears := retirementAget - age
	retirementYear := time.Now().AddDate(afterYears, 0, 0).Year()
	return retirementYear
}

func main() {
	age := getInput("What is your current age? ")
	retirementage := getInput("At what age whould you like to retire? ")
	fmt.Println("You have ", retirementage-age, " years left until you retire.")
	fmt.Printf("It's %d, so you can retire at %d\n", time.Now().Year(), retirementYear(age, retirementage))
}
