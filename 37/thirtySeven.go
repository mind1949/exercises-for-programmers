package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type dataSource struct {
	Numbers          []rune
	SpecialCharaters []rune
}

var ds = dataSource{
	Numbers:          []rune("0123456789"),
	SpecialCharaters: []rune("!@#$%^&*?:"),
}

func prompt(msg string) (input string) {
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

func getcount(msg string) (count int) {
	_count := prompt(msg)
	count, err := strconv.Atoi(_count)
	if err != nil {
		fmt.Println(err)
	}
	return count
}

func charatersSlice(count int, t string) (_charatersSlice []rune) {
	dataSlices := map[string][]rune{
		"numbers":          ds.Numbers,
		"specialCharaters": ds.SpecialCharaters,
	}
	dataSlice := dataSlices[t]

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= count; i++ {
		number := dataSlice[r.Intn(len(dataSlice))]
		_charatersSlice = append(_charatersSlice, number)
	}

	return
}

func passwordSlice(numbersCount, specialCharatersCount int) (_passwordSlice []rune) {
	_numbersSlice := charatersSlice(numbersCount, "numbers")
	_specialCharatersSlice := charatersSlice(specialCharatersCount, "specialCharaters")
	_passwordSlice = append(_numbersSlice, _specialCharatersSlice...)
	return
}

func shuffleSlice(slice []rune) (retSlice []rune) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := len(slice)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(slice))
		retSlice = append(retSlice, slice[randIndex])
		slice = append(slice[:randIndex], slice[randIndex+1:]...)
	}
	return
}

func genPassword(numbersCount, specialCharatersCount int) (password string) {
	_passwordSlice := passwordSlice(numbersCount, specialCharatersCount)
	_password := shuffleSlice(_passwordSlice)
	password = string(_password)
	return
}

func main() {
	minLength := getcount("Whart's the minimum length?")
	specialCharatersCount := getcount("How many special charaters ?")
	numbersCount := getcount("How many numbers ?")

	if minLength < specialCharatersCount+numbersCount {
		fmt.Println("minimu length must be greater than special charaters count and numbers count!")
		os.Exit(1)
	}

	password := genPassword(numbersCount, specialCharatersCount)

	fmt.Println("your password is:\n", password)
}
