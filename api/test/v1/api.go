// Package v1 @Author Bing
// @Date 2024/9/6 16:58:00
// @Desc
package v1

import "github.com/gogf/gf/v2/frame/g"

type MenuReq struct {
	g.Meta `path:"/test/menu" method:"get" dc:"加载测试数据" tag:"test" dc:"菜单添加"`
}

type MenuRes struct{}
