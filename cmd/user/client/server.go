package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zsmartex/pkg/v2/log"
	"github.com/zsmartex/zsmartex/internal/user/router"
	"github.com/zsmartex/zsmartex/pkg/config"
	userv1 "github.com/zsmartex/zsmartex/proto/api/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func customErrorHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Call the gRPC handler function
	resp, err := handler(ctx, req)
	if err != nil {
		log.Errorf("%+v\n", err)
	}

	// Check for an error and convert it to a custom error message
	if err != nil {
		message := "An error occurred: " + err.Error()
		return nil, status.Errorf(codes.Internal, message)
	}

	// Return the response
	return resp, nil
}

func NewGRPCServer(userServiceServer router.UserServiceServer) *grpc.Server {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(customErrorHandler))
	userv1.RegisterUserServiceServer(grpcServer, userServiceServer)

	return grpcServer
}

func NewGRPCGateway(ctx context.Context, grpcServerConfig config.GRPC, opts []grpc.DialOption) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			log.Error(err)
			message := err.Error()
			response := map[string][]string{"errors": {message}}
			body, _ := json.Marshal(response)

			// Write the HTTP response
			writer.Header().Set("Content-Type", marshaler.ContentType(response))
			status := runtime.HTTPStatusFromCode(grpc.Code(err))
			writer.WriteHeader(status)
			writer.Write(body)
		}),
	)
	err := userv1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", grpcServerConfig.Port), opts)
	if err != nil {
		return nil, err
	}

	return mux, err
}
