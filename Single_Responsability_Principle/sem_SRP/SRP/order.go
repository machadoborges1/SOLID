package main

import (
	"errors"
	"fmt"
)

type Order struct {
	ID     int
	Amount float64
}

func ProcessOrder(order Order) error {
	// Validação
	if order.Amount <= 0 {
		return errors.New("invalid order amount")
	}

	// Salvar no banco de dados (simulação)
	fmt.Printf("Order %d saved to database.\n", order.ID)

	// Enviar notificação
	fmt.Printf("Notification sent for order %d.\n", order.ID)

	return nil
}

func main() {
	order := Order{ID: 1, Amount: 100.0}
	err := ProcessOrder(order)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

//Neste exemplo, a função ProcessOrder faz várias coisas diferentes: valida o pedido, salva no banco de dados e envia uma notificação.
//Todas essas responsabilidades estão juntas, violando o SRP.
