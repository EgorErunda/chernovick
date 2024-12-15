package main

import (
	"fmt"
	"log"
)

type Order struct {
	OrderNumber  int
	CustomerName string
	OrderAmount  float64
}

type OrderLogger struct{}

func NewOrderLogger() *OrderLogger {
	return &OrderLogger{}
}

func (logger *OrderLogger) AddOrder(order Order) {
	message := fmt.Sprintf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n",
		order.OrderNumber, order.CustomerName, order.OrderAmount)

	log.Print(message)
}
