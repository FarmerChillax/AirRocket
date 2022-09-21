package main

import (
	"log"
	"net"

	"github.com/FarmerChillax/AirRocket/proto/rocket_server"
	"github.com/FarmerChillax/AirRocket/server"
	"google.golang.org/grpc"
)

const ADDRESS = "127.0.0.1:5000"

func main() {

	lis, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	rocket_server.RegisterAirRocketServerServer(s, server.NewAirRocketServer())

	log.Printf("Starting gRPC listener on port: %v", ADDRESS)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
