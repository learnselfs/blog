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
	IMenu interface {
		Create(ctx context.Context, in entity.Menu) (err error)
		Update(ctx context.Context, in entity.Menu) (err error)
		List(ctx context.Context, in entity.Menu) (out []*model.Menu, err error)
		One(ctx context.Context, in entity.Menu) (out model.Menu, err error)
		Delete(ctx context.Context, in entity.Menu) (err error)
	}
)

var (
	localMenu IMenu
)

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}
