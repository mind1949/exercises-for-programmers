package main

import (
	"fmt"
	"github.com/mind1949/getinput"
	"regexp"
)

type isValider interface {
	isValid() bool
}

type firstName string

func (_firstName firstName) isValid() bool {
	return len(_firstName) > 1
}

type lastName string

func (_lastName lastName) isValid() bool {
	return len(_lastName) > 1
}

type zip string

func (_zip zip) isValid() bool {
	regex := regexp.MustCompile(`\A[0-9]{2,6}\z`)
	return regex.MatchString(string(_zip))
}

type employId string

func (_employId employId) isValid() bool {
	regex := regexp.MustCompile(`\A[A-Za-z]{2}-[0-9]{4}\z`)
	return regex.MatchString(string(_employId))
}

func main() {
	_firstName := firstName(getinput.Call("Enter the first name:", "string").(string))
	_lastName := lastName(getinput.Call("Enter the last name:", "string").(string))
	_zip := zip(getinput.Call("Enter the Zip Code:", "string").(string))
	_employId := employId(getinput.Call("Enter the empley id:", "string").(string))

	elements := []isValider{_firstName, _lastName, _zip, _employId}

	msgsForInvalid := map[string]string{
		"main.firstName": "firstName(%s) must be longer than 1 !\n",
		"main.lastName": "lastName(%s) must be longer than 1 !\n",
		"main.zip": "zip(%s) must be numeric !\n",
		"main.employId": "employId(%s) is not a valid id\n" ,
	}

	for _, element := range elements {
		if !element.isValid() {
			t := fmt.Sprintf("%T",element)
			msg := msgsForInvalid[t]
			fmt.Printf(msg,element)
		}
	}
}
