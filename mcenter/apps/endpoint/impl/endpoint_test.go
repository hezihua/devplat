package impl_test

import (
	"testing"

	"github.com/hezihua/devplat/mcenter/apps/endpoint"
)

func TestCreateService(t *testing.T) {
	req := &endpoint.RegistryRequest{
		Items: []*endpoint.CreateEndpointRequest{},
	}
	req.Items = append(req.Items, &endpoint.CreateEndpointRequest{
		ServiceId: "xx",
		Method: "POST",
		Path: "xxxx",
		Operation: "xxx",
	})
	set, err := impl.RegistryEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
