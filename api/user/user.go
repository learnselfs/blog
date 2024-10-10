// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"blog/api/user/v1"
)

type IUserV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	One(ctx context.Context, req *v1.OneReq) (res *v1.OneRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Post(ctx context.Context, req *v1.PostReq) (res *v1.PostRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}
