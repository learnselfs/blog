// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionRead is the golang structure of table permission_read for DAO operations like Where/Data.
type PermissionRead struct {
	g.Meta       `orm:"table:permission_read, do:true"`
	Id           interface{} // 标识
	PermissionId interface{} // 权限标识
	MenuId       interface{} // 目录标识
	Type         interface{} // 类型: 可选 read/write/null: 读/写/读写)
}
