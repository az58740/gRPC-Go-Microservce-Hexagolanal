package e2e

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/az58740/grpc-microservices-proto/golang/order"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CreateOrderTestSuite struct {
	suite.Suite
	compose *tc.LocalDockerCompose
}

func (c *CreateOrderTestSuite) Setupsuit() {
	composeFilePaths := []string{"resources/docker-compose.yml"}
	identifier := strings.ToLower(uuid.New().String())

	compose := tc.NewLocalDockerCompose(composeFilePaths, identifier)
	c.compose = compose

	exeError := compose.WithCommand([]string{"up", "-d"}).Invoke()
	err := exeError.Error
	if err != nil {
		log.Fatalf("Could not run compose stack: %v", err)
	}

}
func (c *CreateOrderTestSuite) Test_Should_Create_Order() {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("Failed to connect order service. Error:%v", err)
	}
	defer conn.Close()
	orderClient := order.NewOrderClient(conn)
	createOrderResponse, errCreate := orderClient.Create(context.Background(), &order.CreateOrderRequest{
		UserId: 123,
		OrderItems: []*order.OrderItem{
			{
				ProductCode: "CAM123",
				Quantity:    3,
				UnitPrice:   1.236,
			},
		},
	})
	//Verifies there is no error
	c.Nil(errCreate)
	getOrderResponse, errGet := orderClient.Get(context.Background(), &order.GetOrderRequest{
		OrderId: createOrderResponse.OrderId,
	})
	c.Nil(errGet)
	c.Equal(int64(123), getOrderResponse.UserId)
	orderItem := getOrderResponse.OrderItems[0]
	c.Equal(float32(1.236), orderItem.UnitPrice)
	c.Equal(int32(3), orderItem.Quantity)
	c.Equal("CAM123", orderItem.ProductCode)
}
func (c *CreateOrderTestSuite) TearDownSuite() {
	execError := c.compose.WithCommand([]string{"down"}).Invoke()
	err := execError.Error
	if err != nil {
		log.Fatalf("Could not shutdown compose stack: %v", err)
	}
}
func TestCreateOrderTestSuite(t *testing.T) {
	suite.Run(t, new(CreateOrderTestSuite))
}
