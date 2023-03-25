package router

import (
	"context"

	authzv1 "github.com/zsmartex/zsmartex/proto/api/authz/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

type authzGRPCServer struct {
	authzv1.UnimplementedAuthzServiceServer
}

func NewAuthzGRPCServer(
	grpcServer *grpc.Server,
) authzv1.AuthzServiceServer {
	svc := authzGRPCServer{}

	authzv1.RegisterAuthzServiceServer(grpcServer, &svc)
	reflection.Register(grpcServer)

	return &svc
}

func (s *authzGRPCServer) Authorization(ctx context.Context, request *authzv1.AuthorizationRequest) (*authzv1.AuthorizationResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	cookie := md.Get("cookie")
}
