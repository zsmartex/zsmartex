// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/code/v1/code.proto

package codev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CodeServiceClient is the client API for CodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeServiceClient interface {
	GetPendingCode(ctx context.Context, in *GetPendingCodeRequest, opts ...grpc.CallOption) (*GetPendingCodeResponse, error)
	GenerateCode(ctx context.Context, in *GenerateCodeRequest, opts ...grpc.CallOption) (*GenerateCodeResponse, error)
	CheckCodes(ctx context.Context, in *CheckCodesRequest, opts ...grpc.CallOption) (*CheckCodesResponse, error)
	ValidateCodes(ctx context.Context, in *ValidateCodesRequest, opts ...grpc.CallOption) (*ValidateCodesResponse, error)
}

type codeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeServiceClient(cc grpc.ClientConnInterface) CodeServiceClient {
	return &codeServiceClient{cc}
}

func (c *codeServiceClient) GetPendingCode(ctx context.Context, in *GetPendingCodeRequest, opts ...grpc.CallOption) (*GetPendingCodeResponse, error) {
	out := new(GetPendingCodeResponse)
	err := c.cc.Invoke(ctx, "/api.code.v1.CodeService/GetPendingCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServiceClient) GenerateCode(ctx context.Context, in *GenerateCodeRequest, opts ...grpc.CallOption) (*GenerateCodeResponse, error) {
	out := new(GenerateCodeResponse)
	err := c.cc.Invoke(ctx, "/api.code.v1.CodeService/GenerateCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServiceClient) CheckCodes(ctx context.Context, in *CheckCodesRequest, opts ...grpc.CallOption) (*CheckCodesResponse, error) {
	out := new(CheckCodesResponse)
	err := c.cc.Invoke(ctx, "/api.code.v1.CodeService/CheckCodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServiceClient) ValidateCodes(ctx context.Context, in *ValidateCodesRequest, opts ...grpc.CallOption) (*ValidateCodesResponse, error) {
	out := new(ValidateCodesResponse)
	err := c.cc.Invoke(ctx, "/api.code.v1.CodeService/ValidateCodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeServiceServer is the server API for CodeService service.
// All implementations should embed UnimplementedCodeServiceServer
// for forward compatibility
type CodeServiceServer interface {
	GetPendingCode(context.Context, *GetPendingCodeRequest) (*GetPendingCodeResponse, error)
	GenerateCode(context.Context, *GenerateCodeRequest) (*GenerateCodeResponse, error)
	CheckCodes(context.Context, *CheckCodesRequest) (*CheckCodesResponse, error)
	ValidateCodes(context.Context, *ValidateCodesRequest) (*ValidateCodesResponse, error)
}

// UnimplementedCodeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCodeServiceServer struct {
}

func (UnimplementedCodeServiceServer) GetPendingCode(context.Context, *GetPendingCodeRequest) (*GetPendingCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPendingCode not implemented")
}
func (UnimplementedCodeServiceServer) GenerateCode(context.Context, *GenerateCodeRequest) (*GenerateCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateCode not implemented")
}
func (UnimplementedCodeServiceServer) CheckCodes(context.Context, *CheckCodesRequest) (*CheckCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCodes not implemented")
}
func (UnimplementedCodeServiceServer) ValidateCodes(context.Context, *ValidateCodesRequest) (*ValidateCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCodes not implemented")
}

// UnsafeCodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeServiceServer will
// result in compilation errors.
type UnsafeCodeServiceServer interface {
	mustEmbedUnimplementedCodeServiceServer()
}

func RegisterCodeServiceServer(s grpc.ServiceRegistrar, srv CodeServiceServer) {
	s.RegisterService(&CodeService_ServiceDesc, srv)
}

func _CodeService_GetPendingCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPendingCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).GetPendingCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.code.v1.CodeService/GetPendingCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).GetPendingCode(ctx, req.(*GetPendingCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeService_GenerateCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).GenerateCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.code.v1.CodeService/GenerateCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).GenerateCode(ctx, req.(*GenerateCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeService_CheckCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).CheckCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.code.v1.CodeService/CheckCodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).CheckCodes(ctx, req.(*CheckCodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeService_ValidateCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateCodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).ValidateCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.code.v1.CodeService/ValidateCodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).ValidateCodes(ctx, req.(*ValidateCodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeService_ServiceDesc is the grpc.ServiceDesc for CodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.code.v1.CodeService",
	HandlerType: (*CodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPendingCode",
			Handler:    _CodeService_GetPendingCode_Handler,
		},
		{
			MethodName: "GenerateCode",
			Handler:    _CodeService_GenerateCode_Handler,
		},
		{
			MethodName: "CheckCodes",
			Handler:    _CodeService_CheckCodes_Handler,
		},
		{
			MethodName: "ValidateCodes",
			Handler:    _CodeService_ValidateCodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/code/v1/code.proto",
}