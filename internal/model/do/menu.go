// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Menu is the golang structure of table menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta    `orm:"table:menu, do:true"`
	Mid       interface{} // 唯一标识
	ParentMid interface{} // 父节点id
	Id        interface{} // id
	Title     interface{} // 标题
	Icon      interface{} // 图标
	Type      interface{} // 类型
	OpenType  interface{} // 展开类型
	Href      interface{} // 链接
}
