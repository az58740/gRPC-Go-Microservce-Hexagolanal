package ports

import "github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
	GetOrder(id string) (domain.Order, error)
}
