package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/hezihua/devplat/mcenter/apps/token"
)

type Exception struct {
	Code    int
	Message string
}

// 查询
func (h *handler) QueryCluster(r *restful.Request, w *restful.Response) {
	attrTK := r.Attribute(token.ATTRIBUTE_TOKEN_KEY)
	if attrTK != nil {
		tk := attrTK.(*token.Token)
		fmt.Println(tk)
	}
	// 获取用户参数, 强制 content type 来识别body里面内容的格式
	w.Write([]byte("test"))
}
