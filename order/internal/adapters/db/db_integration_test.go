package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/application/core/domain"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type OrderDatabaseTestSuite struct {
	suite.Suite
	DataSourceUrl string
	ctx           context.Context
}

func (o *OrderDatabaseTestSuite) SetupSuite() {
	ctx := context.Background()
	port := "3306/tcp"
	dbURL := func(host string, port nat.Port) string {
		return fmt.Sprintf("root:Az@1358740@tcp(localhost:%s)/grpcorder?charset=utf8mb4&parseTime=True&loc=Local", port.Port())
	}
	req := testcontainers.ContainerRequest{
		Image:        "docker.io/mysql:8.0.36",
		ExposedPorts: []string{port},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "Az@1358740",
			"MYSQL_DATABASE":      "grpcorder",
		},
		WaitingFor: wait.ForSQL(nat.Port(port), "mysql", dbURL),
	}
	mysqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal("Failed to start Mysql.", err)
	}
	endpoint, _ := mysqlContainer.Endpoint(ctx, "")
	o.DataSourceUrl = fmt.Sprintf("root:Az@1358740@tcp(%s)/grpcorder?charset=utf8mb4&parseTime=True&loc=Local", endpoint)
}

func (o *OrderDatabaseTestSuite) Test_Should_Save_Order() {
	t := o.T()
	adapter, err := NewAdapter(o.DataSourceUrl)
	o.Nil(t, err)
	saveErr := adapter.Save(&domain.Order{
		CustomerID: 123,
		Status:     "pending",
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "CAM",
				Quantity:    5,
				UnitPrice:   1.32,
			},
		},
	})
	o.Nil(t, saveErr)
}
func (o *OrderDatabaseTestSuite) Test_Should_Get_Order() {
	adapter, _ := NewAdapter(o.DataSourceUrl)
	order := domain.NewOrder(2, []domain.OrderItem{
		{
			ProductCode: "CAM",
			Quantity:    5,
			UnitPrice:   1.32,
		},
	})
	adapter.Save(&order)
	id := strconv.FormatInt(order.ID, 10)
	ord, _ := adapter.Get(id)
	o.Equal(int64(2), ord.CustomerID)
}

func TestOrderDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(OrderDatabaseTestSuite))
}
