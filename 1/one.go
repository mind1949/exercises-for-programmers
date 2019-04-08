package main

import "fmt"

func main() {
	fmt.Print("What is your name: ")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("Hello, ", name, "nice to meet you!")
}
