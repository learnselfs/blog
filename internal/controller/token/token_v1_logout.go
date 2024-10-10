package token

import (
	"blog/internal/service"
	"context"

	"blog/api/token/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	//return nil, gerror.NewCode(gcode.CodeNotImplemented)
	service.Token().Logout(ctx)
	return
}
