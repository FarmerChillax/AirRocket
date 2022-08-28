package server

import (
	"context"
	"fmt"
	"unsafe"

	"github.com/FarmerChillax/AirRocket/pkg"
	"github.com/FarmerChillax/AirRocket/proto/rocket_server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AirRocketServer struct{}

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
	err := entryCache.Set(accessCode, in)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "generate access code err: %v", err)
	}

	return &rocket_server.AccessCode{Value: accessCode}, nil
}

func (ars *AirRocketServer) GetDetail(ctx context.Context, in *rocket_server.AccessCode) (*rocket_server.EntryMessage, error) {
	panic("not implemented") // TODO: Implement
}
