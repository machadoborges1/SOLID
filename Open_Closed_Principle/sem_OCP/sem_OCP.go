//A função CalculateDiscount precisa ser alterada toda vez que um novo tipo de cliente for adicionado.

package main

import "fmt"

// Tipo de cliente
const (
	RegularCustomer = iota
	PremiumCustomer
	VIPCustomer
)

// Calcula desconto com base no tipo de cliente
func CalculateDiscount(customerType int, amount float64) float64 {
	switch customerType {
	case RegularCustomer:
		return amount * 0.05
	case PremiumCustomer:
		return amount * 0.10
	case VIPCustomer:
		return amount * 0.15
	default:
		return 0
	}
}

func main() {
	amount := 100.0
	discount := CalculateDiscount(PremiumCustomer, amount)
	fmt.Printf("Discount: %.2f\n", discount)
}
