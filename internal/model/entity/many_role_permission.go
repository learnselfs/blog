// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ManyRolePermission is the golang structure for table many_role_permission.
type ManyRolePermission struct {
	Id  int `json:"id"  orm:"id"  ` // id
	Rid int `json:"rid" orm:"rid" ` // 角色 id
	Pid int `json:"pid" orm:"pid" ` // 权限 id
}
