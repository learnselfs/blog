package rbac

import (
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"

	"blog/api/rbac/v1"
)

func (c *ControllerV1) PermissionUpdate(ctx context.Context, req *v1.PermissionUpdateReq) (res *v1.PermissionUpdateRes, err error) {
	permission := entity.Permission{Id: req.Id, RoleId: req.Rid, Name: req.Name, Information: req.Information, MenuId: req.MenuId}
	err = service.Rbac().PermissionUpdate(ctx, permission)
	return
}
