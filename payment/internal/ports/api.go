package ports

import (
	"context"

	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/payment/internal/application/core/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
