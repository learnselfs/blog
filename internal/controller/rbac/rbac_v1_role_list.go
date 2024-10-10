package rbac

import (
	"blog/internal/service"
	"context"

	"blog/api/rbac/v1"
)

func (c *ControllerV1) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	res = &v1.RoleListRes{}
	res.Role, err = service.Rbac().RoleList(ctx)
	return
}
