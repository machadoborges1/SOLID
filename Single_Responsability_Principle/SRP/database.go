//Este arquivo é responsável por salvar os pedidos no banco de dados (simulado).

package main

import "fmt"

func SaveOrder(order Order) {
	fmt.Printf("Order %d saved to database.\n", order.ID)
}
