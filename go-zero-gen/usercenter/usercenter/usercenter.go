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
	ConfirmCaptchaReq    = pb.ConfirmCaptchaReq
	ConfirmCaptchaResp   = pb.ConfirmCaptchaResp
	ConfirmEmailCodeReq  = pb.ConfirmEmailCodeReq
	ConfirmEmailCodeResp = pb.ConfirmEmailCodeResp
	GenerateTokenReq     = pb.GenerateTokenReq
	GenerateTokenResp    = pb.GenerateTokenResp
	GetCaptchaReq        = pb.GetCaptchaReq
	GetCaptchaResp       = pb.GetCaptchaResp
	GetUserReq           = pb.GetUserReq
	GetUserResp          = pb.GetUserResp
	LoginReq             = pb.LoginReq
	LoginResp            = pb.LoginResp
	Point                = pb.Point
	RefreshTokenReq      = pb.RefreshTokenReq
	RefreshTokenResp     = pb.RefreshTokenResp
	RegisterReq          = pb.RegisterReq
	RegisterResp         = pb.RegisterResp
	RetrievePasswordReq  = pb.RetrievePasswordReq
	RetrievePasswordResp = pb.RetrievePasswordResp
	SendEmailCodeReq     = pb.SendEmailCodeReq
	SendEmailCodeResp    = pb.SendEmailCodeResp
	SendEmailReq         = pb.SendEmailReq
	SendEmailResp        = pb.SendEmailResp
	Token                = pb.Token
	UpdateUserReq        = pb.UpdateUserReq
	UpdateUserResp       = pb.UpdateUserResp
	User                 = pb.User

	Usercenter interface {
		SendEmailCode(ctx context.Context, in *SendEmailCodeReq, opts ...grpc.CallOption) (*SendEmailCodeResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
		GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error)
		GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaResp, error)
		ConfirmCaptcha(ctx context.Context, in *ConfirmCaptchaReq, opts ...grpc.CallOption) (*ConfirmCaptchaResp, error)
		ConfirmEmailCode(ctx context.Context, in *ConfirmEmailCodeReq, opts ...grpc.CallOption) (*ConfirmEmailCodeResp, error)
		RetrievePassword(ctx context.Context, in *RetrievePasswordReq, opts ...grpc.CallOption) (*RetrievePasswordResp, error)
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

func (m *defaultUsercenter) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUsercenter) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUsercenter) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUsercenter) RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.RefreshToken(ctx, in, opts...)
}

func (m *defaultUsercenter) GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetCaptcha(ctx, in, opts...)
}

func (m *defaultUsercenter) ConfirmCaptcha(ctx context.Context, in *ConfirmCaptchaReq, opts ...grpc.CallOption) (*ConfirmCaptchaResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.ConfirmCaptcha(ctx, in, opts...)
}

func (m *defaultUsercenter) ConfirmEmailCode(ctx context.Context, in *ConfirmEmailCodeReq, opts ...grpc.CallOption) (*ConfirmEmailCodeResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.ConfirmEmailCode(ctx, in, opts...)
}

func (m *defaultUsercenter) RetrievePassword(ctx context.Context, in *RetrievePasswordReq, opts ...grpc.CallOption) (*RetrievePasswordResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.RetrievePassword(ctx, in, opts...)
}
