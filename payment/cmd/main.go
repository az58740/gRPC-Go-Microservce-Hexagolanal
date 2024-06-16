package main

import (
	"log"

	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/payment/config"
	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/payment/internal/adapters/db"
	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/payment/internal/adapters/grpc"
	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/payment/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()

}
