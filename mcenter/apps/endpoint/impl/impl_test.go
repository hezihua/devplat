package impl_test

import (
	"context"

	"github.com/hezihua/devplat/mcenter/apps/endpoint"
	"github.com/hezihua/devplat/mcenter/test/tools"

	"github.com/infraboard/mcube/app"
)

var (
	impl endpoint.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSet()

	impl = app.GetInternalApp(endpoint.AppName).(endpoint.Service)
}
