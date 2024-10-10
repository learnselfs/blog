package rbac

import (
	"blog/internal/dao"
	"blog/internal/model"
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"
)

type sRbac struct{}

func init() {
	service.RegisterRbac(New())
}

func New() *sRbac { return &sRbac{} }

func (r *sRbac) RolePost(ctx context.Context, role entity.Role) (id int64, err error) {
	d := dao.Role.Ctx(ctx)
	id, err = d.FieldsEx("id").OmitEmpty().InsertAndGetId(role)
	return
}

func (r *sRbac) RoleOne(ctx context.Context, role entity.Role) (roles model.Role, err error) {
	roles = model.Role{}
	d := dao.Role.Ctx(ctx)
	err = d.Where("id", role.Id).WithAll().Limit(1).Scan(&roles)
	return
}

func (r *sRbac) RoleList(ctx context.Context) (roles []model.Role, err error) {
	roles = []model.Role{}
	d := dao.Role.Ctx(ctx)
	err = d.WithAll().Scan(&roles)
	return
}

func (r *sRbac) RoleUpdate(ctx context.Context, role entity.Role) (err error) {
	d := dao.Role.Ctx(ctx)
	_, err = d.Data(role).FieldsEx("id").Where("id", role.Id).Update()
	return
}

func (r *sRbac) RoleDelete(ctx context.Context, role entity.Role) (err error) {
	d := dao.Role.Ctx(ctx)
	_, err = d.Where("id", role.Id).Delete()
	return
}
