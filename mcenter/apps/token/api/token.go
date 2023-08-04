package api

import (
	"hzh/devcloud/mcenter/apps/token"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

type Exception struct {
	Code    int
	Message string
}

// 令牌颁发
func (h *handler) IssueToken(r *restful.Request, w *restful.Response) {
	// 获取用户参数, 强制 content type 来识别body里面内容的格式
	in := token.NewIssueTokenRequest()
	err := r.ReadEntity(in)
	if err != nil {
		// 异常的处理: Accept, 请求者希望 接收的数据格式
		// 抽象到公共库里面, mcube
		response.Failed(w, err)
		// w.WriteEntity(&Exception{Code: 400, Message: err.Error()})
		return
	}

	tk, err := h.service.IssueToken(r.Request.Context(), in)
	if err != nil {
		response.Failed(w, err)
		return
	}

	w.WriteEntity(tk)
}
