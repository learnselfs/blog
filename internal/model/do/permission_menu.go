// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionMenu is the golang structure of table permission_menu for DAO operations like Where/Data.
type PermissionMenu struct {
	g.Meta       `orm:"table:permission_menu, do:true"`
	Id           interface{} // 标识
	PermissionId interface{} // 权限标识
	MenuId       interface{} // 目录标识
	Type         interface{} // 类型: 可选 read/write/null: 读/写/读写)
}
