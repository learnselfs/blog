package rbac

import (
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"blog/api/rbac/v1"
)

func (c *ControllerV1) PermissionDelete(ctx context.Context, req *v1.PermissionDeleteReq) (res *v1.PermissionDeleteRes, err error) {
	err = service.Rbac().PermissionDelete(ctx, entity.Permission{Id: req.Id})
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
