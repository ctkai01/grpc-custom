package main

import (
	"fmt"
	"grpc_example/common/grpc"
	"grpc_example/common/grpc/config"
	"grpc_example/discount/grpc/pb"
	discountServer "grpc_example/discount/server"
	"log"

	googleGrpc "google.golang.org/grpc"
)

func main() {
	config := config.CustomGrpcOptions{
		Server: &config.GrpcOptions{
			Port: ":5002",
			Host: "localhost",
			Name: "discount service",
		},
		Client: nil,
	}

	grpcServer := grpc.NewGrpcServer(&config)
	discountServer := discountServer.NewDiscountServer()

	grpcServer.GrpcServiceBuilder().RegisterRoutes(
		func(s *googleGrpc.Server) {
			pb.RegisterDiscountServer(s, discountServer)
		},
	)

	fmt.Println("Discount service is running")

	if err := grpcServer.RunGrpcServer(); err != nil {
		log.Fatal(err)
	}
}
