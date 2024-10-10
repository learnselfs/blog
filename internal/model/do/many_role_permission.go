// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ManyRolePermission is the golang structure of table many_role_permission for DAO operations like Where/Data.
type ManyRolePermission struct {
	g.Meta `orm:"table:many_role_permission, do:true"`
	Id     interface{} // id
	Rid    interface{} // 角色 id
	Pid    interface{} // 权限 id
}
