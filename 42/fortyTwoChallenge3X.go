package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Name struct{ FirstName, LastName string }

type dollar int

func (d dollar) String() string {
	_slice := []rune(strconv.Itoa(int(d)))
	reverse(_slice)

	var slice []rune
	for i, v := range _slice {
		slice = append(slice, v)
		if i%3 == 0 && i > 0 && i < len(_slice)-1 {
			slice = append(slice, ',')
		}
	}
	reverse(slice)

	retS := "$" + string(slice)
	return retS

}

type Staff struct {
	Name
	Salary dollar
}

func main() {
	file, err := os.Open("data.csv")
	check(err)
	defer file.Close()

	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	var staffs = make([]Staff, len(rows))
	for i, row := range rows {
		// debug
		fmt.Println(row)
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

func parseRow(row []string) (staff Staff) {
	staff.LastName = row[0]
	staff.FirstName = row[1]

	_salary, err := strconv.Atoi(strings.TrimSpace(row[2]))
	check(err)
	staff.Salary = dollar(_salary)

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
		fmt.Printf("%s %s %s\n", staff.LastName, staff.FirstName, staff.Salary)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
