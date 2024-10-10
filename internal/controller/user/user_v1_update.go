package user

import (
	"blog/api/user/v1"
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	in := entity.User{Id: req.Id, Username: req.Username, Password: req.Password, Gender: req.Gender, Age: req.Age, RoleId: req.RoleId, CurrentRole: req.CurrentRole}
	err = service.User().Update(ctx, in)
	return
}
