package impl

import (
	"github.com/hezihua/devplat/mcenter/apps/token"
	"github.com/hezihua/devplat/mcenter/apps/token/provider"
	"github.com/hezihua/devplat/mcenter/conf"

	_ "github.com/hezihua/devplat/mcenter/apps/token/provider/all"

	"github.com/infraboard/mcube/app"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	svc = &impl{}
)

type impl struct {
	token.UnimplementedRPCServer

	col *mongo.Collection
}

func (i *impl) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	i.col = db.Collection("token")

	err = provider.Init()
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) Name() string {
	return token.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	token.RegisterRPCServer(server, i)
}

func init() {
	// grpc
	app.RegistryGrpcApp(svc)
	// internal app
	app.RegistryInternalApp(svc)

}
