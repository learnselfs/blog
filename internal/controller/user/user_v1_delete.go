package user

import (
	"context"

	"blog/api/user/v1"
	"blog/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.User().Delete(ctx, req.Id)
	return
}
