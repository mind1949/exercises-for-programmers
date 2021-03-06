package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Name struct{ FirstName, LastName string }

type Staff struct {
	Name
	Salary int
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
	_Salary, err := strconv.Atoi(rowSlice[2])
	staff.Salary = _Salary
	check(err)
	return
}

func output(staffs []Staff) {
	fmt.Println("Last First Salary\n------------")
	for _, staff := range staffs {
		fmt.Printf("%s %s %d\n", staff.LastName, staff.FirstName, staff.Salary)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
