// Package rbac @Author Bing
// @Date 2024/9/3 15:16:00
// @Desc
package rbac

import (
	"blog/internal/dao"
	"blog/internal/model"
	"blog/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

func (r *sRbac) PermissionPost(ctx context.Context, permission entity.Permission, menu []*model.Menu) (err error) {
	d := dao.Permission.Ctx(ctx)
	err = d.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		return permissionWithMenu(menu, permission, d)
	})
	return
}
func permissionWithMenu(menu []*model.Menu, permission entity.Permission, d *gdb.Model) (err error) {
	for _, m := range menu {
		permission.MenuId = m.Mid
		_, err = d.FieldsEx("id").Data(permission).Save()
		if len(m.Children) > 0 {
			err = permissionWithMenu(m.Children, permission, d)
		}
		if err != nil {
			return err
		}
	}

	return
}
func (r *sRbac) PermissionList(ctx context.Context) (permission []model.Permission, err error) {
	permission = []model.Permission{}
	d := dao.Permission.Ctx(ctx)
	err = d.WithAll().Scan(&permission)
	return
}
func (r *sRbac) PermissionOne(ctx context.Context, p entity.Permission) (permission model.Permission, err error) {
	permission = model.Permission{}
	d := dao.Permission.Ctx(ctx)
	err = d.Where("id", p.Id).WithAll().Limit(1).Scan(&permission)
	return
}

func (r *sRbac) PermissionUpdate(ctx context.Context, permission entity.Permission) (err error) {
	d := dao.Permission.Ctx(ctx)
	_, err = d.Where("id", permission.Id).FieldsEx("id").Data(permission).Update()
	return
}
func (r *sRbac) PermissionDelete(ctx context.Context, permission entity.Permission) (err error) {
	d := dao.Permission.Ctx(ctx)
	_, err = d.Where("id", permission.Id).Delete()
	return
}
