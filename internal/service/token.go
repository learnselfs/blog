// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "blog/api/token/v1"
	"context"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IToken interface {
		Refresh(ctx context.Context) (res *v1.RefreshRes, err error)
		Logout(ctx context.Context)
		JWT(ctx context.Context) *jwt.GfJWTMiddleware
		Login(ctx context.Context) (token string, expire time.Time, err error)
		Authenticator(ctx context.Context) ghttp.HandlerFunc
		// IdentityHandler get the identity from JWT and set the identity for every request
		// Using this function, by r.GetParam("id") get identity
		IdentityHandler(ctx context.Context) interface{}
	}
)

var (
	localToken IToken
)

func Token() IToken {
	if localToken == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localToken
}

func RegisterToken(i IToken) {
	localToken = i
}
