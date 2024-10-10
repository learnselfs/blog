package menu

import (
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"

	"blog/api/menu/v1"
)

func (c *ControllerV1) One(ctx context.Context, req *v1.OneReq) (res *v1.OneRes, err error) {
	res = &v1.OneRes{}
	in := entity.Menu{Id: req.Id, Title: req.Title}
	res.Menu, err = service.Menu().One(ctx, in)
	return
}
