package rbac

import (
	"blog/internal/model/entity"
	"context"

	"blog/api/rbac/v1"
	"blog/internal/service"
)

func (c *ControllerV1) RolePost(ctx context.Context, req *v1.RolePostReq) (res *v1.RolePostRes, err error) {
	role := entity.Role{PermissionId: req.PermissionId, Name: req.Name, Information: req.Information}
	res = &v1.RolePostRes{}
	res.Id, err = service.Rbac().RolePost(ctx, role)

	return
}
