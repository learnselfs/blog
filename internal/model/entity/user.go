// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	Id          int    `json:"id"          orm:"id"           ` // id
	Username    string `json:"username"    orm:"username"     ` // 用户名称
	Age         int    `json:"age"         orm:"age"          ` // 用户年龄
	Gender      int    `json:"gender"      orm:"gender"       ` // 性别
	Password    string `json:"password"    orm:"password"     ` //
	RoleId      string `json:"roleId"      orm:"role_id"      ` // 角色标识
	CurrentRole int    `json:"currentRole" orm:"current_role" ` // 用户当前角色
}
