package main

import (
	"fmt"
	"github.com/mind1949/getinput"
)

func main() {
	age := getinput.Call("Enter your age:","int").(int)
	restPlus := getinput.Call("Enter your rest plus:","int").(int)

	cal := func(age int, restPlus int, intensity int) (heartRate int) {
		return (((220 - age) - restPlus)*intensity/100)+restPlus
	}

	for intensity := 55; intensity <= 95; intensity += 5 {
		if intensity == 55 {
			fmt.Printf("Intensity\t| Rate\n")
			fmt.Printf("---------\t|--------\n")
		}
		fmt.Printf("   %d%%   \t| %d bpm\n", intensity, cal(age, restPlus, intensity))
	}
}
