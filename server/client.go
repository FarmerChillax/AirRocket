package server

import (
	"context"

	"github.com/FarmerChillax/AirRocket/proto/rocket_server"
	"google.golang.org/grpc"
)

// import (
// 	"errors"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"

// 	"github.com/FarmerChillax/AirRocket/pkg"
// 	"github.com/FarmerChillax/AirRocket/api/rocket_client"
// )

type Client struct {
	Addr string
	// airRocketServerClient *rocket_server.AirRocketServerClient
}

func NewClient(address string) *Client {
	return &Client{Addr: address}
}

func (c *Client) GetDetail(ctx context.Context, accessCode string) (entryMessage *rocket_server.EntryMessage, err error) {
	// grpc 连接
	conn, err := grpc.DialContext(ctx, c.Addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	// 初始化客户端
	arsc := rocket_server.NewAirRocketServerClient(conn)
	entryMessage, err = arsc.GetDetail(ctx, &rocket_server.AccessCode{Value: accessCode})
	if err != nil {
		return nil, err
	}
	return entryMessage, nil
}

// type AirRocketClient struct{}

// func NewAirRocketClient() *AirRocketClient {
// 	return &AirRocketClient{}
// }

// func (arc *AirRocketClient) Transfer(in *rocket_client.TransferRequest, out rocket_client.AirRocketClient_TransferServer) error {
// 	fmt.Println(in.URI)
// 	// 判断文件是否存在
// 	fmt.Printf("CheckFileExist result: %v\n", pkg.CheckFileExist(in.URI))
// 	if exist := pkg.CheckFileExist(in.URI); !exist {
// 		return pkg.ErrFileNotExist
// 	}
// 	// 打开文件, 获取文件描述符
// 	f, err := os.Open(in.URI)
// 	if err != nil {
// 		return fmt.Errorf("打开文件出错")
// 	}
// 	defer f.Close()
// 	fileBuffer := make([]byte, 4194000)
// 	// pr, pw := io.Pipe()
// 	// loop
// 	for {
// 		// 读取文件到buffer
// 		n, err := f.Read(fileBuffer)
// 		if err != nil {
// 			if errors.Is(err, io.EOF) {
// 				log.Println("文件已加载完")
// 				return nil
// 			}
// 			return err
// 		}
// 		// 发送文件
// 		if err := out.Send(&rocket_client.TransferResponse{
// 			Data: fileBuffer[:n],
// 		}); err != nil {
// 			log.Printf("发送出错, %v", err)
// 			break
// 		}
// 	}
// 	return nil
// }
