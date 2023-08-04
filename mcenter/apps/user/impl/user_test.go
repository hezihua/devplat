package impl_test

import (
	"testing"

	"github.com/hezihua/devplat/mcenter/apps/user"
	"github.com/hezihua/devplat/mcenter/test/tools"
)

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "123456"
	ins, err := impl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	req.Keywords = "k"
	set, err := impl.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tools.MustToJson(set))

}

func TestDescribeUser(t *testing.T) {
	req := &user.DescribeUserRequest{
		DescribeBy:    user.DESCRIBE_BY_USERNAME,
		DescribeValue: "admin",
	}
	ins, err := impl.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tools.MustToJson(ins))
	err = ins.CheckPassword("123456")
	if err != nil {
		t.Fatal(err)
	}

}

func TestUpdateUser(t *testing.T) {
	impl.CreateUser(ctx, nil)
}

func TestDeleteUser(t *testing.T) {
	impl.CreateUser(ctx, nil)
}
