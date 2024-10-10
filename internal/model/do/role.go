// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Role is the golang structure of table role for DAO operations like Where/Data.
type Role struct {
	g.Meta       `orm:"table:role, do:true"`
	Id           interface{} // id
	Name         interface{} // 角色名称
	Information  interface{} // 角色信息
	PermissionId interface{} // 权限标识
}
