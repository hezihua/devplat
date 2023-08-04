package impl_test

import (
	"hzh/devcloud/mcenter/apps/token"
	"testing"

	"github.com/infraboard/mcube/exception"
)

func TestIssueToken(t *testing.T) {
	// TODO 用户名密码错误，异常信息 自定义
	req := &token.IssueTokenRequest{
		Username: "admin1",
		Password: "123456",
	}
	tk, err := impl.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("fdsafdsa")
	tk, err := impl.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err.(exception.APIException).ToJson())
		t.Fatal(err)
	}
	t.Log(tk)
}
