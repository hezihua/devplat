package rpc_test

import (
	"context"
	"testing"

	"hzh/devcloud/mcenter/apps/token"
	"hzh/devcloud/mcenter/client/rpc"
)

var (
	client *rpc.ClientSet
	ctx    = context.Background()
)

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("cj67d66igiri0j2a76ng")
	tk, err := client.Token().ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func init() {
	conf := rpc.NewDefaultConfig()
	conf.ClientID = "cfsrgnh3n7pi7u2is880"
	conf.ClientSecret = "cfsrgnh3n7pi7u2is88g"
	c, err := rpc.NewClient(conf)
	if err != nil {
		panic(err)
	}
	client = c
}
