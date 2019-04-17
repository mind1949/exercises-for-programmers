package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func prompt(msg string) {
	if strings.HasSuffix(msg, " ") {
		msg += " "
	}
	fmt.Printf(msg)

	_input := bufio.NewScanner(os.Stdin)
	if _input.Scan() {
		input := _input.Text()
		if len(input) <= 0 {
			prompt(msg)
		}
	} else {
		prompt(msg)
	}
}

var ANSWERS = []string{
	"Yes",
	"No",
	"Maybe",
	"Ask Again later",
}

func magicEight() {
	prompt("What's you want know?")

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(ANSWERS))

	answer := ANSWERS[index]
	fmt.Println(answer)

	if index > 2 {
		magicEight()
	}
}

func main() {
	magicEight()
}
