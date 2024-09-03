package main

import "fmt"

func main() {
	order := Order{ID: 1, Amount: 100.0}
	err := ProcessOrder(order)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
