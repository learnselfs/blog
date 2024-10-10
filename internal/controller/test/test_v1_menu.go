package test

import (
	"blog/hack"
	"context"

	"blog/api/test/v1"
)

func (c *ControllerV1) Menu(ctx context.Context, req *v1.MenuReq) (res *v1.MenuRes, err error) {
	hack.MenuLoad()
	return
}
