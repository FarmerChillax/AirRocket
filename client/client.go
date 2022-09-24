package client

import (
	"context"
	"log"
	"os"

	"github.com/FarmerChillax/AirRocket/proto/rocket_client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
			return err
		}
		f, err := os.OpenFile(`F:\Everything.mkv`, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}
		n, err := f.Write(tr.Data)
		if err != nil {
			return err
		}
		log.Println("写入文件大小:", n)

	}

}
