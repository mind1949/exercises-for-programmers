package main

import(
	"fmt"
	"github.com/mind1949/getinput"
)

func main() {
	var sum float64
	for i := 0; i <= 4; i++ {
		sum += getinput.Call("Enter a number:", "float64").(float64)
	}

	fmt.Printf("The total is %.2f\n", sum)
}
