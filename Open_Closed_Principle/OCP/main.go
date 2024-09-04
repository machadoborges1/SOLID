package main

import "fmt"

func main() {
	amount := 100.0

	// Usando um cliente Premium
	customer := PremiumCustomer{}
	discount := customer.CalculateDiscount(amount)
	fmt.Printf("Discount: %.2f\n", discount)
}
