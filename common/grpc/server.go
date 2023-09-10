package grpc

import (
	"fmt"
	"grpc_example/common/grpc/config"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer interface {
	RunGrpcServer(configGrpc ...func(grpcServer *grpc.Server)) error
	GracefulShutdown()
	GetCurrentGrpcServer() *grpc.Server
	GrpcServiceBuilder() *GrpcServiceBuilder
}

type grpcServer struct {
	server         *grpc.Server
	config         *config.CustomGrpcOptions
	serviceName    string
	serviceBuilder *GrpcServiceBuilder
}

func NewGrpcServer(
	config *config.CustomGrpcOptions,
) GrpcServer {
	s := grpc.NewServer()

	return &grpcServer{
		server:         s,
		config:         config,
		serviceName:    config.Server.Name,
		serviceBuilder: NewGrpcServiceBuilder(s),
	}
}

func (s *grpcServer) RunGrpcServer(configGrpc ...func(grpcServer *grpc.Server)) error {
	l, err := net.Listen("tcp",s.config.Server.Port)

	if err != nil {
		return fmt.Errorf("Error net.Listen: ", err)
	}

	if len(configGrpc) > 0 {
		grpcFunc := configGrpc[0]
		if grpcFunc != nil {
			grpcFunc(s.server)
		}
	}

	err = s.server.Serve(l)

	if err != nil {
		return fmt.Errorf("grpc server error: ", err)
	}

	return nil
}

func (s *grpcServer) GrpcServiceBuilder() *GrpcServiceBuilder {
	return s.serviceBuilder
}

func (s *grpcServer) GetCurrentGrpcServer() *grpc.Server {
	return s.server
}

func (s *grpcServer) GracefulShutdown() {
	s.server.Stop()
	s.server.GracefulStop()
}