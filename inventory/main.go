package main

import (
	"log"
	"net"

	"github.com/harikesh-yadav/go-grpc-app/inventory/service"
	pb "github.com/harikesh-yadav/go-grpc-app/proto"
	"google.golang.org/grpc"
)

var addr = "0.0.0.0:8001"

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opt []grpc.ServerOption
	grpcSever := grpc.NewServer(opt...)

	pb.RegisterInventoryServiceServer(grpcSever, &service.InventoryServer{})
	log.Println("Server listning at port ", lis.Addr().String())

	if err := grpcSever.Serve(lis); err != nil {
		log.Printf("Failed to serve at %v", err)
	}
}
