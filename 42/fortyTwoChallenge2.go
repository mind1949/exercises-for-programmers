package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	BySalary(staffs).Sort()

	output(staffs)
}

type BySalary []Staff

func (bs BySalary) Len() int {
	return len(bs)
}

func (bs BySalary) Less(i, j int) bool {
	return bs[i].Salary > bs[j].Salary
}

func (bs BySalary) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func (bs BySalary) Sort() {
	sort.Sort(bs)
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

	_salary, err := strconv.Atoi(rowSlice[2])
	check(err)
	staff.Salary = _salary

	return
}

func dollarize(s int) (retS string) {
	_slice := []rune(strconv.Itoa(s))
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

func reverse(s []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func output(staffs []Staff) {
	fmt.Println("Last First Salary\n------------")
	for _, staff := range staffs {
		fmt.Printf("%s %s %s\n", staff.LastName, staff.FirstName, dollarize(staff.Salary))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
