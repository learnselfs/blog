package rbac

import (
	"blog/internal/model/entity"
	"context"

	"blog/api/rbac/v1"
	"blog/internal/service"
)

func (c *ControllerV1) PermissionPost(ctx context.Context, req *v1.PermissionPostReq) (res *v1.PermissionPostRes, err error) {
	permission := entity.Permission{RoleId: req.RoleId, Name: req.Name, Information: req.Information, MenuId: req.MenuId}
	err = service.Rbac().PermissionPost(ctx, permission, req.Menu)
	return
}
