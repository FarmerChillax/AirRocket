package startup

import (
	"github.com/FarmerChillax/AirRocket/api/rocket_client"
	"github.com/FarmerChillax/AirRocket/api/rocket_server"
	"github.com/FarmerChillax/AirRocket/client"
	"github.com/FarmerChillax/AirRocket/server"
	"google.golang.org/grpc"
)

// RegisterCenterCServer 此处注册中心节点的 grpc Server
func RegisterCenterServer(grpcServer *grpc.Server) error {
	rocket_server.RegisterAirRocketServerServer(grpcServer, server.NewAirRocketServer())
	return nil
}

// RegisterNodeServer 此处注册客户端节点的 grpc Server
func RegisterNodeServer(grpcServer *grpc.Server) error {
	rocket_client.RegisterAirRocketNodeServerServer(grpcServer, client.NewAirRocketService())
	return nil
}
