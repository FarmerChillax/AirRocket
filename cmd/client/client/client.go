package main

import (
	"context"
	"log"

	"github.com/FarmerChillax/AirRocket/client"
)

const ADDRESS = "127.0.0.1:5000"

func main() {
	// ctx := context.WithTimeout(context.Background(), time.Minute*30)
	ctx := context.Background()
	err := client.GetFile(ctx, ADDRESS, `F:\Downloads\Everything.Everywhere.All.At.Once.2022.1080p.WEB-DL.DDP5.1.H.264-EVO.mkv`)
	if err != nil {
		log.Fatalf("client.GetFile err: %v", err)
	}
}
