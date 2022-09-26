package client

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/FarmerChillax/AirRocket/proto/rocket_client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AirRocketClient struct {
	Addr string
}

func NewAirRocketClient(address string) *AirRocketClient {

	return &AirRocketClient{Addr: address}
}

// 从目标(src)获取数据,存到本机目录(dst)中
// src: source; dst: destination.
func (arc *AirRocketClient) Download(ctx context.Context, src, dst string) error {
	// 打开 grpc
	conn, err := grpc.DialContext(ctx, arc.Addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	arcc := rocket_client.NewAirRocketClientClient(conn)
	airRocketClient, err := arcc.Transfer(ctx, &rocket_client.TransferRequest{URI: src})
	if err != nil {
		return err
	}

	// 打开文件描述符
	f, err := os.OpenFile(dst, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	for {
		tr, err := airRocketClient.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("传输完成")
				break
			}
			return err
		}
		n, err := f.Write(tr.Data)
		if err != nil {
			return err
		}
		log.Println("写入文件大小:", n)
	}

	return nil
}

func GetFile(ctx context.Context, address, URI string) error {
	dialOpts := grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(4194000))
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), dialOpts)
	if err != nil {
		return err
	}
	defer conn.Close()
	// grpc.NewClientStream()
	arcc := rocket_client.NewAirRocketClientClient(conn)
	arc, err := arcc.Transfer(ctx, &rocket_client.TransferRequest{
		URI: URI,
	})
	if err != nil {
		return err
	}
	for {

		tr, err := arc.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("传输完成")
				break
			}
			return err
		}
		f, err := os.OpenFile(`F:\Everything.mkv`, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}
		// /home/farmer/go1.19.1.linux-amd64.tar.gz
		n, err := f.Write(tr.Data)
		if err != nil {
			return err
		}
		log.Println("写入文件大小:", n)

	}
	return nil
}
