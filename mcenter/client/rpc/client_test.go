package rpc_test

import (
	"context"
	"testing"

	"github.com/hezihua/devplat/mcenter/apps/token"
	"github.com/hezihua/devplat/mcenter/client/rpc"
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
	conf.ClientID = "cj6ujamigirsc4615i20"
	conf.ClientSecret = "cj6ujamigirsc4615i2g"
	c, err := rpc.NewClient(conf)
	if err != nil {
		panic(err)
	}
	client = c
}
