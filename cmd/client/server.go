package main

import "github.com/FarmerChillax/AirRocket/startup"

func main() {
	// 开启边缘节点服务端
	_ = startup.AirRocketClientServerStartUp()
}
