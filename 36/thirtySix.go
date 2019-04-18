package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var numbers []float64

func appendNumbers() {
	for {
		fmt.Print("Enter a number:")
		_input := bufio.NewScanner(os.Stdin)
		if _input.Scan() {
			input := _input.Text()
			if input == "done" {
				break
			}
			number, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			numbers = append(numbers, number)
		}
	}
}

func mean() float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func standardDeviation() float64 {
	average := mean()
	sum := 0.0
	for _, number := range numbers {
		sum += math.Pow((number - average), 2.0)
	}
	return math.Pow((sum / float64(len(numbers))), 0.5)
}

func min() float64 {
	var _min = numbers[0]
	for _, number := range numbers {
		if _min > number {
			_min = number
		}
	}
	return _min
}

func max() float64 {
	_max := numbers[0]
	for _, number := range numbers {
		if _max < number {
			_max = number
		}
	}
	return _max
}

func joinNumbers() string {
	strs := make([]string, len(numbers))
	for index, number := range numbers {
		strs[index] = strconv.FormatFloat(number, 'f', -1, 64)
	}
	return strings.Join(strs, ", ")
}

func output() {
	fmt.Println("The number is: ", joinNumbers())
	fmt.Printf("The average is %.2f.\n", mean())
	fmt.Printf("The minmum is %.2f.\n", min())
	fmt.Printf("The maxmum is %.2f.\n", max())
	fmt.Printf("The standard devitation is %f.\n", standardDeviation())
}

func main() {
	appendNumbers()
	output()
}
