package impl_test

import (
	"context"

	"github.com/hezihua/devplat/mcenter/apps/token"
	"github.com/hezihua/devplat/mcenter/test/tools"

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
