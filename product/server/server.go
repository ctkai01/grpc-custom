package server

import (
	"context"
	"fmt"
	"grpc_example/product/grpc/pb"
)

var (
	dataProductFake = []pb.Product{
		{
			ProductID: 1,
			Name: "Shoe",
			Price: 1000,
		},
		{
			ProductID: 2,
			Name: "Hat",
			Price: 1500,
		},
		{
			ProductID: 3,
			Name: "Jean",
			Price: 500,
		},
	}
)

type productServer struct {
	pb.ProductsServer
}

func NewProductServer() *productServer {
	return &productServer {}
}

func (ds *productServer) GetProduct(c context.Context, productReq *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	var productById *pb.Product

	for _, product := range dataProductFake {
		if product.ProductID == productReq.ProductID {
			productById = &product
			break
		}
	}

	if productById == nil {
		return nil, fmt.Errorf("error find product")
	}

	return &pb.GetProductResponse{
		Product: productById,
	}, nil
}