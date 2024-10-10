package v1

import (
	"blog/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type PermissionPostReq struct {
	g.Meta      `path:"/rbac/permission" method:"post" tag:"permission" dc:"权限添加"`
	RoleId      int           `p:"roleId"`
	Name        string        `p:"name" v:"required"`
	Information string        `p:"information"`
	MenuId      int           `p:"menuId"`
	Menu        []*model.Menu `p:"menu"`
}

type PermissionPostRes struct{}

type PermissionOneReq struct {
	g.Meta `path:"/rbac/permission/:id" method:"get" tag:"permission" dc:"权限查询"`
	Id     int `p:"id" v:"required"`
}

type PermissionOneRes struct {
	model.Permission `json:"permission"`
}

type PermissionListReq struct {
	g.Meta `path:"/rbac/permission" method:"get" tag:"permission" dc:"权限列表"`
}

type PermissionListRes struct {
	Permission []model.Permission `json:"permission"`
}

type PermissionUpdateReq struct {
	g.Meta      `path:"/rbac/permission/:id" method:"put" tag:"permission" dc:"权限更新"`
	Id          int    `p:"id" v:"required"`
	Rid         int    `p:"rid"`
	Name        string `p:"name" `
	Information string `p:"information" `
	MenuId      int    `p:"menuId"`
	Menu        string `json:"menu"`
}

type PermissionUpdateRes struct {
}

type PermissionDeleteReq struct {
	g.Meta `path:"/rbac/permission" method:"delete" tag:"permission" dc:"权限删除"`
	Id     int `p:"id" v:"required"`
}

type PermissionDeleteRes struct {
}
