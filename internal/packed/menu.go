// Package packed @Author Bing
// @Date 2024/9/24 12:37:00
// @Desc
package packed

import (
	"blog/internal/model"
	"sort"
)

func Menu(permissions []model.Permission) []*model.Menu {
	var menus []*model.Menu
	menuMap := make(map[int]*model.Menu)
	intSlic := make([]int, 0, len(permissions))
	for _, p := range permissions {
		intSlic = append(intSlic, p.Menu.Mid)
		menuMap[p.Menu.Mid] = &model.Menu{
			Menu:     p.Menu,
			Children: []*model.Menu{},
		}
	}
	sort.Ints(intSlic)
	for _, i := range intSlic {
		m := menuMap[i]
		if m.ParentMid == 0 {
			menus = append(menus, menuMap[m.Menu.Mid])
		} else {
			mc := menuMap[m.ParentMid]
			if mc.Children != nil {
				mc.Children = append(mc.Children, menuMap[m.Menu.Mid])
			}

		}
	}
	return menus
}
