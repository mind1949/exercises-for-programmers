package main

import (
	"fmt"
)

func getInput(question string) (answer string) {
	fmt.Print(question)
	fmt.Scanf("%s", &answer)
	fmt.Println(answer)
	return
}

func main() {
	templ := "Do you %s your %s %s %s? That is hilarious!\n"
	verb := getInput("Enter a verb: ")
	adjective := getInput("Enter a adjective: ")
	noun := getInput("Enter a noun: ")
	adverb := getInput("Enter a adverb: ")
	fmt.Printf(templ, verb, adjective, noun, adverb)
}
