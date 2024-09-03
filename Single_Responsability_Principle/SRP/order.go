package main

type Order struct {
	ID     int
	Amount float64
}

func ProcessOrder(order Order) error {
	if err := ValidateOrder(order); err != nil {
		return err
	}

	SaveOrder(order)
	SendNotification(order)

	return nil
}
