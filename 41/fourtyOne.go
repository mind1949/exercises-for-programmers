package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputfile, err := os.Open("input.text")
	check(err)
	defer inputfile.Close()

	outputfile, err := os.Create("output.text")
	check(err)

	err = outputfile.Truncate(0)
	check(err)

	defer outputfile.Close()

	scanner := bufio.NewScanner(inputfile)

	var count int
	for isScan(scanner) {
		count++
		row := scanner.Text() + "\n"
		_, err := outputfile.Write([]byte(row))
		check(err)
	}

	if count > 0 {
		title := fmt.Sprintf("The total is %d names\n", count)
		_, err := outputfile.WriteAt([]byte(title), 0)
		check(err)
	}
}

func isScan(scanner *bufio.Scanner) bool {
		return scanner.Scan()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
