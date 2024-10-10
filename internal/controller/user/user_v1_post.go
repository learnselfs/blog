package user

import (
	"blog/api/user/v1"
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"
)

func (c *ControllerV1) Post(ctx context.Context, req *v1.PostReq) (res *v1.PostRes, err error) {
	// request := g.RequestFromCtx(ctx)
	// request.Response.Writeln(req)
	// u:=dao.User.Ctx(ctx)
	// _, err = u.Data(req).Save()
	res = &v1.PostRes{}
	in := entity.User{Username: req.Username, Age: req.Age, Gender: req.Gender, Password: req.Password, RoleId: req.RoleId, CurrentRole: req.CurrentRole}
	res.User, err = service.User().Create(ctx, in)
	return
}
