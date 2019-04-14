package main

import (
	"fmt"
	"github.com/mind1949/getinput"
)

func validate(username, password string) string {
	users := map[string]string{"pondz":"1949","mindz":"1949"}
	_password, ok := users[username]
	if !ok {
		return "I don't know you!"
	}
	if _password != password {
		return "Password doesn't match!"
	}
	return fmt.Sprintf("Welcome %s", username)
}

func main() {
	username := getinput.Call("What is your name?", "string")
	password := getinput.Call("What is your password?", "string")

	result := validate(username.(string), password.(string))
	fmt.Println(result)
}
