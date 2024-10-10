// Package middleware @Author Bing
// @Date 2024/9/22 16:54:00
// @Desc
package middleware

import (
	"blog/internal/service"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct{}

var middleware = sMiddleware{}

func New() *sMiddleware {
	return &middleware
}

func init() {
	service.RegisterMiddleware(New())
}

func (s *sMiddleware) CORS(ctx context.Context, r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *sMiddleware) Auth(r *ghttp.Request) {
	service.Token().Authenticator(context.TODO())(r)
	r.Middleware.Next()
}
