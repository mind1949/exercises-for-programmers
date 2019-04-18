package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Name struct {
	FirstName string
	LastName  string
}

func (name Name) fullName() (_fullName string) {
	return name.FirstName + " " + name.LastName
}

type User struct {
	Name
	Position       string
	SeparationData string
}

var users []User

func init() {
	name1 := Name{"John", "Johnson"}
	name2 := Name{"Tou", "Xiong"}
	name3 := Name{"Michaela", "Michaelson"}
	name4 := Name{"Jake", "Jacobon"}
	name5 := Name{"Jacquelyn", "Jackson"}
	name6 := Name{"Sally", "Weber"}

	user1 := User{name1, "Manager", "2016-12-31"}
	user2 := User{name2, "Software Engineer", "2016-10-05"}
	user3 := User{name3, "District Manager", "2015-12-19"}
	user4 := User{name4, "Programmer", ""}
	user5 := User{name5, "DBA", ""}
	user6 := User{name6, "Web Developer", "2015-12-08"}

	users = []User{user1, user2, user3, user4, user5, user6}
}

func prompt(msg string) (input string) {
	if !strings.HasSuffix(msg, " ") {
		msg += " "
	}
	fmt.Print(msg)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	return
}

func searchByName(searchString string) (matchedUsers []User) {
	for _, user := range users {
		if strings.Contains(user.fullName(), searchString) {
			matchedUsers = append(matchedUsers, user)
		}
	}
	return
}

func output(slice []User) {
	fmt.Println("Name\t| Position\t| SeparationData")
	fmt.Println("----------|----------|----------")
	for _, user := range slice {
		fmt.Printf("%s\t| %s\t| %s\n", user.fullName(), user.Position, user.SeparationData)
	}
}

func main() {
	searchString := prompt("Enter a search string:")
	matchedUsers := searchByName(searchString)
	output(matchedUsers)
}
