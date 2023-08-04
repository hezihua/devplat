package tools

import (
	"github.com/hezihua/devplat/mcenter/conf"

	_ "github.com/hezihua/devplat/mcenter/apps"

	"github.com/infraboard/mcube/app"
)

// 开发环境 单元测试的Setup
func DevelopmentSet() {
	// 测试测试用例的配置
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	// 初始化测试类的注册
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}
}
