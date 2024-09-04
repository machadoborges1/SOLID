package main

// Customer define a interface para calcular o desconto
type Customer interface {
	CalculateDiscount(amount float64) float64
}

// RegularCustomer implementa Customer
type RegularCustomer struct{}

func (c RegularCustomer) CalculateDiscount(amount float64) float64 {
	return amount * 0.05
}

// PremiumCustomer implementa Customer
type PremiumCustomer struct{}

func (c PremiumCustomer) CalculateDiscount(amount float64) float64 {
	return amount * 0.10
}

// VIPCustomer implementa Customer
type VIPCustomer struct{}

func (c VIPCustomer) CalculateDiscount(amount float64) float64 {
	return amount * 0.15
}
