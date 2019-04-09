package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func getInput(question string) (answer string) {
	for done := false; !done; {
		fmt.Print(question)
		answer, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		answer = strings.TrimSpace(answer)

		if len(answer) > 0 {
			done = true
		}
	}
	return
}

func main() {
	quote := getInput("What is the quote ?")
	who := getInput("Who said it ?")

	fmt.Printf("%s says, %q\n", who, quote)
}


