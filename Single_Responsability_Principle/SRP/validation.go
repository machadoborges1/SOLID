//Este arquivo é responsável pela validação dos pedidos.

package main

import "errors"

func ValidateOrder(order Order) error {
	if order.Amount <= 0 {
		return errors.New("invalid order amount")
	}
	return nil
}
