// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// PermissionMenu is the golang structure for table permission_menu.
type PermissionMenu struct {
	Id           int    `json:"id"           orm:"id"            ` // 标识
	PermissionId int    `json:"permissionId" orm:"permission_id" ` // 权限标识
	MenuId       int    `json:"menuId"       orm:"menu_id"       ` // 目录标识
	Type         string `json:"type"         orm:"type"          ` // 类型: 可选 read/write/null: 读/写/读写)
}
