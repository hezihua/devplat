package impl_test

import (
	"context"

	"github.com/hezihua/devplat/mcenter/apps/service"
	"github.com/hezihua/devplat/mcenter/test/tools"

	"github.com/infraboard/mcube/app"
)

var (
	impl service.ServiceManager
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSet()

	impl = app.GetInternalApp(service.AppName).(service.ServiceManager)
}
