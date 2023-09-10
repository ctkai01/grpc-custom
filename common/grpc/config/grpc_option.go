package config

type GrpcOptions struct {
	Port string 
	Host string
	Name string
}

type CustomGrpcOptions struct {
	Server *GrpcOptions
	Client []*GrpcOptions
}