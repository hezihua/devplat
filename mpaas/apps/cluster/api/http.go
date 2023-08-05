package api

import (
	"fmt"

	"github.com/hezihua/devplat/mpaas/apps/cluster"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
)

var (
	h = &handler{}
)

type handler struct {
	service cluster.Service
}

func (h *handler) Config() error {

	return nil
}

func (h *handler) Name() string {
	return cluster.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"k8s集群配置管理"}
	fmt.Println(tags)
	// r.POST("/xxxxx", HandleFunc)
	// Go Restful 路由配置是链式的: ws.Router(路由配置)
	// ws.POST("/")  To(h.IssueToken)  都是路由装饰
	ws.Route(ws.GET("/").To(h.QueryCluster).
		Doc("查询集群列表").
		Metadata("auth", true).
		Metadata(restfulspec.KeyOpenAPITags, tags))

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
