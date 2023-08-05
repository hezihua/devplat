package auth

import (
	"strings"

	"github.com/hezihua/devplat/mcenter/apps/token"
	"github.com/hezihua/devplat/mcenter/common/logger"

	"github.com/emicklei/go-restful/v3"
	"github.com/hezihua/devplat/mcenter/apps/endpoint"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"
)

// "github.com/emicklei/go-restful/v3"

func NewHttpAuther() *httpAuther {
	return &httpAuther{
		token: app.GetInternalApp(token.AppName).(token.Service),
	}
}

type httpAuther struct {
	token token.Service
}

func (a *httpAuther) FilterFunction(
	req *restful.Request,
	resp *restful.Response,
	next *restful.FilterChain) {
	// 权限判断 哪个功能
	accessRoute := req.SelectedRoute()

	ep := endpoint.Endpoint{
		Spec: &endpoint.CreateEndpointRequest{
			ServiceId: "cj6ujamigirsc4615i1g",
			Method:    accessRoute.Method(),
			Path:      accessRoute.Path(),
			Operation: accessRoute.Operation(),
		},
	}
	isAuth := accessRoute.Metadata()["auth"]
	if isAuth != nil {
		ep.Spec.Auth = isAuth.(bool)
	}
	if ep.Spec.Auth {

		//处理请求
		logger.L().Info().Msgf("%s", req)

		// 处理token 使用者是谁
		authHeader := req.HeaderParameter(token.TOKEN_HEADER_KEY)
		tkl := strings.Split(authHeader, " ")
		if len(tkl) != 2 {
			response.Failed(resp, exception.NewUnauthorized("令牌不合法, 格式：breaer xxx"))
			return
		}
		tk := tkl[1]
		logger.L().Debug().Msgf("get token: %s", tk)
		// 检查token 合法性

		tkobj, err := a.token.ValidateToken(
			req.Request.Context(),
			token.NewValidateTokenRequest(tk),
		)
		if err != nil {
			response.Failed(resp, exception.NewUnauthorized("令牌校验不合法"))
			return
		}
		// 放入上下文
		req.SetAttribute(token.ATTRIBUTE_TOKEN_KEY, tkobj)
	}
	next.ProcessFilter(req, resp)
	// 处理响应
	logger.L().Info().Msgf("%s", resp)

}
