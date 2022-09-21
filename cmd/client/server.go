package main

import (
	"log"
	"net"

	"github.com/FarmerChillax/AirRocket/proto/rocket_client"
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
	rocket_client.RegisterAirRocketClientServer(s, server.NewAirRocketClient())

	log.Printf("Starting AirRocket client service on port: %v\n", ADDRESS)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to start AirRocket client serve: %v", err)
	}
}

/*
题目：实现简易版load balancer
用户当前搜索的request
->发送到分流器
->分流器发送到一台机器的load balancer模块
->load balancer接到请求，实时决定是否交给当前机器的业务模块处理
-> Yes, 则将当前request发送给后续模块；No，则将request返回给分流器

判断Yes或No的标准：使得*任意时刻*的业务模块的QPS都不能大于N

request
|
分流器
|                               |
服务器1 Load balancer模块             服务器2 Load balancer模块
当前时刻QPS>N?
| Y             | N
request传给后续模块   拒绝请求, request返回分流器
*/
