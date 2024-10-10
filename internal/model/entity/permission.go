// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Permission is the golang structure for table permission.
type Permission struct {
	Id             int    `json:"id"             orm:"id"              ` // id
	Name           string `json:"name"           orm:"name"            ` // 权限名称
	Information    string `json:"information"    orm:"information"     ` // 权限信息
	RoleId         int    `json:"roleId"         orm:"role_id"         ` // 角色标识
	MenuId         int    `json:"menuId"         orm:"menu_id"         ` // 菜单标识
	PermissionType string `json:"permissionType" orm:"permission_type" ` // 权限类型：read、write、'' 分别对应读权限、写权限、读写权限
}
