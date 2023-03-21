package server

import (
	"context"
	"fmt"

	"github.com/FarmerChillax/AirRocket/api/rocket_server"
	"github.com/FarmerChillax/AirRocket/pkg"
	"github.com/FarmerChillax/AirRocket/vars"
)

type AirRocketServer struct {
	rocket_server.UnimplementedAirRocketServerServer
}

func NewAirRocketServer() *AirRocketServer {
	return &AirRocketServer{}
}

func (ars *AirRocketServer) GenerateAccessCode(ctx context.Context, in *rocket_server.EntryMessage) (*rocket_server.AccessCode, error) {
	// 生成提取码
	var accessCode string
	for i := 0; i < vars.AppSettings.GenerateCount; i++ {
		accessCode = pkg.GenerateAccessCode()
		if _, exist := vars.Cache.Get(accessCode); !exist {
			break
		}
	}
	// vars.Cache.Set(accessCode, in, cache.DefaultExpiration)
	vars.Cache.SetDefault(accessCode, in)
	data, exist := vars.Cache.Get(accessCode)
	fmt.Println(accessCode, data, exist)
	return &rocket_server.AccessCode{Value: accessCode}, nil
}

func (ars *AirRocketServer) GetDetail(ctx context.Context, in *rocket_server.AccessCode) (*rocket_server.EntryMessage, error) {
	accessCode := in.GetValue()
	entryDetail, exist := vars.Cache.Get(accessCode)
	if !exist {
		return nil, ErrNotFound
	}

	return entryDetail.(*rocket_server.EntryMessage), nil
}
