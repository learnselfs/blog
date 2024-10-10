// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"blog/internal/model"
	"blog/internal/model/entity"
	"context"
)

type (
	IRbac interface {
		PermissionPost(ctx context.Context, permission entity.Permission, menu []*model.Menu) (err error)
		PermissionList(ctx context.Context) (permission []model.Permission, err error)
		PermissionOne(ctx context.Context, p entity.Permission) (permission model.Permission, err error)
		PermissionUpdate(ctx context.Context, permission entity.Permission) (err error)
		PermissionDelete(ctx context.Context, permission entity.Permission) (err error)
		RolePost(ctx context.Context, role entity.Role) (id int64, err error)
		RoleOne(ctx context.Context, role entity.Role) (roles model.Role, err error)
		RoleList(ctx context.Context) (roles []model.Role, err error)
		RoleUpdate(ctx context.Context, role entity.Role) (err error)
		RoleDelete(ctx context.Context, role entity.Role) (err error)
	}
)

var (
	localRbac IRbac
)

func Rbac() IRbac {
	if localRbac == nil {
		panic("implement not found for interface IRbac, forgot register?")
	}
	return localRbac
}

func RegisterRbac(i IRbac) {
	localRbac = i
}
