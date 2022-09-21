package server

import (
	"fmt"

	"github.com/FarmerChillax/AirRocket/proto/rocket_client"
)

type AirRocketClient struct{}

func NewAirRocketClient() *AirRocketClient {
	return &AirRocketClient{}
}

func (arc *AirRocketClient) Transfer(in *rocket_client.TransferRequest, out rocket_client.AirRocketClient_TransferServer) error {
	// panic("not implemented") // TODO: Implement
	fmt.Println(in.URI)
	// 判断文件是否存在
	// loop
	// 读取文件到buffer
	// 发送文件
	// out.Send(&rocket_client.TransferResponse{})
	// 清空buffer
	return nil
}

// func (arc *AirRocketClient) Transfer(ctx context.Context, in *rocket_client.TransferRequest, opts ...grpc.CallOption) (rocket_client.AirRocketClient_TransferClient, error) {
// 	panic("not implemented") // TODO: Implement
// 	// 判断文件是否存在

// 	// loop
// 	// 读取文件到buffer
// 	// 发送文件
// 	// rocket_client.AirRocketClient_TransferClient
// 	// 清空buffer

// }
