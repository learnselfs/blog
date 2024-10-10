// Package packed @Author Bing
// @Date 2024/9/5 11:03:00
// @Desc
package packed

import (
	"blog/internal/consts"
	"blog/internal/model"
	"blog/internal/model/entity"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func JwtClaim(ctx context.Context, in entity.User) (token string, err error) {

	claim := model.UserClaims{
		UserID:   uint(in.Id),
		UserName: in.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   in.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(consts.JwtKey))
	return
}

func JwtParse(ctx context.Context, token string) (claim *model.UserClaims, err error) {
	result := &jwt.Token{}
	result, err = jwt.ParseWithClaims(token, &model.UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})

	if err == nil && result.Valid {
		claim = &model.UserClaims{}
		var ok bool
		claim, ok = result.Claims.(*model.UserClaims)
		if ok {
			return claim, nil
		}
	} else {
		return nil, err
	}
	return
}
