// Package v1 @Author Bing
// @Date 2024/9/6 9:49:00
// @Desc
package v1

import "github.com/gogf/gf/v2/frame/g"

type ConfigReq struct {
	g.Meta `path:"/config" method:"get" tag:"config" dc:"配置接口"`
}

type ConfigRes struct {
	Config g.Map `json:"config"`
}
