package service

import "context"

const (
	AppName = "services"
)

type ServiceManager interface {
	CreateService(context.Context, *CreateServiceRequest) (*Service, error)
	RPCServer
}
