package main

import (
	"fmt"
	"github.com/mind1949/getinput"
)

const legalYears = 16

func main() {
	userYears := getinput.Call("Enter your years:", "int").(int)
	for userYears <= 0 {
		userYears = getinput.Call("Enter your years(it should be greater than 0):", "int").(int)
	}

	if userYears >= legalYears {
		fmt.Println("You are old enough to legally drive.")
	}
	if userYears < legalYears {
		fmt.Println("You are not old enough to legally drive.")
	}
}
