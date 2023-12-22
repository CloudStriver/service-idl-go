// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package usercenter

import (
	"context"

	"github.com/CloudStriver/service-idl-go/go-zero-gen/usercenter/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FindUserReq       = pb.FindUserReq
	FindUserResp      = pb.FindUserResp
	GenerateTokenReq  = pb.GenerateTokenReq
	GenerateTokenResp = pb.GenerateTokenResp
	RegisterReq       = pb.RegisterReq
	RegisterResp      = pb.RegisterResp
	SendEmailCodeReq  = pb.SendEmailCodeReq
	SendEmailCodeResp = pb.SendEmailCodeResp
	SendEmailReq      = pb.SendEmailReq
	SendEmailResp     = pb.SendEmailResp
	Token             = pb.Token
	UpdateUserReq     = pb.UpdateUserReq
	UpdateUserResp    = pb.UpdateUserResp
	User              = pb.User

	Usercenter interface {
		SendEmailCode(ctx context.Context, in *SendEmailCodeReq, opts ...grpc.CallOption) (*SendEmailCodeResp, error)
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
		FindUser(ctx context.Context, in *FindUserReq, opts ...grpc.CallOption) (*FindUserResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

func (m *defaultUsercenter) SendEmailCode(ctx context.Context, in *SendEmailCodeReq, opts ...grpc.CallOption) (*SendEmailCodeResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SendEmailCode(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUsercenter) FindUser(ctx context.Context, in *FindUserReq, opts ...grpc.CallOption) (*FindUserResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.FindUser(ctx, in, opts...)
}

func (m *defaultUsercenter) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUsercenter) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}
