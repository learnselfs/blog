package rbac

import (
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"

	"blog/api/rbac/v1"
)

func (c *ControllerV1) RoleUpdate(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error) {
	err = service.Rbac().RoleUpdate(ctx, entity.Role{Id: req.Id, Name: req.Name, PermissionId: req.PermissionId, Information: req.Information})
	return
}
