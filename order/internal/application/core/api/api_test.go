package api

import (
	"errors"
	"testing"

	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/application/core/domain"
	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/ports/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockePayment struct {
	mock.Mock
}

func (p *mockePayment) Charge(order *domain.Order) error {
	args := p.Called(order)
	return args.Error(0)
}

func Test_Should_Place_Order(t *testing.T) {
	payment := new(mocks.PaymentPort)
	db := new(mocks.DBPort)
	payment.On("Charge", mock.Anything).Return(nil) //There is no error on payment.Charge.
	db.On("Save", mock.Anything).Return(nil)        //There is no error on db.Save.
	aplication := NewApplication(db, payment)
	_, err := aplication.PlaceOrder(domain.Order{
		CustomerID: 123,
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "camera",
				UnitPrice:   12.3,
				Quantity:    3,
			},
		},
		CreatedAt: 0,
	})
	assert.Nil(t, err) //error is null in this case
}
func Test_Should_Retern_Error_When_Db_Persistence_Fail(t *testing.T) {
	payment := new(mockePayment)
	db := new(mocks.DBPort)
	payment.On("Charge", mock.Anything).Return(nil)
	db.On("Save", mock.Anything).Return(errors.New("connection error!"))
	aplication := NewApplication(db, payment)
	_, err := aplication.PlaceOrder(domain.Order{
		CustomerID: 123,
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "phone",
				UnitPrice:   14.7,
				Quantity:    1,
			},
		},
		CreatedAt: 0,
	})
	assert.EqualError(t, err, "connection error!")
}
func Test_Should_Retern_Error_When_Payment_Fail(t *testing.T) {
	payment := new(mockePayment)
	db := new(mocks.DBPort)
	payment.On("Charge", mock.Anything).Return(errors.New("insufficient balance"))
	db.On("Save", mock.Anything).Return(nil)
	aplication := NewApplication(db, payment)
	_, err := aplication.PlaceOrder(domain.Order{
		CustomerID: 123,
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "bag",
				UnitPrice:   2.5,
				Quantity:    6,
			},
		},
		CreatedAt: 0,
	})
	st, _ := status.FromError(err)
	assert.Equal(t, st.Message(), "order creation failed")
	assert.Equal(t, st.Details()[0].(*errdetails.BadRequest).FieldViolations[0].Description, "insufficient balance")
	assert.Equal(t, st.Code(), codes.InvalidArgument)
}
