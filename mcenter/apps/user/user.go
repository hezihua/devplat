package user

import (
	"encoding/json"

	"github.com/hezihua/devplat/mcenter/common/meta"

	"github.com/infraboard/mcube/http/request"
	"golang.org/x/crypto/bcrypt"
)

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

func (s *UserSet) Add(item *User) {
	s.Items = append(s.Items, item)
}

func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*meta.Meta
		*CreateUserRequest
	}{u.Meta, u.Spec})
}

func NewDefaultPageRequest() *User {
	return &User{
		Spec: NewCreateUserRequest(),
		Meta: meta.NewMeta(),
	}
}
func New(req *CreateUserRequest) (*User, error) {
	// 校验必填
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if err := req.HashPassword(); err != nil {
		return nil, err
	}
	return &User{
		Meta: meta.NewMeta(),
		Spec: req,
	}, nil
}

func (req *CreateUserRequest) HashPassword() error {
	hp, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return err
	}
	req.Password = string(hp)
	return nil
}

func (i *User) Desense() {
	i.Spec.Password = "****"
}

func (i *User) CheckPassword(pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(i.Spec.Password), []byte(pass))
}
