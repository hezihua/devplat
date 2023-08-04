package user

import (
	context "context"

	"github.com/hezihua/devplat/mcenter/apps/domain"
	"github.com/hezihua/devplat/mcenter/common/validator"
)

const (
	AppName = "users"
)

type Service interface {
	// 创建用户
	CreateUser(ctx context.Context, in *CreateUserRequest) (*User, error)
	// 更新用户
	UpdateUser(ctx context.Context, in *UpdateUserRequest) (*User, error)
	// 删除用户
	DeleteUser(ctx context.Context, in *DeleteUserRequest) (*User, error)

	RPCServer
}

func (req *CreateUserRequest) Validate() error {
	if req.Domain == "" {
		req.Domain = domain.DEFAULT_DOMAIN
	}
	return validator.V().Struct(req)
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{}
}

func NewDescibeUserRequest(username string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy:    DESCRIBE_BY_USERNAME,
		DescribeValue: username,
	}
}
