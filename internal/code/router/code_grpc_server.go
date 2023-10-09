package router

import (
	"context"

	servicesv1 "github.com/zsmartex/zsmartex/proto/services/v1"
)

var _ servicesv1.CodeServiceServer = (*CodeServiceServer)(nil)

type CodeServiceServer struct {
	servicesv1.UnimplementedCodeServiceServer
}

func NewUserServiceServer() servicesv1.CodeServiceServer {
	return &CodeServiceServer{}
}

func (s *CodeServiceServer) GenerateCode(context.Context, *servicesv1.GenerateCodeRequest) (*servicesv1.GenerateCodeResponse, error) {
	return &servicesv1.GenerateCodeResponse{}, nil
}

func (s *CodeServiceServer) CheckCodes(context.Context, *servicesv1.CheckCodesRequest) (*servicesv1.CheckCodesResponse, error) {
	return &servicesv1.CheckCodesResponse{}, nil
}

func (s *CodeServiceServer) GetPendingCode(context.Context, *servicesv1.GetPendingCodeRequest) (*servicesv1.GetPendingCodeResponse, error) {
	return &servicesv1.GetPendingCodeResponse{}, nil
}

func (s *CodeServiceServer) ValidateCodes(context.Context, *servicesv1.ValidateCodesRequest) (*servicesv1.ValidateCodesResponse, error) {
	return &servicesv1.ValidateCodesResponse{}, nil
}
