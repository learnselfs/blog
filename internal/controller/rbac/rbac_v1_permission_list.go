package rbac

import (
	"blog/internal/service"
	"context"

	"blog/api/rbac/v1"
)

func (c *ControllerV1) PermissionList(ctx context.Context, req *v1.PermissionListReq) (res *v1.PermissionListRes, err error) {
	res = &v1.PermissionListRes{}
	res.Permission, err = service.Rbac().PermissionList(ctx)
	return
}
