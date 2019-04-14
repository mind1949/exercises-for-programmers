package main

import (
	"fmt"
	"github.com/mind1949/getinput"
	"strings"
)

var countriesLegalYears = map[string]int{"china": 18, "america": 16}

func legalcountries(years int) (countries []string) {
	for country, _years := range countriesLegalYears {
		if years >= _years {
			countries = append(countries, country)
		}
	}
	return
}

func main() {
	years := getinput.Call("Enter years:", "int").(int)
	for years <= 0 {
		years = getinput.Call("enter years(it should be greater than 0):", "int").(int)
	}

	countries := legalcountries(years)

	if len(countries) > 0 {
		fmt.Println("you can legal drive at ", strings.Join(countries, " | "))
	} else {
		fmt.Println("you are too young to legal drive")
	}
}
