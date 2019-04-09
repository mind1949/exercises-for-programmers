package main

import (
	"fmt"
)

type people struct {
	Name string
	Say string
}

func (p *people) say() {
	fmt.Println(p.Name, "says " + "\"" + p.Say + "\"")
}

func main() {
	peoples := []people{
		people{"mindz", "pondz, i love you !"},
		people{"pondz", "mindz, i love you !"},
	}

	for _, p := range peoples {
		p.say()
	}
}
