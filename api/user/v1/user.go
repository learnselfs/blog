package v1

import (
	"blog/internal/model"
	"blog/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta `path:"/user" tags:"user" method:"get" tag:"user" dc:"用户列表"`
}

type ListRes struct {
	List []model.User `json:"list"`
}

type OneReq struct {
	g.Meta `path:"/user/:id" tags:"user" method:"get"  tag:"user" dc:"用户列表"`
	Id     int `p:"id" v:"required"`
}

type OneRes struct {
	model.User
}

type UpdateReq struct {
	g.Meta      `path:"/user/:id" tags:"user" method:"put" tag:"user" dc:"用户更新"`
	Id          int    `p:"id" v:"required"`
	Username    string `p:"username" `
	Age         int    `p:"age" `
	Gender      int    `p:"gender" v:"in:0,1#gender只能传0，1" dc:"类型说明：（0：女，1：男）"`
	Password    string `p:"password"`
	RoleId      string `p:"role_id"`
	CurrentRole int    `p:"current_role"`
}

type UpdateRes struct {
	model.User
}

type PostReq struct {
	g.Meta `path:"/user/" tags:"user" method:"post" tag:"user" dc:"用户创建"`
	// Id int
	Username    string `p:"username" v:"required"`
	Age         int    `p:"age"`
	Gender      int    `p:"gender" v:"required|in:0,1#gender只能传0，1" dc:"类型说明：（0：女，1：男）"`
	Password    string `p:"password" v:"required"`
	RoleId      string `p:"role_id"`
	CurrentRole int    `p:"current_role"`
}

type PostRes struct {
	entity.User
}

type DeleteReq struct {
	g.Meta `path:"/user/:id" tags:"user" method:"delete" tag:"user" dc:"用户删除"`
	Id     int `p:"id"`
}

type DeleteRes struct {
	Username string `json:"username"`
}
