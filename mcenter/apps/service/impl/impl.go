package impl

import (
	"github.com/hezihua/devplat/mcenter/apps/service"
	"github.com/hezihua/devplat/mcenter/conf"

	"github.com/infraboard/mcube/app"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	svc = &impl{}
)

type impl struct {
	service.UnimplementedRPCServer

	col *mongo.Collection
}

func (i *impl) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	i.col = db.Collection(service.AppName)
	return nil
}

func (i *impl) Name() string {
	return service.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	service.RegisterRPCServer(server, i)
}

func init() {
	// grpc
	app.RegistryGrpcApp(svc)
	// internal app
	app.RegistryInternalApp(svc)

}
