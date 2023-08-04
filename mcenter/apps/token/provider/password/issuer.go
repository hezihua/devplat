package password

import (
	"context"
	"fmt"

	"github.com/hezihua/devplat/mcenter/apps/token"
	"github.com/hezihua/devplat/mcenter/apps/token/provider"
	"github.com/hezihua/devplat/mcenter/apps/user"
	"github.com/hezihua/devplat/mcenter/common/logger"

	"github.com/infraboard/mcube/app"
	"github.com/rs/xid"
)

type issuer struct {
	user user.Service
}

func (i *issuer) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (
	*token.Token, error) {
	req := user.NewDescibeUserRequest(in.Username)
	u, err := i.user.DescribeUser(ctx, req)
	if err != nil {
		return nil, err
	}
	err = u.CheckPassword(in.Password)
	if err != nil {
		logger.L().Debug().Msg(err.Error())
		return nil, fmt.Errorf("用户名或密码错误")
	}

	tk := token.NewToken()
	tk.AccessToken = xid.New().String()
	tk.RefreshToken = xid.New().String()

	return tk, nil
}

func (i *issuer) Config() error {
	i.user = app.GetInternalApp(user.AppName).(user.Service)
	return nil
}

func init() {
	provider.Registry(token.GRANT_TYPE_PASSWORD, &issuer{})
}
