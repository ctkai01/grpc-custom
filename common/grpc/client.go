package grpc

import (
	"fmt"
	"grpc_example/common/grpc/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	// conn *grpc.ClientConn
	conns map[string]*grpc.ClientConn
}

type GrpcClient interface {
	GetGrpcConnection(name string) (*grpc.ClientConn, error)
	Close(name string) error
}

func NewGrpcClient(config *config.CustomGrpcOptions) (GrpcClient, error) {
	// Grpc Client to call Grpc Server
	// https://sahansera.dev/building-grpc-client-go/

	if len(config.Client) == 0 {
		return nil, fmt.Errorf("client empty")
	}
	
	clientConns := make(map[string]*grpc.ClientConn)
	for _, optionClient := range config.Client {
		conn, err := grpc.Dial(fmt.Sprintf("%s%s", optionClient.Host, optionClient.Port),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, err
		}
		clientConns[optionClient.Name] = conn
	}

	return &grpcClient{conns: clientConns}, nil
}

func (g *grpcClient) GetGrpcConnection(name string) (*grpc.ClientConn, error) {
	conn, isExist := g.conns[name]

	if !isExist {
		return nil, fmt.Errorf("error get connect not exist")
	}
	return conn, nil
}

func (g *grpcClient) Close(name string) error {
	conn, isExist := g.conns[name]

	if !isExist {
		return fmt.Errorf("error get connect not exist")
	}
	return conn.Close()
}
