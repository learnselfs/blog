// Package model @Author Bing
// @Date 2024/9/6 12:11:00
// @Desc
package model

import (
	"blog/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type MenuList struct {
}

type Menu struct {
	g.Meta `table:"menu"`
	entity.Menu
	Children []*Menu `orm:"with:parent_mid=mid" json:"children"`
}
