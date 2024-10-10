package menu

import (
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"

	"blog/api/menu/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	res = &v1.CreateRes{}
	in := entity.Menu{
		ParentMid: req.ParentMid,
		Id:        req.Id,
		Title:     req.Title,
		Icon:      req.Icon,
		Type:      req.Type,
		OpenType:  req.Opentype,
		Href:      req.Href,
	}
	err = service.Menu().Create(ctx, in)
	if err == nil {
		res.Menu, err = service.Menu().One(ctx, in)
	}
	return
}
