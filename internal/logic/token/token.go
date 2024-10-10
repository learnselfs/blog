// Package token @Author Bing
// @Date 2024/9/22 11:19:00
// @Desc
package token

import (
	v1 "blog/api/token/v1"
	"blog/internal/consts"
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

type sToken struct{ *jwt.GfJWTMiddleware }

func New() *sToken {
	s := &sToken{}
	s.GfJWTMiddleware = jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour * 2,
		MaxRefresh:      time.Hour * 2,
		IdentityKey:     "id",
		TokenLookup:     "header: Authorization, query: token, cookie: token",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: s.IdentityHandler,
	})
	return s
}

func init() {
	service.RegisterToken(New())
}
func (s *sToken) Refresh(ctx context.Context) (res *v1.RefreshRes, err error) {
	res = &v1.RefreshRes{}
	res.Token, res.Expire = s.RefreshHandler(ctx)
	return
}

func (s *sToken) Logout(ctx context.Context) {
	s.LogoutHandler(ctx)
	return
}
func (s *sToken) JWT(ctx context.Context) *jwt.GfJWTMiddleware {
	return s.GfJWTMiddleware
}
func (s *sToken) Login(ctx context.Context) (token string, expire time.Time, err error) {
	token, expire = s.LoginHandler(ctx)
	return
}
func (s *sToken) Authenticator(ctx context.Context) ghttp.HandlerFunc {
	return s.MiddlewareFunc()
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func (s *sToken) IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[s.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	//r.Response.WriteJson(g.Map{
	//	"code":    code,
	//	"message": message,
	//})
	consts.HandleException404(r)
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(ctx context.Context) (any, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in = entity.User{}
	)
	var user entity.User
	var err error
	if err = r.Parse(&in); err != nil {
		return nil, err
	}
	if user, err = service.User().Login(ctx, in); err == nil {

		return gconv.Map(user), nil
	}

	return nil, jwt.ErrFailedAuthentication
}
