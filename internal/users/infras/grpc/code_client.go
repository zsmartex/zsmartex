package grpc

import (
	"github.com/google/wire"
	"github.com/zsmartex/zsmartex/cmd/users/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type codeGRPCClient struct {
	conn *grpc.ClientConn
}

var CodeGRPCClientSet = wire.NewSet(NewCodeGRPCClient)

func NewCodeGRPCClient(cfg *config.Config) {
	conn, err := grpc.Dial(cfg.CodeClient.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &codeGRPCClient{
		conn: conn,
	}, nil
}
