package router

import (
	"google.golang.org/grpc"

	userv1 "github.com/zsmartex/zsmartex/proto/api/user/v1"
	servicesv1 "github.com/zsmartex/zsmartex/proto/services/v1"
)

func registerRouterHooks(grpcServer *grpc.Server, router UserServiceServer) {
	userv1.RegisterUserServiceServer(grpcServer, router)
	servicesv1.RegisterUserServiceServer(grpcServer, router)
}
