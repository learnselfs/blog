package menu

import (
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"

	"blog/api/menu/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	in := entity.Menu{Mid: req.Mid}
	err = service.Menu().Delete(ctx, in)
	return
}
