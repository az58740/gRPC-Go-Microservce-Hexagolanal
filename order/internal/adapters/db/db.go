package db

import (
	"fmt"

	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}
type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint //back refrence to order model
}
type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	//Auto migration for givin models
	//Be sure the tables are created correctly
	err := db.AutoMigrate(&Order{}, OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", openErr)
	}
	return &Adapter{db: db}, nil
}
func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order
	//find order by ID and put it into order entity
	res := a.db.First(&orderEntity, id)
	var orderItems []domain.OrderItem
	for _, item := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}
	return order, res.Error
}
func (a Adapter) Save(order *domain.Order) error {
	var orderItems []OrderItem
	for _, item := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	orderModel := Order{
		CustomerID: order.ID,
		Status:     order.Status,
		OrderItems: orderItems,
	}
	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = int64(order.ID)
	}
	return res.Error
}
