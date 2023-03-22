package main

import (
	"log"

	"github.com/FarmerChillax/AirRocket/startup"
)

func main() {

	// 开启中心节点服务端
	err := startup.AirRocketCenterServerStartUp()
	if err != nil {
		log.Printf("startup.StartUp err: %v", err)
	}
}
