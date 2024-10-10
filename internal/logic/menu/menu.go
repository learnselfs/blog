// Package menu @Author Bing
// @Date 2024/9/6 11:45:00
// @Desc
package menu

import (
	"blog/internal/dao"
	"blog/internal/model"
	"blog/internal/model/entity"
	"blog/internal/packed"
	"blog/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMenu struct{}

func init() {
	service.RegisterMenu(New())
}
func New() *sMenu {
	return &sMenu{}
}

func (s *sMenu) Create(ctx context.Context, in entity.Menu) (err error) {
	orm := dao.Menu.Ctx(ctx)
	_, err = orm.Data(in).Save()
	return
}

func (s *sMenu) Update(ctx context.Context, in entity.Menu) (err error) {
	orm := dao.Menu.Ctx(ctx)
	_, err = orm.FieldsEx("mid").Where("mid", in.Mid).Update(in)
	return
}

func (s *sMenu) List(ctx context.Context, in entity.Menu) (out []*model.Menu, err error) {
	//orm := dao.Menu.Ctx(ctx)
	//out = make([]model.Menu, 0)
	//if len(in.Title) > 0 || len(in.Id) > 0 {
	//	err = orm.Where("parent_mid", 0).FieldsEx("parent_mid").WithAll().OmitNil().WhereOr("title", in.Title).WhereOr("id", in.Id).Scan(&out)
	//}
	//err = orm.Where("parent_mid", 0).FieldsEx("parent_mid").WithAll().OmitNil().Scan(&out)

	j := service.Token().JWT(ctx)
	var claims g.Map
	var user entity.User
	var u model.User
	var permission []model.Permission
	claims, _, err = j.GetClaimsFromJWT(ctx)
	gconv.Struct(claims, &user)
	orm := dao.User.Ctx(ctx)
	err = orm.Where("id", user.Id).Scan(&u)
	if err == nil {
		if err = dao.Permission.Ctx(ctx).WithAll().Where("role_id", u.CurrentRole).Scan(&permission); err == nil {
			out = packed.Menu(permission)
			// 包含 系统设置模块
		}
	}
	return
}

func (s *sMenu) One(ctx context.Context, in entity.Menu) (out model.Menu, err error) {
	orm := dao.Menu.Ctx(ctx)
	out = model.Menu{}
	err = orm.Where("id", in.Id).WhereOr("title", in.Id).With(entity.Menu{}).Limit(1).Scan(&out)
	return
}

func (s *sMenu) Delete(ctx context.Context, in entity.Menu) (err error) {
	orm := dao.Menu.Ctx(ctx)
	_, err = orm.Where("mid", in.Mid).Delete()
	return
}
