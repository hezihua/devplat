package provider

import (
	"context"

	"github.com/hezihua/devplat/mcenter/apps/token"
)

type Issuer interface {
	Config() error
	TokenIssuer
}

type TokenIssuer interface {
	IssueToken(context.Context, *token.IssueTokenRequest) (
		*token.Token, error)
}
