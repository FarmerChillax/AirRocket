package startup

import (
	"github.com/FarmerChillax/AirRocket/api/rocket_server"
	"github.com/FarmerChillax/AirRocket/server"
	"google.golang.org/grpc"
)

// RegisterGRPCServer 此处注册pb的Server
func RegisterGRPCServer(grpcServer *grpc.Server) error {
	rocket_server.RegisterAirRocketServerServer(grpcServer, server.NewAirRocketServer())
	return nil
}
