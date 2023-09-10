package server

import (
	"context"
	"fmt"
	"grpc_example/discount/grpc/pb"
)

type discountDBFake struct {
	discount int
	productID int
}

var (
	dataDiscountFake = []discountDBFake{
		{
			discount: 10,
			productID: 1,
		},
		{
			discount: 50,
			productID: 2,
		},
		{
			discount: 70,
			productID: 3,
		},
	}
)

type discountService struct {
	pb.DiscountServer
}

func NewDiscountServer() *discountService {
	return &discountService {}
}

func (ds *discountService) GetDiscountByProductID(c context.Context, discountReq *pb.DiscountRequest) (*pb.DiscountResponse, error) {
	
	
	var discountProduct *discountDBFake

	for _, discount := range dataDiscountFake {
		if discount.productID == int(discountReq.ProductID) {
			discountProduct = &discount
			break
		}
	}

	if discountProduct == nil {
		return nil, fmt.Errorf("error find discount")
	}

	return &pb.DiscountResponse{
		PercentDiscount: int32(discountProduct.discount),
	}, nil
}