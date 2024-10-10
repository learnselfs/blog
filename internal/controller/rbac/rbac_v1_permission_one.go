package rbac

import (
	"blog/internal/dao"
	"blog/internal/model"
	"context"

	"blog/api/rbac/v1"
)

func (c *ControllerV1) PermissionOne(ctx context.Context, req *v1.PermissionOneReq) (res *v1.PermissionOneRes, err error) {
	d := dao.Permission.Ctx(ctx)
	var p model.Permission
	err = d.Where("id", req.Id).With(model.Role{}).Scan(&p)
	res = &v1.PermissionOneRes{Permission: p}
	return
}
