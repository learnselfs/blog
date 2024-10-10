package menu

import (
	"blog/api/menu/v1"
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	res = &v1.ListRes{}
	in := entity.Menu{Id: req.Id, Title: req.Title}
	res.Menu, err = service.Menu().List(ctx, in)
	return
}
