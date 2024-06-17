package api

import (
	"strings"

	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/application/core/domain"
	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{db: db, payment: payment}
}
func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	//Charges for current order
	paymentError := a.payment.Charge(&order)
	if paymentError != nil {
		st := status.Convert(paymentError) //convert a complex error to a status
		var allErrors []string             //slices for whole errors
		for _, detail := range st.Details() {
			//Details returns a slice of details messages attached to the status
			switch t := detail.(type) {
			case *errdetails.BadRequest:
				for _, violation := range t.GetFieldViolations() {
					allErrors = append(allErrors, violation.Description)
				}
			}
		}
		//Payment errors as fields
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field:       "payment",
			Description: strings.Join(allErrors, "\n"),
		}
		badReq := &errdetails.BadRequest{}
		badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation failed")
		statusWithDetails, _ := orderStatus.WithDetails(badReq)
		return domain.Order{}, statusWithDetails.Err()
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
