package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" tag:"hello" dc:"测试接口"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
