package model

import (
	"blog/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

type UserId struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender bool   `json:"gender"`
}

type User struct {
	g.Meta `table:"user"`
	entity.User
	Role `orm:"with:id=role_id"`
}

type UserID struct {
	Id int `json:"id"`
}

type UserList []User

type UserClaims struct {
	UserID   uint
	UserName string
	jwt.RegisteredClaims
}
