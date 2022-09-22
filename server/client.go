package server

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/FarmerChillax/AirRocket/pkg"
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
	fmt.Printf("CheckFileExist result: %v\n", pkg.CheckFileExist(in.URI))
	if exist := pkg.CheckFileExist(in.URI); !exist {
		return pkg.ErrFileNotExist
	}
	// 打开文件, 获取文件描述符
	f, err := os.Open(in.URI)
	if err != nil {
		return fmt.Errorf("打开文件出错")
	}
	defer f.Close()
	fileBuffer := make([]byte, 2048)
	// loop
	for {
		// 读取文件到buffer
		_, err := f.Read(fileBuffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("文件已加载完")
				return nil
			}
			return fmt.Errorf("读取出错")
		}
		// 发送文件
		if err := out.Send(&rocket_client.TransferResponse{
			Data: fileBuffer,
		}); err != nil {
			log.Fatal("发送出错")
		}
		// 清空buffer
		fileBuffer = fileBuffer[:0]
	}
}
