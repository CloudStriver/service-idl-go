// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: usercenter.proto

package pb

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

const (
	Usercenter_FindUserByUserIdList_FullMethodName     = "/pb.usercenter/FindUserByUserIdList"
	Usercenter_SendEmailCode_FullMethodName            = "/pb.usercenter/SendEmailCode"
	Usercenter_AddUserRemainder_FullMethodName         = "/pb.usercenter/AddUserRemainder"
	Usercenter_UpdateUserInfo_FullMethodName           = "/pb.usercenter/UpdateUserInfo"
	Usercenter_AddUserRemainderRollBack_FullMethodName = "/pb.usercenter/AddUserRemainderRollBack"
	Usercenter_FindUserInfoByUserIdList_FullMethodName = "/pb.usercenter/FindUserInfoByUserIdList"
)

// UsercenterClient is the client API for Usercenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsercenterClient interface {
	FindUserByUserIdList(ctx context.Context, in *FindUserByUserIdListReq, opts ...grpc.CallOption) (*FindUserByUserIdListResp, error)
	SendEmailCode(ctx context.Context, in *SendEmailCodeReq, opts ...grpc.CallOption) (*SendEmailCodeResp, error)
	AddUserRemainder(ctx context.Context, in *AddUserRemainderReq, opts ...grpc.CallOption) (*AddUserRemainderResp, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoResp, error)
	AddUserRemainderRollBack(ctx context.Context, in *AddUserRemainderReq, opts ...grpc.CallOption) (*AddUserRemainderResp, error)
	FindUserInfoByUserIdList(ctx context.Context, in *FindUserInfoByUserIdListReq, opts ...grpc.CallOption) (*FindUserInfoByUserIdListResp, error)
}

type usercenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUsercenterClient(cc grpc.ClientConnInterface) UsercenterClient {
	return &usercenterClient{cc}
}

func (c *usercenterClient) FindUserByUserIdList(ctx context.Context, in *FindUserByUserIdListReq, opts ...grpc.CallOption) (*FindUserByUserIdListResp, error) {
	out := new(FindUserByUserIdListResp)
	err := c.cc.Invoke(ctx, Usercenter_FindUserByUserIdList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) SendEmailCode(ctx context.Context, in *SendEmailCodeReq, opts ...grpc.CallOption) (*SendEmailCodeResp, error) {
	out := new(SendEmailCodeResp)
	err := c.cc.Invoke(ctx, Usercenter_SendEmailCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) AddUserRemainder(ctx context.Context, in *AddUserRemainderReq, opts ...grpc.CallOption) (*AddUserRemainderResp, error) {
	out := new(AddUserRemainderResp)
	err := c.cc.Invoke(ctx, Usercenter_AddUserRemainder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoResp, error) {
	out := new(UpdateUserInfoResp)
	err := c.cc.Invoke(ctx, Usercenter_UpdateUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) AddUserRemainderRollBack(ctx context.Context, in *AddUserRemainderReq, opts ...grpc.CallOption) (*AddUserRemainderResp, error) {
	out := new(AddUserRemainderResp)
	err := c.cc.Invoke(ctx, Usercenter_AddUserRemainderRollBack_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) FindUserInfoByUserIdList(ctx context.Context, in *FindUserInfoByUserIdListReq, opts ...grpc.CallOption) (*FindUserInfoByUserIdListResp, error) {
	out := new(FindUserInfoByUserIdListResp)
	err := c.cc.Invoke(ctx, Usercenter_FindUserInfoByUserIdList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsercenterServer is the server API for Usercenter service.
// All implementations must embed UnimplementedUsercenterServer
// for forward compatibility
type UsercenterServer interface {
	FindUserByUserIdList(context.Context, *FindUserByUserIdListReq) (*FindUserByUserIdListResp, error)
	SendEmailCode(context.Context, *SendEmailCodeReq) (*SendEmailCodeResp, error)
	AddUserRemainder(context.Context, *AddUserRemainderReq) (*AddUserRemainderResp, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoReq) (*UpdateUserInfoResp, error)
	AddUserRemainderRollBack(context.Context, *AddUserRemainderReq) (*AddUserRemainderResp, error)
	FindUserInfoByUserIdList(context.Context, *FindUserInfoByUserIdListReq) (*FindUserInfoByUserIdListResp, error)
	mustEmbedUnimplementedUsercenterServer()
}

// UnimplementedUsercenterServer must be embedded to have forward compatible implementations.
type UnimplementedUsercenterServer struct {
}

func (UnimplementedUsercenterServer) FindUserByUserIdList(context.Context, *FindUserByUserIdListReq) (*FindUserByUserIdListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByUserIdList not implemented")
}
func (UnimplementedUsercenterServer) SendEmailCode(context.Context, *SendEmailCodeReq) (*SendEmailCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailCode not implemented")
}
func (UnimplementedUsercenterServer) AddUserRemainder(context.Context, *AddUserRemainderReq) (*AddUserRemainderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserRemainder not implemented")
}
func (UnimplementedUsercenterServer) UpdateUserInfo(context.Context, *UpdateUserInfoReq) (*UpdateUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUsercenterServer) AddUserRemainderRollBack(context.Context, *AddUserRemainderReq) (*AddUserRemainderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserRemainderRollBack not implemented")
}
func (UnimplementedUsercenterServer) FindUserInfoByUserIdList(context.Context, *FindUserInfoByUserIdListReq) (*FindUserInfoByUserIdListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserInfoByUserIdList not implemented")
}
func (UnimplementedUsercenterServer) mustEmbedUnimplementedUsercenterServer() {}

// UnsafeUsercenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsercenterServer will
// result in compilation errors.
type UnsafeUsercenterServer interface {
	mustEmbedUnimplementedUsercenterServer()
}

func RegisterUsercenterServer(s grpc.ServiceRegistrar, srv UsercenterServer) {
	s.RegisterService(&Usercenter_ServiceDesc, srv)
}

func _Usercenter_FindUserByUserIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByUserIdListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).FindUserByUserIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_FindUserByUserIdList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).FindUserByUserIdList(ctx, req.(*FindUserByUserIdListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_SendEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).SendEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_SendEmailCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).SendEmailCode(ctx, req.(*SendEmailCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_AddUserRemainder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRemainderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).AddUserRemainder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_AddUserRemainder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).AddUserRemainder(ctx, req.(*AddUserRemainderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_UpdateUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_AddUserRemainderRollBack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRemainderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).AddUserRemainderRollBack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_AddUserRemainderRollBack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).AddUserRemainderRollBack(ctx, req.(*AddUserRemainderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_FindUserInfoByUserIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserInfoByUserIdListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).FindUserInfoByUserIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_FindUserInfoByUserIdList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).FindUserInfoByUserIdList(ctx, req.(*FindUserInfoByUserIdListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Usercenter_ServiceDesc is the grpc.ServiceDesc for Usercenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usercenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.usercenter",
	HandlerType: (*UsercenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindUserByUserIdList",
			Handler:    _Usercenter_FindUserByUserIdList_Handler,
		},
		{
			MethodName: "SendEmailCode",
			Handler:    _Usercenter_SendEmailCode_Handler,
		},
		{
			MethodName: "AddUserRemainder",
			Handler:    _Usercenter_AddUserRemainder_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _Usercenter_UpdateUserInfo_Handler,
		},
		{
			MethodName: "AddUserRemainderRollBack",
			Handler:    _Usercenter_AddUserRemainderRollBack_Handler,
		},
		{
			MethodName: "FindUserInfoByUserIdList",
			Handler:    _Usercenter_FindUserInfoByUserIdList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usercenter.proto",
}
