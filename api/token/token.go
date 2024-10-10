// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package token

import (
	"context"

	"blog/api/token/v1"
)

type ITokenV1 interface {
	Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
}
