package rpc

import (
	"context"
	"time"

	"gitee.com/go-course/go9/projects/devcloud/mcenter/apps/service"
)

// NewDefaultConfig todo
func NewDefaultConfig() *Config {
	return &Config{
		Address:       "localhost:18010",
		TimeoutSecond: 10,
	}
}

// Config 客户端配置
type Config struct {
	Address string `json:"adress" toml:"adress" yaml:"adress" env:"MCENTER_GRPC_ADDRESS"`
	// 客户端认证
	ClientID     string `json:"client_id" toml:"client_id" yaml:"client_id" env:"MCENTER_CLINET_ID"`
	ClientSecret string `json:"client_secret" toml:"client_secret" yaml:"client_secret" env:"MCENTER_CLIENT_SECRET"`
	// 默认值10秒
	TimeoutSecond uint `json:"timeout_second" toml:"timeout_second" yaml:"timeout_second" env:"MCENTER_GRPC_TIMEOUT_SECOND"`
}

func (c *Config) Timeout() time.Duration {
	return time.Second * time.Duration(c.TimeoutSecond)
}

// GetRequestMetadata gets the current request metadata, refreshing tokens
// if required. This should be called by the transport layer on each
// request, and the data should be populated in headers or other
// context. If a status code is returned, it will be used as the status for
// the RPC (restricted to an allowable set of codes as defined by gRFC
// A54). uri is the URI of the entry point for the request.  When supported
// by the underlying implementation, ctx can be used for timeout and
// cancellation. Additionally, RequestInfo data will be available via ctx
// to this call.  TODO(zhaoq): Define the set of the qualified keys instead
// of leaving it as an arbitrary string.
func (c *Config) GetRequestMetadata(ctx context.Context, uri ...string) (
	map[string]string, error) {
	return map[string]string{
		service.ClientHeaderKey: c.ClientID,
		service.ClientSecretKey: c.ClientSecret,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.
func (c *Config) RequireTransportSecurity() bool {
	return false
}

// PerRPCCredentials defines the common interface for the credentials which need to
// attach security information to every RPC (e.g., oauth2).
