package grpc

import (
	"context"

	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/application/core/domain"
	pb "github.com/az58740/grpc-microservices-proto/golang/order"
)

func (a Adapter) Create(ctx context.Context, request *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, item := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{OrderId: result.ID}, nil

}
func (a Adapter) Get(ctx context.Context, request *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	orderID := string(rune(request.GetOrderId()))
	result, err := a.api.GetOrder(orderID)
	if err != nil {
		return nil, err
	}
	var orderItems []*pb.OrderItem
	for _, item := range result.OrderItems {
		orderItems = append(orderItems, &pb.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	return &pb.GetOrderResponse{
		UserId:     result.CustomerID,
		OrderItems: orderItems,
	}, nil
}
