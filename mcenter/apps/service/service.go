package service

import (
	"time"

	"github.com/hezihua/devplat/mcenter/common/meta"
	"github.com/rs/xid"
)

func New(in *CreateServiceRequest) *Service {
	return &Service{
		Meta: meta.NewMeta(),
		Spec: in,
		Credential: &Crendential{
			ClientId:     xid.New().String(),
			ClientSecret: xid.New().String(),
			CreateAt:     time.Now().Unix(),
		},
	}
}

func NewDefaultService() *Service {
	return New(&CreateServiceRequest{})
}
