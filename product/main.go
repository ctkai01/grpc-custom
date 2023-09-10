package main

import (
	"fmt"
	"grpc_example/common/grpc"
	"grpc_example/common/grpc/config"
	"grpc_example/product/grpc/pb"
	productServer "grpc_example/product/server"
	"log"

	googleGrpc "google.golang.org/grpc"
)

func main() {
	config := config.CustomGrpcOptions{
		Server: &config.GrpcOptions{
			Port: ":5001",
			Host: "localhost",
			Name: "product service",
		},
		Client: nil,
	}

	grpcServer := grpc.NewGrpcServer(&config)
	productServer := productServer.NewProductServer()

	grpcServer.GrpcServiceBuilder().RegisterRoutes(
		func(s *googleGrpc.Server) {
			pb.RegisterProductsServer(s, productServer)
		},
	)

	fmt.Println("Product service is running")

	if err := grpcServer.RunGrpcServer(); err != nil {
		log.Fatal(err)
	}
}