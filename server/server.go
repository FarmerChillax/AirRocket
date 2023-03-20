package server

import (
	"context"
	"fmt"
	"unsafe"

	"github.com/FarmerChillax/AirRocket/api/rocket_server"
	"github.com/FarmerChillax/AirRocket/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AirRocketServer struct {
	rocket_server.UnimplementedAirRocketServerServer
}

func NewAirRocketServer() *AirRocketServer {
	return &AirRocketServer{}
}

func (ars *AirRocketServer) GenerateAccessCode(ctx context.Context,
	in *rocket_server.EntryMessage) (*rocket_server.AccessCode, error) {
	defer func() {
		entryCache := NewEntryCache()
		result, _ := entryCache.All()
		fmt.Printf("entry cache size: %v value: %+v\n", unsafe.Sizeof(result), result)
	}()

	accessCode := pkg.GenerateAccessCode()
	entryCache := NewEntryCache()
	err := entryCache.Set(accessCode, in, 10)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "generate access code err: %v", err)
	}

	return &rocket_server.AccessCode{Value: accessCode}, nil
}

func (ars *AirRocketServer) GetDetail(ctx context.Context, in *rocket_server.AccessCode) (*rocket_server.EntryMessage, error) {
	accessCode := in.GetValue()
	entryMessage, err := entryCache.Get(accessCode)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "entryCache.Get err: %v", err)
	}
	return entryMessage, nil
}
