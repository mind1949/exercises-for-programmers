package main

import (
	"fmt"
	"github.com/mind1949/getinput"
	"crypto/sha256"
)


func encrypteData(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	encryptedData := fmt.Sprintf("%x",h.Sum(nil))
	return encryptedData
}

func main() {
	users := make(map[string]string,2)
	users["mindz"] = encrypteData("1949")
	users["pondz"] = encrypteData("1949")

	username := getinput.Call("Enter your name:", "string").(string)
	password := getinput.Call("Enter your password:", "string").(string)

	encryptedPassword, ok := users[username]

	if !ok {
		fmt.Println("I don't know you!")
		return
	}
	if encryptedPassword != encrypteData(password) {
		fmt.Println("Password doesn't match!")
	}
	if encryptedPassword == encrypteData(password) {
		fmt.Printf("Welcome %s!\n",username)
	}
}
