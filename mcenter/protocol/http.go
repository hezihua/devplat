package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hezihua/devplat/mcenter/apps/endpoint"
	"github.com/hezihua/devplat/mcenter/common/logger"
	"github.com/hezihua/devplat/mcenter/conf"
	"github.com/hezihua/devplat/mcenter/protocol/auth"
	"github.com/hezihua/devplat/mcenter/swagger"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
)

// NewHTTPService 构建函数
func NewHTTPService() *HTTPService {
	// restful root router
	r := restful.DefaultContainer
	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	// http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/emicklei/Projects/swagger-ui/dist"))))

	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"*"},
		AllowedDomains: []string{"*"},
		AllowedMethods: []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"},
		CookiesAllowed: false,
		Container:      r,
	}
	r.Filter(cors.Filter)
	r.Filter(auth.NewHttpAuther().FilterFunction)

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HTTP.Addr(),
		Handler:           r,
	}

	return &HTTPService{
		r:      r,
		server: server,
		c:      conf.C(),
	}
}

// HTTPService http服务
type HTTPService struct {
	r      *restful.Container
	c      *conf.Config
	server *http.Server
}

// mcenter/api/tokens/v1
// mcenter/api/users/v1
func (s *HTTPService) PathPrefix() string {
	return fmt.Sprintf("/%s/api", s.c.App.Name)
}

// Start 启动服务
func (s *HTTPService) Start() error {
	// 装置子服务路由
	app.LoadRESTfulApp(s.PathPrefix(), s.r)
	// 获取当前所有以装载的web sevice

	es := endpoint.NewRegistryRequest()
	wss := s.r.RegisteredWebServices()
	for i := range wss {
		ws := wss[i]
		routes := ws.Routes()
		for m := range routes {
			route := routes[m]
			fmt.Println(route)
			ep := &endpoint.CreateEndpointRequest{
				ServiceId: "cj6ujamigirsc4615i1g",
				Method:    route.Method,
				Path:      route.Path,
				Operation: route.Operation,
			}
			es.Add(ep)

		}
	}
	// API Doc
	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: swagger.Docs,
		DefinitionNameHandler: func(name string) string {
			if name == "state" || name == "sizeCache" || name == "unknownFields" {
				return ""
			}
			return name
		},
	}
	s.r.Add(restfulspec.NewOpenAPIService(config))
	logger.L().Info().Msgf("Get the API using http://%s%s", s.c.App.HTTP.Addr(), config.APIPath)

	// 启动 HTTP服务
	logger.L().Info().Msgf("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logger.L().Info().Msgf("service is stopped")
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

// Stop 停止server
func (s *HTTPService) Stop() error {
	logger.L().Info().Msgf("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		logger.L().Error().Msgf("graceful shutdown timeout, force exit")
	}
	return nil
}
