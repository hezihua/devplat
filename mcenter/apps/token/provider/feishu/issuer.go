package feishu

import (
	"context"
	"hzh/devcloud/mcenter/apps/token"
	"hzh/devcloud/mcenter/apps/token/provider"
)

type issuer struct{}

func (i *issuer) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (
	*token.Token, error) {
	return nil, nil
}

func (i *issuer) Config() error {
	return nil
}

func init() {
	provider.Registry(token.GRANT_TYPE_FEISHU, &issuer{})
}
