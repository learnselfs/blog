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
	IUser interface {
		Create(ctx context.Context, in entity.User) (out entity.User, err error)
		List(ctx context.Context) (out model.UserList, err error)
		One(ctx context.Context, id int) (out model.User, err error)
		Update(ctx context.Context, user entity.User) (err error)
		Delete(ctx context.Context, id int) (err error)
		Login(ctx context.Context, in entity.User) (user entity.User, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
