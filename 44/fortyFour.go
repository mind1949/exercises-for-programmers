package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func main() {
	productName := prompt("What is the product name?")

	jsonfile, err := os.Open("data.json")
	check(err)
	defer jsonfile.Close()

	dec := json.NewDecoder(jsonfile)

	ps := map[string][]Product{}
	products := []Product{}
	ps["products"] = products
	dec.Decode(&ps)

	product, ok := isContain(ps["products"], productName)
	if !ok {
		fmt.Println("Sorry, that product was not found in our inventory.")
		os.Exit(1)
	}

	fmt.Printf("Name\t%s\nPrice\t%.4f\nQuantity\t%d\n", product.Name, product.Price, product.Quantity)
}

func isContain(ps []Product, name string) (p Product, ok bool) {
	for _, _p := range ps {
		if _p.Name == name {
			p = _p
			ok = true
		}
	}

	return
}

func prompt(msg string) (input string) {
	if !strings.HasSuffix(msg, " ") {
		msg += " "
	}
	fmt.Printf(msg)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
