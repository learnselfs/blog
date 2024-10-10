package v1

import (
	"blog/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type RolePostReq struct {
	g.Meta       `path:"/rbac/role" method:"post" tag:"role" dc:"角色添加"`
	PermissionId int    `p:"permissionId" `
	Name         string `p:"name" v:"required"`
	Information  string `p:"information" `
	//entity.Role
}

type RolePostRes struct {
	Id int64 `json:"id"`
}

type RoleOneReq struct {
	g.Meta `path:"/rbac/role/:id" method:"get" tag:"role" dc:"角色查询"`
	Id     int `p:"id" v:"required"`
}

type RoleOneRes struct {
	model.Role `json:"role"`
}

type RoleListReq struct {
	g.Meta `path:"/rbac/role" method:"get" tag:"role" dc:"角色列表"`
}

type RoleListRes struct {
	Role []model.Role `json:"role"`
}
type RoleUpdateReq struct {
	g.Meta       `path:"/rbac/role/:id" method:"put" tag:"role" dc:"角色更新"`
	Id           int    `p:"id" v:"required"`
	PermissionId int    `p:"permissionId" `
	Name         string `p:"name" `
	Information  string `p:"information" `
}

type RoleUpdateRes struct {
}

type RoleDeleteReq struct {
	g.Meta `path:"/rbac/role/:id" method:"delete" tag:"role" dc:"角色删除"`
	Id     int `p:"id" v:"required"`
}

type RoleDeleteRes struct {
}
