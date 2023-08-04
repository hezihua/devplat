package auth

import (
	"strings"

	"github.com/hezihua/devplat/mpaas/common/logger"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"
)

// "github.com/emicklei/go-restful/v3"

func NewHttpAuther() *httpAuther {
	return &httpAuther{

	}
}

type httpAuther struct {

}

func (a *httpAuther) FilterFunction  (
	req *restful.Request, 
	resp *restful.Response, 
	next *restful.FilterChain) {
		//处理请求
		logger.L().Info().Msgf("%s", req)

		// 处理token
		authHeader := req.HeaderParameter("Authorization")
		tkl := strings.Split(authHeader, " ")
		if len(tkl) != 2 {
			response.Failed(resp, exception.NewUnauthorized("令牌不合法, 格式：breaer xxx"))
			return 
		}
		tk := tkl[1]
		logger.L().Debug().Msgf("get token: %s", tk)
    // 检查token 合法性


		next.ProcessFilter(req, resp)
		// 处理响应
		logger.L().Info().Msgf("%s", resp)

}