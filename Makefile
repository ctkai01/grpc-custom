gen-grpc-discount:
	protoc -Idiscount/grpc/pb --go_out=. --go_opt=module=grpc_example --go-grpc_out=. --go-grpc_opt=module=grpc_example discount/grpc/pb/discount.proto 
gen-grpc-product:
	protoc -Iproduct/grpc/pb --go_out=. --go_opt=module=grpc_example --go-grpc_out=. --go-grpc_opt=module=grpc_example product/grpc/pb/product.proto 