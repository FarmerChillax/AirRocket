package client

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/FarmerChillax/AirRocket/pkg"
	"github.com/FarmerChillax/AirRocket/proto/rocket_client"
)

type AirRocketService struct{}

func NewAirRocketService() *AirRocketService {
	return &AirRocketService{}
}

func (arc *AirRocketService) Transfer(in *rocket_client.TransferRequest, out rocket_client.AirRocketClient_TransferServer) error {
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
	fileBuffer := make([]byte, 4194000)
	// pr, pw := io.Pipe()
	// loop
	for {
		// 读取文件到buffer
		n, err := f.Read(fileBuffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("文件已加载完")
				return nil
			}
			return err
		}
		// 发送文件
		if err := out.Send(&rocket_client.TransferResponse{
			Data: fileBuffer[:n],
		}); err != nil {
			log.Printf("发送出错, %v", err)
			break
		}
	}
	return nil
}
