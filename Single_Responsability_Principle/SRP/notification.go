package main

import "fmt"

func SendNotification(order Order) {
	fmt.Printf("Notification sent for order %d.\n", order.ID)
}
