// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package config

import (
	"context"

	"blog/api/config/v1"
)

type IConfigV1 interface {
	Config(ctx context.Context, req *v1.ConfigReq) (res *v1.ConfigRes, err error)
}
