package main

import (
	"fmt"
	"sort"
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

type byName []User

func (us byName) Len() int { return len(us) }
func (us byName) Less(i, j int) bool {
	return us[i].fullName() < us[j].fullName()
}
func (us byName) Swap(i, j int) {
	us[i], us[j] = us[j], us[i]
}

func (us byName) Sort() { sort.Sort(us) }

func output() {
	fmt.Println("Name\t| Position\t| Separation Date")
	fmt.Println("----\t|---------\t|----------------")
	for _, user := range users {
		fmt.Printf("%s\t | %s\t | %s\t\n", user.fullName(), user.Position, user.SeparationData)
	}
}

func main() {
	byName(users).Sort()
	output()
}
