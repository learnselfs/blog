package token

import (
	"blog/internal/service"
	"context"

	"blog/api/token/v1"
)

func (c *ControllerV1) Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error) {
	res, err = service.Token().Refresh(ctx)
	return
}
