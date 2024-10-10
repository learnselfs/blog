package rbac

import (
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"

	"blog/api/rbac/v1"
)

func (c *ControllerV1) RoleOne(ctx context.Context, req *v1.RoleOneReq) (res *v1.RoleOneRes, err error) {
	res = &v1.RoleOneRes{}
	role := entity.Role{Id: req.Id}
	res.Role, err = service.Rbac().RoleOne(ctx, role)
	return
}
