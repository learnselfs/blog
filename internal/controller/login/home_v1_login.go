package login

import (
	"blog/api/login/v1"
	"blog/internal/service"
	"context"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	//res = &v1.LoginRes{}
	//user := entity.User{}
	//user, err = service.User().Login(ctx, entity.User{Password: req.Password, Username: req.Username})
	//if err == nil {
	// go-jwt 5.0
	//res.Token, err = packed.JwtClaim(ctx, user)
	//}
	res = &v1.LoginRes{}
	//_, err = service.User().Login(ctx, entity.User{Password: req.Password, Username: req.Username})
	//if err == nil {
	// gf-jwt v2
	res.Token, res.Expire, err = service.Token().Login(ctx)
	//}
	return
}
