// Package v1 @Author Bing
// @Date 2024/9/6 23:52:00
// @Desc
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type LoginReq struct {
	g.Meta   `path:"/login" tags:"user" method:"post" summary:"user login" dc:"用户登陆"`
	Username string `p:"username" v:"bail|required|length:4,10# 用户名必须4-10位"`
	Password string `p:"password" v:"required|length:4,16# 密码必填"`
}

type LoginRes struct {
	//entity.User
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}
