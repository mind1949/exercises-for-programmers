package main

import (
	"bufio"
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

	for isScan(scanner) {
		row := scanner.Text() + "\n"
		_, err := outputfile.Write([]byte(row))
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
