package user

import (
	"blog/api/user/v1"
	"blog/internal/service"
	"context"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	res = &v1.ListRes{}
	res.List, err = service.User().List(ctx)
	return
}
