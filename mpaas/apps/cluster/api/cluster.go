package api

import (
	"github.com/emicklei/go-restful/v3"
)

type Exception struct {
	Code    int
	Message string
}

// 令牌颁发
func (h *handler) QueryCluster(r *restful.Request, w *restful.Response) {
	// 获取用户参数, 强制 content type 来识别body里面内容的格式
	w.Write([]byte("test"))
}
