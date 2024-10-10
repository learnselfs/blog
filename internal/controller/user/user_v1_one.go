package user

import (
	"context"

	"blog/api/user/v1"
	"blog/internal/service"
)

func (c *ControllerV1) One(ctx context.Context, req *v1.OneReq) (res *v1.OneRes, err error) {
	res = &v1.OneRes{}
	res.User, err = service.User().One(ctx, req.Id)
	return
}
