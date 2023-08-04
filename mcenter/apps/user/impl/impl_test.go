package impl_test

import (
	"context"
	"hzh/devcloud/mcenter/apps/user"
	"hzh/devcloud/mcenter/test/tools"

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
