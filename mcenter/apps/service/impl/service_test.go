package impl_test

import (
	"testing"

	"github.com/hezihua/devplat/mcenter/apps/service"
	"github.com/hezihua/devplat/mcenter/test/tools"
)

func TestCreateService(t *testing.T) {
	req := &service.CreateServiceRequest{
		Domain:    "default",
		Namespace: "default",
		Name:      "mpaas",
	}
	ins, err := impl.CreateService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestDescribeService(t *testing.T) {
	req := &service.DescribeServiceRequest{DescribeValue: "cj6v2suigirpq254poq0", 
	DescribeBy: service.DESCRIBE_BY_SERVICE_CREDENTIAL_ID}
	ins, err := impl.DescribeService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}
