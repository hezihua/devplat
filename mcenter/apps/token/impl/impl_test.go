package impl_test

import (
	"context"
	"hzh/devcloud/mcenter/apps/token"
	"hzh/devcloud/mcenter/test/tools"

	"github.com/infraboard/mcube/app"
)

var (
	impl token.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSet()

	impl = app.GetInternalApp(token.AppName).(token.Service)
}
