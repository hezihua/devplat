package auth

import (
	"context"

	"github.com/hezihua/devplat/mcenter/apps/service"
	"github.com/infraboard/mcube/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// grpc 服务端 auth 中间件

func NewGrpcAuther() *grpcAuther {
	return &grpcAuther{
		service: app.GetInternalApp(service.AppName).(service.ServiceManager),
	}
}

type grpcAuther struct {
	service service.ServiceManager
}

func (g *grpcAuther) GetClientCredentialsFromMeta(md metadata.MD) (
	clientId, clientSecret string) {
	cids := md.Get(service.ClientHeaderKey)
	sids := md.Get(service.ClientSecretKey)
	if len(cids) > 0 {
		clientId = cids[0]
	}
	if len(sids) > 0 {
		clientSecret = sids[0]
	}
	return
}

func (g *grpcAuther) Authfunc(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// 获取客户端凭证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "需要认真")
	}
	clientId, clientSecret := g.GetClientCredentialsFromMeta(md)
	if clientId == "" || clientSecret == "" {

		return nil, status.Error(codes.PermissionDenied, "需要认证")
	}
	svc, err := g.service.DescribeService(ctx, &service.DescribeServiceRequest{
		DescribeBy:    service.DESCRIBE_BY_SERVICE_CREDENTIAL_ID,
		DescribeValue: clientId,
	})

	if err != nil {

		return nil, status.Error(codes.PermissionDenied, "认证错误")
	}

	if clientSecret != svc.Credential.ClientSecret {

		return nil, status.Error(codes.PermissionDenied, "用户名凭证认证失败")
	}
	resp, err = handler(ctx, req)
	return resp, err
}
