package main

import (
	"github.com/FarmerChillax/AirRocket/startup"
)

const ADDRESS = "127.0.0.1:5000"

func main() {
	// 开启中心节点服务端
	// err := startup.AirRocketCenterServerStartUp()
	// if err != nil {
	// 	log.Printf("startup.StartUp err: %v", err)
	// }
	// 开启边缘节点服务端
	_ = startup.AirRocketClientServerStartUp()
}
