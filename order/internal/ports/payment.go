package ports

import "github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
