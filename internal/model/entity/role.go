// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Role is the golang structure for table role.
type Role struct {
	Id           int    `json:"id"           orm:"id"            ` // id
	Name         string `json:"name"         orm:"name"          ` // 角色名称
	Information  string `json:"information"  orm:"information"   ` // 角色信息
	PermissionId int    `json:"permissionId" orm:"permission_id" ` // 权限标识
}
