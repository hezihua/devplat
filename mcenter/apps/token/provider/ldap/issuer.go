package ldap

import (
	"context"

	"github.com/hezihua/devplat/mcenter/apps/token"
	"github.com/hezihua/devplat/mcenter/apps/token/provider"
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
	provider.Registry(token.GRANT_TYPE_LDAP, &issuer{})
}
