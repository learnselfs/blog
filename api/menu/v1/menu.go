// Package v1 @Author Bing
// @Date 2024/9/6 9:49:00
// @Desc
package v1

import (
	"blog/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta    `path:"/menu" method:"post" tag:"menu" dc:"菜单添加"`
	Id        string `p:"id"         `              // id
	Title     string `p:"title"      v:"required" ` // 标题
	Icon      string `p:"icon"       `              // 图标
	Type      int    `p:"type"       `              // 类型
	Opentype  string `p:"openType"   `              // 展开类型
	Href      string `p:"href"       `
	ParentMid int    `json:"parentMid" orm:"parent_mid" ` // 父节点id

}

type CreateRes struct {
	model.Menu `json:"menu"`
}
type UpdateReq struct {
	g.Meta    `path:"/menu/:mid" method:"put" tag:"menu" dc:"菜单修改"`
	Mid       int    `p:"mid"         v:"required"` // id
	Id        string `p:"id"         `              // id
	Title     string `p:"title"      `              // 标题
	Icon      string `p:"icon"       `              // 图标
	Type      int    `p:"type"       `              // 类型
	Opentype  string `p:"openType"   `              // 展开类型
	Href      string `p:"href"`
	ParentMid int    `p:"parent_mid"`
}

type UpdateRes struct {
}

type OneReq struct {
	g.Meta `path:"/menu/:id" method:"get" tag:"menu" dc:"单个菜单查询"`
	Id     string `p:"id"         ` // id
	Title  string `p:"title"      ` // 标题

}

type OneRes struct {
	model.Menu `json:"menu"`
}
type ListReq struct {
	g.Meta `path:"/menu" method:"get" tag:"menu" dc:"查询菜单"`
	Id     string `p:"id"         ` // id
	Title  string `p:"title"      ` // 标题
}

type ListRes struct {
	Menu []*model.Menu `json:"menu"`
}
type DeleteReq struct {
	g.Meta `path:"/menu/:mid" method:"delete" tag:"menu" dc:"菜单删除"`
	Mid    int `p:"mid"         v:"required"` // id
}

type DeleteRes struct{}
