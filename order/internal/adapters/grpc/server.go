package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/ports"
	pb "github.com/az58740/grpc-microservices-proto/golang/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	pb.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}
func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port: %v ", err)
	}
	var opt []grpc.ServerOption
	grpcServer := grpc.NewServer(opt...)
	pb.RegisterOrderServer(grpcServer, a.UnimplementedOrderServer)
	if "config" == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d  err:%v", a.port, err)
	}
}
