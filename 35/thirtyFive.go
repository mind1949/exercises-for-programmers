package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"math/rand"
	"time"
)

var participants []string

func prompt(msg string) (input string)  {
	if !strings.HasSuffix(msg, " ") {
		msg += " "
	}
	fmt.Print(msg)

	_input := bufio.NewScanner(os.Stdin)
	if _input.Scan() {
		input = _input.Text()
	}

	return
}

func getParticipants() {
	for {
		_participant := prompt("Enter a name:")
		if _participant != "" {
			participants = append(participants, _participant)
		} else {
			break
		}
	}
}

func selectWinner() (winner string) {
	rand.Seed(time.Now().UnixNano())

	winner = participants[rand.Intn(len(participants))]

	return
}

func main() {
	getParticipants()
	winner := selectWinner()
	fmt.Println("The winner is ...", winner)
}
