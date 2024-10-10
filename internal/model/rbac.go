package model

import (
	"blog/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type Role struct {
	g.Meta `orm:"table:role"`
	entity.Role
	Permission []*entity.Permission `orm:"with:role_id=permission_id" json:"permission"`
}

type Permission struct {
	g.Meta `orm:"table:permission"`
	Id     int `json:"id" orm:"id"`
	entity.Permission
	Role        []*entity.Role `orm:"with:permission_id=role_id" json:"role"`
	entity.Menu `orm:"with:mid=menu_id"`
	Href        string `json:"-,omitempty"`
	//PermissionMenu `orm:"with:permission_id=id" `
}

//type PermissionMenu struct {
//	g.Meta `orm:"table:permission_menu"`
//	entity.PermissionMenu
//	entity.Menu `orm:"with:mid=menu_id"`
//}

//type ManyRolePermission struct {
//	Id         int `json:"id"`
//	Rid        int `json:"rid"`
//	Pid        int `json:"pid"`
//	g.Meta     `orm:"table:many_role_permission"`
//	Permission []*Permission `orm:"with:id=pid"`
//	Role       []*Role       `orm:"with:id=rid"`
//}
