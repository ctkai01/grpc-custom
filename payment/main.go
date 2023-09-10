package main

import (
	"context"
	"log"

	"grpc_example/common/grpc"
	"grpc_example/common/grpc/config"
	discountPb "grpc_example/discount/grpc/pb"
	productPb "grpc_example/product/grpc/pb"
)

func main() {

	config := &config.CustomGrpcOptions{
		Server: nil,
		Client: []*config.GrpcOptions{
			{
				Port: ":5001",
				Host: "localhost",
				Name: "product service",
			},
			{
				Port: ":5002",
				Host: "localhost",
				Name: "discount service",
			},
		},
	}
	paymentClient, err := grpc.NewGrpcClient(config)

	if err != nil {
		log.Fatal(err)
	}

	defer paymentClient.Close("product service")
	defer paymentClient.Close("discount service")

	connProduct, err := paymentClient.GetGrpcConnection("product service")

	if err != nil {
		log.Fatal(err)
	}

	connDiscount, err := paymentClient.GetGrpcConnection("discount service")

	if err != nil {
		log.Fatal(err)
	}

	clientProduct := productPb.NewProductsClient(connProduct)
	clientDiscount := discountPb.NewDiscountClient(connDiscount)

	repProduct, err := clientProduct.GetProduct(context.Background(), &productPb.GetProductRequest{
		ProductID: 1,
	})

	if err != nil {
		log.Fatal(err)
	}

	repDiscount, err := clientDiscount.GetDiscountByProductID(context.Background(), &discountPb.DiscountRequest{
		ProductID: repProduct.Product.ProductID,
	})

	if err != nil {
		log.Fatal(err)
	}

	actualPrice := (repProduct.Product.Price * (100 - repDiscount.PercentDiscount)) / 100

	log.Printf("Payment product %s with price: %d$", repProduct.Product.Name, actualPrice)
}
