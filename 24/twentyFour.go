package main

import (
	"fmt"
	"github.com/mind1949/getinput"
)

func anagram(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == anagram(s2) {
		return true
	}
	return false
}

func main() {
	fmt.Println("Enter two strings and i'll tell you if they are anagram.")
	s1 := getinput.Call("Enter the first string:","string").(string)
	s2 := getinput.Call("Enter the second string:","string").(string)

	if isAnagram(s1, s2) {
		fmt.Printf("%s and %s are anagrams.\n", s1, s2)
	} else {
		fmt.Printf("%s and %s aren't anagrams.\n", s1, s2)
	}
}
