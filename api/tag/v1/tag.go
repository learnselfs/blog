// Package v1 @Author Bing
// @Date 2024/8/27 15:02:00
// @Desc
package v1

import "github.com/gogf/gf/v2/frame/g"

type TagReq struct {
	g.Meta `path:"/tag" method:"get" tag:"tag" dc:"标签添加"`
}

type TagRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
