package main

import (
	"fmt"
	"github.com/mind1949/getinput"
	"strings"
)

func BAC(a, w, r, h float64) (bac float64) {
	bac = (a*5.14/w*r)-0.015*h
	return
}

func getPositiveInput(msg string) (retInput float64) {
	retInput = getinput.Call(msg, "float64").(float64)
	for ; retInput <= 0; {
		retInput = getinput.Call(strings.TrimRight(msg,":")+"(it should be greater than 0):", "float64").(float64)
	}
	return
}

func main() {
	var a, w, r, h float64
	a = getPositiveInput("A:")
	w = getPositiveInput("W:")
	r = getPositiveInput("r:")
	h = getPositiveInput("h:")

	bac := BAC(a, w, r, h)

	fmt.Println("Your BAC is ", bac)
	if bac >= 0.08 {
		fmt.Println("It is not legal for you to drive.")
	} else {
		fmt.Println("It is legal for you to drive.")
	}
}
