package main

import (
	"fmt"
	"github.com/mind1949/getinput"
)

func calTax(orderAmount float64, state string) float64 {
	var taxRate float64
	if state == "WI" {
		taxRate = 5.5
	}
	tax := orderAmount * taxRate / 100
	return tax
}

func main() {
	_orderAmount := getinput.Call("What is the order amount?", "float64")
	_state := getinput.Call("What is the state?", "string")
	//orderAmount := _orderAmount.(float64)
	//state := _state.(string)

	tax := calTax(_orderAmount.(float64), _state.(string))
	subtotal := _orderAmount.(float64)
	total := subtotal + tax

	if tax > 0 {
		fmt.Println("The subtotal is $", subtotal)
		fmt.Println("The tax is $", tax)
	}
	fmt.Println("The total is $", total)
}
