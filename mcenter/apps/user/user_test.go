package user_test

import (
	"testing"

	"github.com/hezihua/devplat/mcenter/apps/user"
)

func TestNewUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Password = "123456"
	err := req.HashPassword()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(req.Password)
}
