package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getIntInput(msg string) (retInput int) {
	for undone := true; undone; {
		if !strings.HasSuffix(msg, " ") {
			msg = msg + " "
		}
		fmt.Print(msg)
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			continue
		}
		input = strings.TrimSpace(input)
		retInput, err = strconv.Atoi(input)
		if err != nil {
			continue
		}
		undone = false
	}
	return
}

type item struct {
	Price    int
	Quantity int
}

func main() {
	items := []item{3: item{}}
	for i := 1; i < 4; i++ {
		items[i].Price = getIntInput(fmt.Sprintf("item %d price:", i))
		items[i].Quantity = getIntInput(fmt.Sprintf("item %d quantity:", i))
	}
	var subtotal float64
	for i := 1; i < 4; i++ {
		subtotal = subtotal + float64(items[i].Price*items[i].Quantity)
	}
	Tax := subtotal * 5.0 / 100
	Total := subtotal - Tax
	fmt.Println("Subtotal: ", subtotal)
	fmt.Println("Tax: ", Tax)
	fmt.Println("Total: ", Total)
}
