package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\t")
	for i := 0; i <= 12; i++ {
		fmt.Printf("%d\t", i)
		if i == 12 {
			fmt.Printf("\n")
		}
	}

	for i := 0; i <= 12; i++ {
		fmt.Printf("%d\t", i)
		for j := 0; j <= 12; j++ {
			fmt.Printf("%d\t", i*j)
			if j == 12 {
				fmt.Printf("\n")
			}
		}
	}
}
