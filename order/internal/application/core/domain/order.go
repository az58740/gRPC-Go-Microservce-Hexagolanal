package domain

import "time"

type OrderItem struct {
	ProductCode string  `json:"product_code"` //Unique code for product
	UnitPrice   float32 `json:"unit_price"`   //Price of a single product
	Quantity    int32   `json:"quantity"`     //Count of the product
}

type Order struct {
	ID         int64       `json:"id"`
	CustomerID int64       `json:"customer_id"` //Owner of the order
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"` //List of items purchased in an order
	CreatedAt  int64       `json:"created_at"`
}

func NewOrder(customerId int64, orderItems []OrderItem) Order {
	return Order{
		CreatedAt:  time.Now().Unix(),
		Status:     "Pending",
		CustomerID: customerId,
		OrderItems: orderItems,
	}
}
