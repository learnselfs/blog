// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package test

import (
	"context"

	"blog/api/test/v1"
)

type ITestV1 interface {
	Menu(ctx context.Context, req *v1.MenuReq) (res *v1.MenuRes, err error)
}
