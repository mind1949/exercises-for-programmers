package main

//What is the length of the room in feet? 15
//What is the width of the room in feet? 20
//You entered dimensions of 15 feet by 20 feet.
//The area is
//300 square feet
//27.871 square meters
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const rate = 0.09290304

func getInput(question string) (answer float64) {
	for done := false; !done; {
		fmt.Print(question)
		_answer, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			continue
		}
		_answer = strings.TrimSpace(_answer)
		_answer1, _ := strconv.ParseFloat(_answer, 64)
		if err != nil || _answer1 <= 0 {
			continue
		}
		done = true
		answer = _answer1
	}
	return answer
}

func getInputUnit() string {
	fmt.Print("use meter or feet as unit?(please input meter or feet)")
	answer, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return getInputUnit()
	}
	answer = strings.TrimSpace(answer)

	switch answer {
	case "meter":
		return "meter"
	case "feet":
		return "feet"
	default:
		return getInputUnit()
	}
}

func main() {
	unit := getInputUnit()
	length := getInput(fmt.Sprintf("What is the length of the room in %s? ", unit))
	width := getInput(fmt.Sprintf("What is the width of the room in %s? ", unit))
	area := length * width

	fmt.Printf("You entered dimensions of %f %s by %f %s.\n", length, unit, width, unit)
	var xArea float64
	if unit == "meter" {
		xArea = area / rate
	} else {
		xArea = area * rate
	}
	fmt.Printf("The area is %f sqare feet, %f square meters\n", area, xArea)
}
