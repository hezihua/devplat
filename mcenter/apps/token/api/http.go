package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"

	"hzh/devcloud/mcenter/apps/token"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
)

var (
	h = &handler{}
)

type handler struct {
	service token.Service
}

func (h *handler) Config() error {
	h.service = app.GetInternalApp(token.AppName).(token.Service)
	return nil
}

func (h *handler) Name() string {
	return token.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"登录"}
	// r.POST("/xxxxx", HandleFunc)
	// Go Restful 路由配置是链式的: ws.Router(路由配置)
	// ws.POST("/")  To(h.IssueToken)  都是路由装饰
	ws.Route(ws.POST("/").To(h.IssueToken).
		Doc("颁发令牌").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.IssueTokenRequest{}).
		// 添加了一个权限注释
		Metadata("auth", false).
		Metadata("resource", "令牌").
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))

	// ws.Route(ws.GET("/").To(h.ValidateToken).
	// 	Doc("验证令牌").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(token.ValidateTokenRequest{}).
	// 	Writes(token.Token{}).
	// 	Returns(200, "OK", token.Token{}))

	// ws.Route(ws.DELETE("/").To(h.RevolkToken).
	// 	Doc("撤销令牌").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Writes(token.Token{}).
	// 	Returns(200, "OK", token.Token{}).
	// 	Returns(404, "Not Found", nil))

	// ws.Route(ws.PATCH("/").To(h.ChangeNamespace).
	// 	Doc("切换空间").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(token.ChangeNamespaceRequest{}))

}

func init() {
	app.RegistryRESTfulApp(h)
}
