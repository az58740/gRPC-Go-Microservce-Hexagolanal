package ports

import "github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
