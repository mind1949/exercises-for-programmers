package main

import (
	"fmt"
	"github.com/mind1949/getinput"
)

type question string

func (q question) mark() (a bool) {
	var input string
	_q := fmt.Sprintf("%v?", q)
	for {
		input = getinput.Call(_q, "string").(string)
		if input == "y" || input == "n" {
			break
		}
	}

	m := map[string]bool{"y": true, "n": false}
	a = m[input]
	return
}

func main() {
	var q question
	q = "Is the car silent when you turn"
	if q.mark() {
		q = "Are the battery terminals corroded"
		if q.mark() {
			fmt.Println("Clean the terminals and try starting again.")
		} else {
			fmt.Println("Replace the cables and try again.")
		}
	} else {
		q = "Does the car make a clicking noise"
		if q.mark() {
			fmt.Println("Replace the battery.")
		} else {
			q = "Does the car crank up but fail to start"
			if q.mark() {
				fmt.Println("Check spark plug connection.")
			} else {
				q = "Does the engine start and then die"
				if q.mark() {
					q = "Does your car have fuel injection"
					if q.mark() {
						fmt.Println("Get it for service.")
					} else {
						fmt.Println("Check to ensure choke is openning and closing.")
					}
				} else {
					fmt.Println("I'm sorry , we are powerless.")
				}
			}
		}
	}
}
