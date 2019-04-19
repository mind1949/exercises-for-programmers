package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct{ FirstName, LastName string }

type Staff struct {
	Name
	Salary string
}

func main() {
	rows := readAndParse("data.csv")
	rows = filter(rows)
	var staffs = make([]Staff, len(rows))
	for i, row := range rows {
		staffs[i] = parseRow(row)
	}

	output(staffs)
}

func filter(slice []string) []string {
	retSlice := make([]string, 0, len(slice))
	for _, element := range slice {
		if strings.Count(element, ", ") == 2 {
			retSlice = append(retSlice, element)
		}
	}
	return retSlice
}

func readAndParse(fileName string) (rows []string) {
	file, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return
}

func parseRow(row string) (staff Staff) {
	rowSlice := strings.Split(row, ", ")
	if len(rowSlice) < 3 {
		return
	}
	staff.LastName = rowSlice[0]
	staff.FirstName = rowSlice[1]
	staff.Salary = dollarize(rowSlice[2])

	return
}

func dollarize(s string) (retS string) {
	_slice := []rune(s)
	reverse(_slice)

	var slice []rune
	for i, v := range _slice {
		slice = append(slice, v)
		if i%3 == 0 && i > 0 && i < len(_slice)-1 {
			slice = append(slice, ',')
		}
	}
	reverse(slice)

	retS = "$" + string(slice)
	return
}

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func output(staffs []Staff) {
	fmt.Println("Last First Salary\n------------")
	for _, staff := range staffs {
		fmt.Printf("%s %s %s\n", staff.LastName, staff.FirstName, staff.Salary)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
