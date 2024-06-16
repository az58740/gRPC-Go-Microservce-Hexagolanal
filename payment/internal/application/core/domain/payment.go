package domain

import (
	"time"
)

type Payment struct {
	ID         int64   `json:"id"`
	CustomerID int64   `json:"customer_id"`
	Status     string  `json:"status"`
	OrderID    int64   `json:"order_id"`
	TotalPrice float32 `json:"total_price"`
	CreateAt   int64   `json:"create_at"`
}

func NewPayment(cusomerId int64, orderId int64, totalPrice float32) Payment {
	return Payment{
		CustomerID: cusomerId,
		Status:     "pending",
		OrderID:    orderId,
		TotalPrice: totalPrice,
		CreateAt:   time.Now().Unix(),
	}
}
