package sUser

import (
	"blog/internal/dao"
	"blog/internal/model"
	"blog/internal/model/entity"
	"blog/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.RegisterUser(New())
}

type sUser struct{}

func New() *sUser { return &sUser{} }

func (u *sUser) Create(ctx context.Context, in entity.User) (out entity.User, err error) {
	orm := dao.User.Ctx(ctx)
	_, err = orm.Data(in).Save()
	out = in
	return
}

func (u *sUser) List(ctx context.Context) (out model.UserList, err error) {
	orm := dao.User.Ctx(ctx)
	err = orm.WithAll().Scan(&out)
	return
}

func (u *sUser) One(ctx context.Context, id int) (out model.User, err error) {
	out = model.User{}
	orm := dao.User.Ctx(ctx)
	err = orm.Where("id", id).WithAll().Limit(1).Scan(&out)
	return
}

func (u *sUser) Update(ctx context.Context, user entity.User) (err error) {
	m := dao.User.Ctx(ctx)
	_, err = m.Where("id", user.Id).FieldsEx("id").OmitEmpty().Update(gconv.Map(user))
	return
}

func (u *sUser) Delete(ctx context.Context, id int) (err error) {
	m := dao.User.Ctx(ctx)
	_, err = m.Where("id", id).Delete()
	return
}

func (u *sUser) Login(ctx context.Context, in entity.User) (user entity.User, err error) {
	user = entity.User{}
	orm := dao.User.Ctx(ctx)
	err = orm.OmitEmpty().Where(in).Scan(&user)
	return

}
