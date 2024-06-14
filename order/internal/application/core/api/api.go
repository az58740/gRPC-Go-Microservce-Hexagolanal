package api

import (
	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/application/core/domain"
	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}
func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
func (a Application) GetOrder(id string) (domain.Order, error) {
	order, err := a.db.Get(id)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
