package rpc

import (
	"context"

	"github.com/hezihua/devplat/mcenter/apps/endpoint"
	"github.com/hezihua/devplat/mcenter/apps/token"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewClient todo
func NewClient(conf *Config) (*ClientSet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), conf.Timeout())
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		// mcenter Grpc 服务地址
		conf.Address,
		// 不使用TLS
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(conf),
	)

	if err != nil {
		return nil, err
	}

	return &ClientSet{
		conf: conf,
		conn: conn,
	}, nil
}

type ClientSet struct {
	conf *Config
	conn *grpc.ClientConn
}

// token rpc 客户端
func (c *ClientSet) Token() token.RPCClient {
	return token.NewRPCClient(c.conn)
}

// 功能管理 rpc 客户端
func (c *ClientSet) Endpoint() endpoint.RPCClient {
	return endpoint.NewRPCClient(c.conn)
}
