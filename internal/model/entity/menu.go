// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Menu is the golang structure for table menu.
type Menu struct {
	Mid       int    `json:"mid"       orm:"mid"        ` // 唯一标识
	ParentMid int    `json:"parentMid" orm:"parent_mid" ` // 父节点id
	Id        string `json:"id"        orm:"id"         ` // id
	Title     string `json:"title"     orm:"title"      ` // 标题
	Icon      string `json:"icon"      orm:"icon"       ` // 图标
	Type      int    `json:"type"      orm:"type"       ` // 类型
	OpenType  string `json:"openType"  orm:"open_type"  ` // 展开类型
	Href      string `json:"href"      orm:"href"       ` // 链接
}
