package menu

import (
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"

	"blog/api/menu/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	in := entity.Menu{
		Id:        req.Id,
		Mid:       req.Mid,
		Title:     req.Title,
		Type:      req.Type,
		OpenType:  req.Opentype,
		Icon:      req.Icon,
		Href:      req.Href,
		ParentMid: req.ParentMid,
	}
	err = service.Menu().Update(ctx, in)
	return
}
