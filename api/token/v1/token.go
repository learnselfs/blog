// Package v1 @Author Bing
// @Date 2024/9/22 11:14:00
// @Desc
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

//type AuthLoginReq struct {
//	g.Meta `path:"/login" method:"post"`
//}
//
//type AuthLoginRes struct {
//	Token  string    `json:"token"`
//	Expire time.Time `json:"expire"`
//}

type RefreshReq struct {
	g.Meta `path:"/refresh_token" method:"post" tag:"token" dc:"令牌更新"`
}

type RefreshRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"get" tag:"user" dc:"用户退出" `
}

type LogoutRes struct {
}
