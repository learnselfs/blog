package rbac

import (
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"

	"blog/api/rbac/v1"
)

func (c *ControllerV1) RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	err = service.Rbac().RoleDelete(ctx, entity.Role{Id: req.Id})
	return
}
