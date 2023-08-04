package impl_test

import (
	"context"

	"github.com/hezihua/devplat/mcenter/apps/user"
	"github.com/hezihua/devplat/mcenter/test/tools"

	"github.com/infraboard/mcube/app"
)

var (
	impl user.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSet()

	impl = app.GetInternalApp(user.AppName).(user.Service)
}
