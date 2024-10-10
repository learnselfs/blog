// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package rbac

import (
	"context"

	"blog/api/rbac/v1"
)

type IRbacV1 interface {
	PermissionPost(ctx context.Context, req *v1.PermissionPostReq) (res *v1.PermissionPostRes, err error)
	PermissionOne(ctx context.Context, req *v1.PermissionOneReq) (res *v1.PermissionOneRes, err error)
	PermissionList(ctx context.Context, req *v1.PermissionListReq) (res *v1.PermissionListRes, err error)
	PermissionUpdate(ctx context.Context, req *v1.PermissionUpdateReq) (res *v1.PermissionUpdateRes, err error)
	PermissionDelete(ctx context.Context, req *v1.PermissionDeleteReq) (res *v1.PermissionDeleteRes, err error)
	RolePost(ctx context.Context, req *v1.RolePostReq) (res *v1.RolePostRes, err error)
	RoleOne(ctx context.Context, req *v1.RoleOneReq) (res *v1.RoleOneRes, err error)
	RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error)
	RoleUpdate(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error)
	RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error)
}
