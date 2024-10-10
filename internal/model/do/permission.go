// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Permission is the golang structure of table permission for DAO operations like Where/Data.
type Permission struct {
	g.Meta         `orm:"table:permission, do:true"`
	Id             interface{} // id
	Name           interface{} // 权限名称
	Information    interface{} // 权限信息
	RoleId         interface{} // 角色标识
	MenuId         interface{} // 菜单标识
	PermissionType interface{} // 权限类型：read、write、'' 分别对应读权限、写权限、读写权限
}
