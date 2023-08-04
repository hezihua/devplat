package impl

import (
	"context"

	"github.com/hezihua/devplat/mcenter/apps/token"
	"github.com/hezihua/devplat/mcenter/apps/token/provider"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (i *impl) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (
	*token.Token, error) {
	// 校验必填
	if err := in.Validate(); err != nil {
		return nil, err
	}

	issuer := provider.Get(in.GrantType)
	ins, err := issuer.IssueToken(ctx, in)
	if err != nil {
		return nil, err
	}

	_, err = i.col.InsertOne(ctx, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) ValidateToken(ctx context.Context, in *token.ValidateTokenRequest) (
	*token.Token, error) {
	// 查找token
	tk := token.NewToken()
	err := i.col.FindOne(ctx, bson.M{
		"_id": in.AccessToken,
	}).Decode(tk)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("%s not found", in.AccessToken)
		}
		return nil, err
	}

	err = tk.CheckAliable()
	if err != nil {
		return nil, err
	}
	return tk, nil
}
