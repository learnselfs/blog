// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta      `orm:"table:user, do:true"`
	Id          interface{} // id
	Username    interface{} // 用户名称
	Age         interface{} // 用户年龄
	Gender      interface{} // 性别
	Password    interface{} //
	RoleId      interface{} // 角色标识
	CurrentRole interface{} // 用户当前角色
}
