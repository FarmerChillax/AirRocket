package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/FarmerChillax/AirRocket/vars"
	"google.golang.org/grpc"
)

func AirRocketClientServerStartUp() error {
	// 加载配置项
	err := LoadNodeServerConfig()
	if err != nil {
		panic(fmt.Errorf("load config err: %v", err))
	}

	// 初始化公共变量
	err = SetupNodeServerVars()
	if err != nil {
		panic(fmt.Errorf("setup vars err: %v", err))
	}

	// 初始化服务端
	lis, err := net.Listen("tcp", vars.ServerSettings.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	grpcServer := grpc.NewServer()

	if err := RegisterCenterServer(grpcServer); err != nil {
		log.Printf("StartUp.RegisterGRPCServer err: %v", err)
		return err
	}

	log.Printf("Starting AirRocket client listener on: %v", vars.ServerSettings.Address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}
