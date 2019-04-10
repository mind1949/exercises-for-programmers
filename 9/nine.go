package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

const SquareFeetCoveredByOneGallon = 350

func getIntInput(question string) (retInput int) {
	fmt.Print(question)
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return getIntInput(question)
	}
	input = strings.TrimSpace(input)
	retInput, err = strconv.Atoi(input)

	if err != nil {
		fmt.Println(err)
		return getIntInput(question)
	}
	return retInput
}

func main() {
	length := getIntInput("Length:")
	width := getIntInput("Width:")

	area := length * width

	GallonsNeeded := int(math.Ceil(float64(area) / SquareFeetCoveredByOneGallon))

	fmt.Printf("You need %d gallons of paint to cover %d square feet.\n", GallonsNeeded, area)
}
