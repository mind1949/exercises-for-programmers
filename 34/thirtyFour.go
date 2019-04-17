package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func prompt(msg string) (input string) {
	if !strings.HasSuffix(msg, " ") {
		msg += " "
	}
	fmt.Printf(msg)

	_input := bufio.NewScanner(os.Stdin)

	if _input.Scan() {
		input = _input.Text()
	} else {
		return prompt(msg)
	}

	return
}

var employees = []string{
	"zhao",
	"qian",
	"sun",
	"li",
	"zhou",
}

func outputEmployees() {
	count := len(employees)
	fmt.Printf("There are %d employees.\n", count)
	for _, employee := range employees {
		fmt.Println(employee)
	}
}

func deleteEmployee(employee string) {
	findIndex := func(s []string, element string) (index int) {
		index = -1
		for _index, value := range s {
			if value == element {
				index = _index
				break
			}
		}
		return
	}

	remove := func(s []string, index int) []string {
		copy(s[index:], s[index+1:])
		return s[:len(s)-1]
	}

	employeeIndex := findIndex(employees, employee)
	employees = remove(employees, employeeIndex)
}

func main() {
	outputEmployees()

	employeeName := prompt("enter an employee name to remove: ")
	deleteEmployee(employeeName)

	fmt.Println("")

	outputEmployees()
}
