package config

import (
	"blog/api/config/v1"
	"blog/internal/consts"
	"context"
)

func (c *ControllerV1) Config(ctx context.Context, req *v1.ConfigReq) (res *v1.ConfigRes, err error) {
	res = &v1.ConfigRes{}
	res.Config = consts.PearConfig
	return
}
