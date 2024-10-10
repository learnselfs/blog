// Package middleware @Author Bing
// @Date 2024/9/6 23:22:00
// @Desc
package middleware

import (
	"blog/internal/packed"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareJwt(r *ghttp.Request) {
	token := r.Cookie.Get("token")
	_, err := packed.JwtParse(context.TODO(), token.String())
	if err != nil {
		r.Response.RedirectTo("/login", 302)
	}

	r.Middleware.Next()
}
